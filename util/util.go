package util

import (
	"net/http"
)

// 현재 미사용중
// 얜 뭐임? 어따쓰는거지?
func CheckMethod(w http.ResponseWriter, r *http.Request, expectedMethod string) bool {
	if r.Method != expectedMethod {
		log := GetLogger()

		h := w.Header()
		h.Del("Content-Length")
		h.Set("Content-Type", "text/plain; charset=utf-8")
		h.Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusMethodNotAllowed)

		log.Error("Invalid request method: %s", r.Method, true)

		return false
	}
	return true
}
