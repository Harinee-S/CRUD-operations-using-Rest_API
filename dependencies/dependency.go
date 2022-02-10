package dependencies

import (
	"Project2/handlers"
	"Project2/services"
)

type Dependency struct {
	Handlers handler
}

func InitDependency() Dependency {
	sampleServ := services.NewSampleService()
	applicationHandler := handlers.NewApplicationHandler(sampleServ)
	return Dependency{Handlers: applicationHandler}
}
