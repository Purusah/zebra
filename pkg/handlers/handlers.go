// General purpose default handlers
package handlers

import "net/http"

//NotFoundHandler Default handler to respond not defined endpoint
func NotFoundHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Not Found", http.StatusNotFound)
}
