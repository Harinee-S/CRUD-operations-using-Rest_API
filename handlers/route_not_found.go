package handlers

import (
	"Project2/contracts"
	"fmt"
	"net/http"
)

// RouteNotFoundHandler - All requests are redirected to this router when a request is received for an unsupported path
func (handler *ApplicationHandler) RouteNotFoundHandler(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	errors := []string{fmt.Sprintf("route %s not found", path)}
	contracts.ErrorResponse(rw, errors, contracts.ErrorNotFound)
}
