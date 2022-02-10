package handlers

import (
	//"github.com/somnath-b/go-sample/services"
	//"Project2/services"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/suite"
)

type PingHandlerTestSuite struct {
	suite.Suite
	handler *ApplicationHandler
}

/*func (suite *PingHandlerTestSuite) SetupSuite() {
	suite.handler = NewApplicationHandler(services.NewSampleService())
}*/

func (suite *PingHandlerTestSuite) TestPingHandler() {
	rw := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/ping", nil)

	suite.handler.Ping(rw, r)

	suite.Nil(err)
	suite.Equal("{\"success\":true,\"data\":{\"message\":\"pong\"}}", rw.Body.String())
	suite.Equal(http.StatusOK, rw.Code)
	suite.Equal("application/json", rw.Header().Get("Content-Type"))
}

func TestPingHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerTestSuite))
}
