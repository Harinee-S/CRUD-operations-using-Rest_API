package server

import (
	"net/http"

	"Project2/dependencies"

	"github.com/gorilla/mux"
)

func Router(dep dependencies.Dependency) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping/", dep.Handlers.Ping).Methods("GET")
	router.HandleFunc("/sample", dep.Handlers.SampleHandler)
	router.HandleFunc("/getusers", dep.Handlers.GetUsers).Methods("GET")
	router.HandleFunc("/createusers", dep.Handlers.CreateUsers).Methods("POST")
	router.HandleFunc("/deleteuser/{email}", dep.Handlers.DeleteUsers).Methods("DELETE")
	router.HandleFunc("/updateuser/{email}", dep.Handlers.UpdateUsers).Methods("PUT")

	router.NotFoundHandler = http.HandlerFunc(dep.Handlers.RouteNotFoundHandler)
	return router
}
