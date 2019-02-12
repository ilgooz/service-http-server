package httpserver

import (
	"github.com/mesg-foundation/core/client/service"
)

// list of task keys:
const (
	completeSessionTaskKey = "completeSession"
)

// list of events:
const (
	requestEventKey = "request"
)

// listenTasks starts listening for service's tasks.
func (s *HTTPServerService) listenTasks() error {
	return s.s.Listen(
		service.Task(completeSessionTaskKey, s.completeSessionHandler),
	)
}

// output key for errors.
const errOutputKey = "error"

// errorOutput is the error output data.
type errorOutput struct {
	Message string `json:"message"`
}

// newErrorOutput returns a new error output from given err.
func newErrorOutput(err error) (outputKey string, outputData interface{}) {
	return errOutputKey, errorOutput{Message: err.Error()}
}
