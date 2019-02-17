package httpserver

import (
	"bytes"
	"io"
	"net/http"
)

// response keeps information about http response.
type response struct {
	code     int
	mimeType string
	content  []byte
}

// sendResponse sends a response to http request.
func sendResponse(w http.ResponseWriter, resp response) error {
	if resp.mimeType != "" {
		w.Header().Set("Content-Type", resp.mimeType)
	}

	if resp.code != 0 {
		w.WriteHeader(resp.code)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	if len(resp.content) == 0 && resp.code != http.StatusOK {
		if _, err := w.Write([]byte(http.StatusText(resp.code))); err != nil {
			return err
		}
	}
	if _, err := io.Copy(w, bytes.NewReader(resp.content)); err != nil {
		return err
	}
	return nil
}
