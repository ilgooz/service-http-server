package httpserver

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joncalhoun/qson"
)

func (s *HTTPServerService) startHTTPServer() error {
	s.server = &http.Server{Handler: s}
	return s.server.Serve(s.ln)
}

func (s *HTTPServerService) shutdownHTTPServer() error {
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

	// Body is the request data.
	Body string `json:"body"`

	// IP address of client.
	IP string `json:"ip"`
}

func (s *HTTPServerService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	se, err := newSession(w, req)
	if err != nil {
		log.Println(err)
		return
	}
	s.addSession(se.id, se)

	var body string
	qs := req.URL.Query().Encode()
	if qs != "" {
		data, err := qson.ToJSON(qs)
		if err != nil {
			log.Println(err)
			return
		}
		body = string(data)
	} else {
		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println(err)
			return
		}
		body = string(data)
	}

	if err := s.s.Emit(requestEventKey, requestEventData{
		SessionID: se.id,
		Method:    req.Method,
		Host:      req.Host,
		Path:      req.URL.Path,
		Body:      body,
		IP:        req.RemoteAddr,
	}); err != nil {
		log.Println(err)
		s.removeSession(se.id)
	}
	se.wait()
}
