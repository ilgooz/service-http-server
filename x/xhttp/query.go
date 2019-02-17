package xhttp

import (
	"net/http"

	"github.com/joncalhoun/qson"
)

// JSONQuery converts query string to json encoded byte slice.
func JSONQuery(req *http.Request) ([]byte, error) {
	qs := req.URL.Query().Encode()
	if qs != "" {
		return qson.ToJSON(qs)
	}
	return []byte("{}"), nil
}
