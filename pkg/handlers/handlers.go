package handlers

import "net/http"

// Default handler to respond unknown endpoint
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
