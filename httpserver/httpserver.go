package httpserver

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/mesg-foundation/core/client/service"
)

// HTTPServerService is a MESG service to serve content over HTTP.
type HTTPServerService struct {
	s *service.Service

	ListeningAddr         string
	serverShutdownTimeout time.Duration
	server                *http.Server
	ln                    net.Listener

	// sessions is a request id, session pair.
	sessions map[string]*session
	ms       sync.Mutex
}

// New creates a new http server service with given mesg service.
func New(service *service.Service, listeningAddr string) (*HTTPServerService, error) {
	s := &HTTPServerService{
		s:                     service,
		sessions:              make(map[string]*session),
		ListeningAddr:         listeningAddr,
		serverShutdownTimeout: time.Second * 5,
	}
	ln, err := net.Listen("tcp", s.ListeningAddr)
	if err != nil {
		return nil, err
	}
	s.ln = ln
	s.ListeningAddr = fmt.Sprintf(":%d", ln.Addr().(*net.TCPAddr).Port)
	return s, nil
}

// Start starts the location service.
func (s *HTTPServerService) Start() error {
	defer s.Close()
	errC := make(chan error, 2)
	go func() { errC <- s.startHTTPServer() }()
	go func() { errC <- s.listenTasks() }()
	return <-errC
}

// Close gracefully closes the http server service.
func (s *HTTPServerService) Close() error {
	if err := s.s.Close(); err != nil {
		return err
	}
	return s.shutdownHTTPServer()
}
