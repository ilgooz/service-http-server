package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlainTextResponse(t *testing.T) {
	w, server := newHTTPServerAndServer(t)

	go server.Start()
	go w.Start()
	defer w.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		event := <-server.LastEmit()
		var data requestEventData
		require.NoError(t, event.Data(&data))
		require.Equal(t, "/route0", data.Path)

		_, execution, err := server.Execute("completeSession", completeSessionInputs{
			SessionID: data.SessionID,
			Content:   "hello world!",
		})
		require.NoError(t, err)
		require.Equal(t, "success", execution.Key())

		var outputs completeSessionSuccessOutputs
		require.NoError(t, execution.Data(&outputs))
		require.Equal(t, data.SessionID, outputs.SessionID)
		require.True(t, outputs.ElapsedTime > 0)
	}()

	resp, err := http.Get(fmt.Sprintf("http://localhost%s/route0", w.ListeningAddr))
	require.NoError(t, err)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, "hello world!", string(data))

	wg.Wait()
}

func TestCompleteSessionError(t *testing.T) {
	var (
		data = completeSessionInputs{
			SessionID: "not-exists",
		}
		w, server = newHTTPServerAndServer(t)
	)

	go server.Start()
	go w.Start()
	defer w.Close()

	_, execution, err := server.Execute("completeSession", data)
	require.NoError(t, err)
	require.Equal(t, "error", execution.Key())

	var outputs errorOutput
	require.NoError(t, execution.Data(&outputs))
	require.Contains(t, "session not found", outputs.Message)
}
