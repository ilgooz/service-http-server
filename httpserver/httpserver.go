package httpserver

import (
	"sync"

	"github.com/ilgooz/service-http-server/httpserver/server"
	"github.com/mesg-foundation/core/client/service"
)

// HTTPServerService is a MESG service to serve content over HTTP.
type HTTPServerService struct {
	service *service.Service
	server  *server.Server

	// sessions is a request id, session pair.
	sessions map[string]*server.Session
	ms       sync.RWMutex

	caches []*cache
	mc     sync.RWMutex
}

// New creates a new http server service with given mesg service.
func New(mesgService *service.Service, addr string) (*HTTPServerService, error) {
	s := &HTTPServerService{
		service:  mesgService,
		sessions: make(map[string]*server.Session),
		caches:   make([]*cache, 0),
	}
	server, err := server.New(addr)
	if err != nil {
		return nil, err
	}
	s.server = server
	go func() {
		for ses := range s.server.Sessions {
			go s.handleSession(ses)
		}
	}()
	return s, nil
}

// Start starts the location service.
func (s *HTTPServerService) Start() error {
	defer s.Close()
	errC := make(chan error, 2)
	go func() { errC <- s.server.Listen() }()
	go func() { errC <- s.listenTasks() }()
	err := <-errC
	s.Close()
	return err
}

func (s *HTTPServerService) ListeningAddr() string {
	return s.server.ListeningAddr
}

// listenTasks starts listening for service's tasks.
func (s *HTTPServerService) listenTasks() error {
	return s.service.Listen(
		service.Task("completeSession", s.completeSessionHandler),
		service.Task("cache", s.cacheHandler),
		service.Task("breakCache", s.breakCacheHandler),
	)
}

// Close gracefully closes the http server service.
func (s *HTTPServerService) Close() error {
	if err := s.service.Close(); err != nil {
		return err
	}
	return s.server.Close()
}
