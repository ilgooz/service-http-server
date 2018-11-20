package httpserver

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	mesg "github.com/mesg-foundation/go-service"
)

const completeSessionSuccessOutputKey = "success"

type completeSessionInputs struct {
	// SessionID id of request.
	SessionID string `json:"sessionID"`

	// Error should be filed if an error message or page is
	// showing as a response.
	// Content will still be used as the response to requested path.
	// If both Error and Code is set, Code has a higher priority.
	Error struct {
		IsNotFound bool `json:"isNotFound"`
		IsInternal bool `json:"isInternal"`
	} `json:"error"`

	// Code is the HTTP response code.
	// Default is 200.
	Code int `json:"code"`

	// MIMEType of content.
	// Server will try to figure out MIME type from content if not provided.
	MIMEType string `json:"mimeType"`

	// Content is the body of HTTP response.
	// If Content is empty but Error is not, a standard error
	// message will be shown related to the error type.
	// If Content is empty and Code is other than 200, a content message will
	// bbe shown related with the Code.
	// If Content is empty but MIME type is set, content will be formated
	// according to that otherwise it will be in plain text.
	Content string `json:"content"`
}

type completeSessionSuccessOutputs struct {
	SessionID string `json:"sessionID"`

	// ElapsedTime is in nanoseconds.
	ElapsedTime time.Duration `json:"elapsedTime"`
}

// completeSessionHandler is a task handler to complete a waiting request by
// sending a response to it with content, code, headers, MIME types and other configs.
func (s *HTTPServerService) completeSessionHandler(execution *mesg.Execution) (string, mesg.Data) {
	var inputs completeSessionInputs
	if err := execution.Data(&inputs); err != nil {
		return newErrorOutput(err)
	}

	se, found := s.getSession(inputs.SessionID)
	if !found {
		return newErrorOutput(errors.New("session not found"))
	}
	defer s.removeSession(se.id)
	defer se.done()

	if inputs.MIMEType != "" {
		se.w.Header().Set("Content-Type", inputs.MIMEType)
	}

	if inputs.Code != 0 {
		se.w.WriteHeader(inputs.Code)
	} else {
		se.w.WriteHeader(http.StatusOK)
	}

	if inputs.Content == "" && inputs.Code != http.StatusOK {
		if _, err := se.w.Write([]byte(http.StatusText(inputs.Code))); err != nil {
			return newErrorOutput(err)
		}
	} else {
		if _, err := io.Copy(se.w, strings.NewReader(inputs.Content)); err != nil {
			return newErrorOutput(err)
		}
	}

	return completeSessionSuccessOutputKey, completeSessionSuccessOutputs{
		SessionID:   se.id,
		ElapsedTime: time.Since(se.issuedAt),
	}
}
