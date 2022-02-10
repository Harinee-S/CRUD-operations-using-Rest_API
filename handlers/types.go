package handlers

import "Project2/construct"

type ApplicationHandler struct {
	sampleService SampleService
}

type SampleService interface {
	DoSomething(construct.Users) (construct.JSONResponse, string, error)
}
