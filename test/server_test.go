package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	api "github.com/nuno-bastos/gin-gonic-wire-api/api"
	controller "github.com/nuno-bastos/gin-gonic-wire-api/api/controller"
)

func TestServerHTTP_Start(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	mockController := &controller.CalculateSecurityMatrixController{}

	// Create a new instance of the HTTP server
	server := api.StartServer(mockController)

	// Use httptest to create a mock HTTP request
	req, err := http.NewRequest("POST", "/github.com/nuno-bastos/gin-gonic-wire-api/CalculateGoSecurityMatrix", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Serve HTTP
	go server.Start() // Run the server asynchronously in a goroutine

	// Make the request against the server's engine
	server.GetEngine().ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
