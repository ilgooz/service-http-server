package server

import (
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Session keeps a new http requests and its response writer.
type Session struct {
	// ID is a unique id given to the session.
	ID string

	// IssuedAt is the creation time of session.
	IssuedAt time.Time

	// W holds response writer.
	W http.ResponseWriter

	// Req holds the http request.
	Req *http.Request

	end chan struct{}
}

func newSession(w http.ResponseWriter, req *http.Request) *Session {
	return &Session{
		ID:       uuid.NewV4().String(),
		IssuedAt: time.Now(),
		W:        w,
		Req:      req,
		end:      make(chan struct{}),
	}
}

// End ends response stream.
// It should be called after done writing to response writer.
func (s *Session) End() {
	close(s.end)
}
