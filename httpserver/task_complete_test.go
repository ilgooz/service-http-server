package httpserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"

	"github.com/mesg-foundation/core/client/service/servicetest"
	"github.com/stretchr/testify/require"
)

type request struct {
	name        string
	mimeType    string
	setMIMEType string
	content     string
	route       string
}

func TestResponses(t *testing.T) {
	var requests = []request{
		{
			name:     "text",
			mimeType: "text/plain",
			content:  "hello world!",
			route:    "/0",
		},
		{
			name:     "html",
			mimeType: "text/html",
			content:  "<b>hello world!</b>",
			route:    "/1",
		},
		{
			name:        "json",
			mimeType:    "application/json",
			setMIMEType: "application/json",
			content:     `{"hello": "world"}`,
			route:       "/2",
		},
	}

	hs, server := newHTTPServerAndServer(t)

	go server.Start()
	go hs.Start()
	defer hs.Close()

	for _, req := range requests {
		req := req
		t.Run("group", func(t *testing.T) {
			t.Run(req.name, func(t *testing.T) { testRequest(t, server, hs, req) })
		})
	}
}

func testRequest(t *testing.T, server *servicetest.Server, hs *HTTPServerService, req request) {
	t.Parallel()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		event := <-server.LastEmit()
		var data requestEventData
		require.NoError(t, event.Data(&data))
		require.Equal(t, req.route, data.Path)

		_, execution, err := server.Execute("completeSession", completeSessionInputs{
			SessionID: data.SessionID,
			Content:   req.content,
			MIMEType:  req.setMIMEType,
		})
		require.NoError(t, err)
		require.Equal(t, "success", execution.Key())

		var outputs completeSessionSuccessOutput
		require.NoError(t, execution.Data(&outputs))
		require.Equal(t, data.SessionID, outputs.SessionID)
		require.True(t, outputs.ElapsedTime > 0)
	}()

	resp, err := http.Get(fmt.Sprintf("http://localhost%s%s", hs.ListeningAddr(), req.route))
	require.NoError(t, err)
	defer resp.Body.Close()
	require.Contains(t, resp.Header.Get("content-type"), req.mimeType)

	data, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Equal(t, req.content, string(data))

	wg.Wait()
}

func TestCompleteSessionError(t *testing.T) {
	var (
		data = completeSessionInputs{
			SessionID: "not-exists",
		}
		hs, server = newHTTPServerAndServer(t)
	)

	go server.Start()
	go hs.Start()
	defer hs.Close()

	_, execution, err := server.Execute("completeSession", data)
	require.NoError(t, err)
	require.Equal(t, "error", execution.Key())

	var outputs errorOutput
	require.NoError(t, execution.Data(&outputs))
	require.Contains(t, "session not found", outputs.Message)
}
