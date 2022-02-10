package services

import (
	"Project2/construct"
	"Project2/contracts"
	"errors"
	//"github.com/somnath-b/go-sample/contracts"
	//"github.com/somnath-b/go-sample/contracts/sample"
)

func (service *SampleService) DoSomething(request construct.Users) (construct.JSONResponse, string, error) {
	var response construct.JSONResponse
	if request.Email == "" || request.Password == "" || request.FullName == "" || request.PromoCode == "" || request.Reference == "" {
		return response, contracts.ErrorBadRequest, errors.New("invalid request")
	}
	response.Type = "type"
	//response.Data = []construct.Users
	response.Message = "message"
	return response, contracts.SuccessOK, nil
}

func NewSampleService() *SampleService {
	return &SampleService{}
}
