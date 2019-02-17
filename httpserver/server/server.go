package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	// ListeningAddr is full listening address of server.
	// it set after service start listening for connections.
	ListeningAddr   string
	addr            string
	hs              *http.Server
	ln              net.Listener
	shutdownTimeout time.Duration

	// Sessions filled with new sessions when new http requests received.
	Sessions chan *Session
}

// New creates a new HTTP server listening on addr.
func New(addr string) (*Server, error) {
	s := &Server{
		addr:            addr,
		shutdownTimeout: time.Minute,
		Sessions:        make(chan *Session),
	}
	s.hs = &http.Server{Handler: s}
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return nil, err
	}
	s.ln = ln
	s.ListeningAddr = fmt.Sprintf(":%d", s.ln.Addr().(*net.TCPAddr).Port)
	return s, nil
}

// Listen starts listening for client connections.
func (s *Server) Listen() error {
	return s.hs.Serve(s.ln)
}

// ServeHTTP implements http.Handler interface to handle http requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ses := newSession(w, req)
	s.Sessions <- ses
	<-ses.end
}

// Close gracefully closes http server.
func (s *Server) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	err := s.hs.Shutdown(ctx)
	close(s.Sessions)
	return err
}
