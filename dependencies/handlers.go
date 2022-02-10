package dependencies

import "net/http"

type handler interface {
	Ping(http.ResponseWriter, *http.Request)
	RouteNotFoundHandler(http.ResponseWriter, *http.Request)
	SampleHandler(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	CreateUsers(http.ResponseWriter, *http.Request)
	DeleteUsers(http.ResponseWriter, *http.Request)
	UpdateUsers(http.ResponseWriter, *http.Request)
}
