package website

import (
	"context"
	"log"
	"net/http"
)

func (s *WebsiteService) startHTTPServer() error {
	s.server = &http.Server{Handler: s}
	return s.server.Serve(s.ln)
}

func (s *WebsiteService) shutdownHTTPServer() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.serverShutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}

// requestEventData is request event's data.
type requestEventData struct {
	// SessionID is the of corresponding request..
	SessionID string `json:"sessionID"`

	// Method is the request's method.
	Method string `json:"method"`

	// Host is the request's host.
	Host string `json:"host"`

	// Path is requested page's path.
	Path string `json:"path"`

	// IP address of client.
	IP string `json:"ip"`
}

func (s *WebsiteService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	se, err := newSession(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	s.addSession(se.id, se)
	if err := s.s.Emit(requestEventKey, requestEventData{
		SessionID: se.id,
		Method:    req.Method,
		Host:      req.Host,
		Path:      req.URL.Path,
	}); err != nil {
		log.Println(err)
		s.removeSession(se.id)
	}
	se.wait()
}
