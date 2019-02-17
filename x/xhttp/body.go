package xhttp

import (
	"io/ioutil"
	"net/http"
)

// BodyAll reads all body data into a byte slice.
func BodyAll(req *http.Request) ([]byte, error) {
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return data, nil
}
