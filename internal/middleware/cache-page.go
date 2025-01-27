package middleware

import (
	"net/http"
	"strings"
	"sync"
)

var cache = make(map[string][]byte)
var mu sync.RWMutex

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
			mu.Lock()
			delete(cache, key)
			mu.Unlock()
			return
		}

		mu.RLock()
		if html, ok := cache[key]; ok {
			w.WriteHeader(http.StatusOK)
			w.Write(html)
			mu.RUnlock()
			return
		}

		rec := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK, body: []byte{}}
		next(rec, r)

		// Cache the response if the status code is 200
		if rec.statusCode == http.StatusOK {
			mu.Lock()
			cache[key] = rec.body
			mu.Unlock()
		}
	})
}
