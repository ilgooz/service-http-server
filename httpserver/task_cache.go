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

type cacheSuccessOutput struct {
	Message string `json:"message"`
}

func (s *HTTPServerService) cacheHandler(execution *service.Execution) (string, interface{}) {
	var inputs cacheInputs
	if err := execution.Data(&inputs); err != nil {
		return errorOutputFrom(err)
	}

	s.deleteCache(inputs.Method, inputs.Path)
	s.addCache(&cache{
		method:   inputs.Method,
		path:     inputs.Path,
		code:     inputs.Code,
		mimeType: inputs.MIMEType,
		content:  []byte(inputs.Content),
	})
	logrus.WithFields(logrus.Fields{
		"method": inputs.Method,
		"path":   inputs.Path,
	}).Info("cached")

	return "success", cacheSuccessOutput{"ok"}
}
