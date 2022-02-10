package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	//"github.com/somnath-b/go-sample/contracts"
	"Project2/construct"
	"Project2/contracts"
	//"github.com/somnath-b/go-sample/contracts/sample"
	//"github.com/somnath-b/go-sample/contracts/sample"
)

func (handler *ApplicationHandler) SampleHandler(rw http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		contracts.ErrorResponse(rw, []string{"couldn't read body successfully"}, contracts.ErrorBadRequest)
	}

	var request construct.Users
	err = json.Unmarshal(body, &request)
	if err != nil {
		contracts.ErrorResponse(rw, []string{"couldn't unmarshal"}, contracts.ErrorBadRequest)
	}

	response, status, err := handler.sampleService.DoSomething(request)
	if err != nil {
		contracts.ErrorResponse(rw, []string{"invalid request"}, contracts.ErrorBadRequest)
	}

	contracts.SuccessResponse(rw, response, status)
}
