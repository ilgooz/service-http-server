package httpserver

import (
	"log"
	"os"

	"github.com/ilgooz/service-http-server/httpserver/server"
	"github.com/ilgooz/service-http-server/x/xhttp"
	"github.com/sirupsen/logrus"
)

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

	// QS is the query string data.
	QS string `json:"qs"`

	// Body is the request data.
	Body string `json:"body"`

	// IP address of client.
	IP string `json:"ip"`
}

func (s *HTTPServerService) handleSession(ses *server.Session) {
	logrus.WithFields(logrus.Fields{
		"method": ses.Req.Method,
		"path":   ses.Req.URL.Path,
	}).Info("new request")

	if os.Getenv("ENABLE_CORS") == "true" {
		xhttp.CORS(ses.W)
	}

	if ses.Req.Method == "OPTIONS" {
		ses.End()
		return
	}

	// use cached response if exists.
	c := s.findCache(ses.Req.Method, ses.Req.URL.Path)
	if c != nil {
		if err := sendResponse(ses.W, response{
			code:     c.code,
			mimeType: c.mimeType,
			content:  c.content,
		}); err != nil {
			log.Println(err)
		}
		ses.End()
		logrus.WithFields(logrus.Fields{
			"method": ses.Req.Method,
			"path":   ses.Req.URL.Path,
		}).Info("responded from cache")
		return
	}

	qs, err := xhttp.JSONQuery(ses.Req)
	if err != nil {
		ses.End()
		log.Println(err)
		return
	}

	body, err := xhttp.BodyAll(ses.Req)
	if err != nil {
		ses.End()
		log.Println(err)
		return
	}

	s.addSession(ses)
	if err := s.service.Emit("request", requestEventData{
		SessionID: ses.ID,
		Method:    ses.Req.Method,
		Host:      ses.Req.Host,
		Path:      ses.Req.URL.Path,
		QS:        string(qs),
		Body:      string(body),
		IP:        ses.Req.RemoteAddr,
	}); err != nil {
		log.Println(err)
		s.removeSession(ses.ID)
	}
}

func (s *HTTPServerService) addSession(ses *server.Session) {
	s.ms.Lock()
	defer s.ms.Unlock()
	s.sessions[ses.ID] = ses
}

func (s *HTTPServerService) getSession(id string) (ses *server.Session, found bool) {
	s.ms.RLock()
	defer s.ms.RUnlock()
	ses, ok := s.sessions[id]
	return ses, ok
}

func (s *HTTPServerService) removeSession(id string) {
	s.ms.Lock()
	defer s.ms.Unlock()
	delete(s.sessions, id)
}
