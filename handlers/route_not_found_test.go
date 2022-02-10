package handlers

import (
	//"github.com/somnath-b/go-sample/services"
	"Project2/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
)

type RouteNotFoundHandlerTestSuite struct {
	suite.Suite
	handlers *ApplicationHandler
}

func (suite *RouteNotFoundHandlerTestSuite) SetupSuite() {
	suite.handlers = NewApplicationHandler(services.NewSampleService())
}

func (suite *RouteNotFoundHandlerTestSuite) TestRouteNotFoundHandlerShouldReturnNotFound() {
	rw := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/invalid_path", nil)
	suite.NoError(err, "failed to create a request")

	suite.handlers.RouteNotFoundHandler(rw, r)

	suite.Equal(http.StatusNotFound, rw.Code)
	suite.Equal("{\"success\":false,\"errors\":[{\"message\":\"route /invalid_path not found\"}]}", rw.Body.String())
	suite.Equal("application/json", rw.Header().Get("Content-Type"))
}

func TestRouteNotFoundHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RouteNotFoundHandlerTestSuite))
}
