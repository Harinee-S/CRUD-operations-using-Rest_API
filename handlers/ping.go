package handlers

import (
	"net/http"

	//"github.com/somnath-b/go-sample/contracts"
	"Project2/contracts"
)

func (handler *ApplicationHandler) Ping(rw http.ResponseWriter, _ *http.Request) {
	response := contracts.PingResponse{Message: "pong"}
	contracts.SuccessResponse(rw, response, contracts.SuccessOK)
}
