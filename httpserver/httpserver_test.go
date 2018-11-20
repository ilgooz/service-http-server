package httpserver

import (
	"testing"

	mesg "github.com/mesg-foundation/go-service"
	"github.com/mesg-foundation/go-service/mesgtest"
	"github.com/stretchr/testify/require"
)

const (
	token    = "token"
	endpoint = "endpoint"
)

func newHTTPServerAndServer(t *testing.T) (*HTTPServerService, *mesgtest.Server) {
	testServer := mesgtest.NewServer()
	service, err := mesg.New(
		mesg.DialOption(testServer.Socket()),
		mesg.TokenOption(token),
		mesg.EndpointOption(endpoint),
	)
	require.NoError(t, err)
	require.NotNil(t, service)

	w, err := New(service, ":0")
	require.NoError(t, err)

	return w, testServer
}
