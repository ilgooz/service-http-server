package httpserver

import (
	"testing"

	"github.com/mesg-foundation/core/client/service"
	"github.com/mesg-foundation/core/client/service/servicetest"
	"github.com/stretchr/testify/require"
)

const (
	token    = "token"
	endpoint = "endpoint"
)

func newHTTPServerAndServer(t *testing.T) (*HTTPServerService, *servicetest.Server) {
	testServer := servicetest.NewServer()
	s, err := service.New(
		service.DialOption(testServer.Socket()),
		service.TokenOption(token),
		service.EndpointOption(endpoint),
	)
	require.NoError(t, err)
	require.NotNil(t, s)

	w, err := New(s, ":0")
	require.NoError(t, err)

	return w, testServer
}
