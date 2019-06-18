package httpserver

import (
	"github.com/mesg-foundation/core/client/service"
	"github.com/sirupsen/logrus"
)

type cacheInputs struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	Code     int    `json:"code"`
	MIMEType string `json:"mimeType"`
	Content  string `json:"content"`
}

type cacheOutput struct {
	Message string `json:"message"`
}

func (s *HTTPServerService) cacheHandler(execution *service.Execution) (interface{}, error) {
	var inputs cacheInputs
	if err := execution.Data(&inputs); err != nil {
		return nil, err
	}
	s.cache(&cache{
		method:   inputs.Method,
		path:     inputs.Path,
		code:     inputs.Code,
		mimeType: inputs.MIMEType,
		content:  []byte(inputs.Content),
	})
	return cacheOutput{"ok"}, nil
}

func (s *HTTPServerService) cache(c *cache) {
	s.deleteCache(c.method, c.path)
	s.addCache(c)
	logrus.WithFields(logrus.Fields{
		"method": c.method,
		"path":   c.path,
	}).Info("cached")
}
