package httpserver

import (
	"github.com/mesg-foundation/core/client/service"
	"github.com/sirupsen/logrus"
)

type breakCacheInputs struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type breakCacheSuccessOutput struct {
	Message string `json:"message"`
}

func (s *HTTPServerService) breakCacheHandler(execution *service.Execution) (string, interface{}) {
	var inputs breakCacheInputs
	if err := execution.Data(&inputs); err != nil {
		return errorOutputFrom(err)
	}

	s.deleteCache(inputs.Method, inputs.Path)
	logrus.WithFields(logrus.Fields{
		"method": inputs.Method,
		"path":   inputs.Path,
	}).Info("cache deleted")

	return "success", breakCacheSuccessOutput{"ok"}
}
