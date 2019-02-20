package xhttp

import (
	"encoding/json"
	"net/http"
)

// JSONQuery converts query string to json encoded byte slice.
func JSONQuery(req *http.Request) ([]byte, error) {
	return json.Marshal(req.URL.Query())
}
