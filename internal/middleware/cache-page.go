package middleware

import (
	"net/http"
	"strings"
)

var cache = make(map[string][]byte)

// responseRecorder captures the HTTP response
type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       []byte
}

func (rec *responseRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func (rec *responseRecorder) Write(body []byte) (int, error) {
	rec.body = append(rec.body, body...)
	return rec.ResponseWriter.Write(body)
}

func CachePage(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path

		if strings.HasPrefix(key, "/product") {
			next(w, r)
			delete(cache, key)
			return
		}
		if html, ok := cache[key]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write(html)
			return
		}

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK, body: []byte{}}
		next(rec, r)

		// Cache the response if the status code is 200
		if rec.statusCode == http.StatusOK {
			cache[key] = rec.body
		}
	})
}
