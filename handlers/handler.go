package handlers

//NewApplicationHandler ...
func NewApplicationHandler(sampleService SampleService) *ApplicationHandler {
	return &ApplicationHandler{
		sampleService: sampleService,
	}
}
