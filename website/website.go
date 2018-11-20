package website

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	mesg "github.com/mesg-foundation/go-service"
)

// WebsiteService is a MESG service to serve website content over http.
type WebsiteService struct {
	s *mesg.Service

	ListeningAddr         string
	serverShutdownTimeout time.Duration
	server                *http.Server
	ln                    net.Listener

	// sessions is a request id, session pair.
	sessions map[string]*session
	ms       sync.Mutex
}

// New creates a new website service with given mesg service.
func New(service *mesg.Service, listeningAddr string) (*WebsiteService, error) {
	s := &WebsiteService{
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
func (s *WebsiteService) Start() error {
	defer s.Close()
	errC := make(chan error, 2)
	go func() { errC <- s.startHTTPServer() }()
	go func() { errC <- s.listenTasks() }()
	return <-errC
}

// Close gracefully closes the website service.
func (s *WebsiteService) Close() error {
	if err := s.s.Close(); err != nil {
		return err
	}
	return s.shutdownHTTPServer()
}
