package httpserver

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// session keeps request, response writer pair for an HTTP request.
type session struct {
	req *http.Request
	w   http.ResponseWriter

	// id of the session
	id string

	// issuedAt is the creation time of session.
	issuedAt time.Time

	// waitC blocks http handler until is response sent.
	waitC chan struct{}
}

func newSession(w http.ResponseWriter, req *http.Request) (*session, error) {
	id := uuid.NewV4()
	idStr := id.String()
	return &session{
		w:        w,
		req:      req,
		id:       idStr,
		issuedAt: time.Now(),
		waitC:    make(chan struct{}),
	}, nil
}

func (s *session) wait() {
	<-s.waitC
}

func (s *session) done() {
	close(s.waitC)
}

func (s *HTTPServerService) addSession(id string, se *session) {
	s.ms.Lock()
	defer s.ms.Unlock()
	s.sessions[id] = se
}

func (s *HTTPServerService) getSession(id string) (se *session, found bool) {
	s.ms.Lock()
	defer s.ms.Unlock()
	se, ok := s.sessions[id]
	return se, ok
}

func (s *HTTPServerService) removeSession(id string) {
	s.ms.Lock()
	defer s.ms.Unlock()
	delete(s.sessions, id)
}
