package httpserver

import "strings"

// cache keeps cache for http request.
type cache struct {
	// method of http request.
	method string

	// path of http request.
	path string

	// code of http response.
	code int

	// mimeType of http response.
	mimeType string

	// content of http response.
	content []byte
}

// addCache adds a new cache.
func (s *HTTPServerService) addCache(c *cache) {
	s.mc.Lock()
	defer s.mc.Unlock()
	s.caches = append(s.caches, c)
}

// findCache finds a cache for given filters.
func (s *HTTPServerService) findCache(method, path string) *cache {
	s.mc.RLock()
	defer s.mc.RUnlock()
	for _, c := range s.caches {
		if strings.ToLower(c.method) == strings.ToLower(method) &&
			strings.ToLower(c.path) == strings.ToLower(path) {
			return c
		}
	}
	return nil
}

// deleteCache deletes a cache for given filters.
func (s *HTTPServerService) deleteCache(method, path string) {
	s.mc.Lock()
	defer s.mc.Unlock()
	for i, c := range s.caches {
		if c.method == method && c.path == path {
			s.caches = append(s.caches[:i], s.caches[i+1:]...)
			return
		}
	}
}
