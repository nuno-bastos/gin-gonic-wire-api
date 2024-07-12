package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	services "github.com/nuno-bastos/gin-gonic-wire-api/service/interface"
)

type CalculateSecurityMatrixController struct {
	_service services.SecurityMatrixService
}

func NewCalculateGoSecurityMatrixController(service services.SecurityMatrixService) *CalculateSecurityMatrixController {
	return &CalculateSecurityMatrixController{
		_service: service,
	}
}

// CalculateGoSecurityMatrix handles the HTTP POST request to calculate the security matrix.
// It delegates the security matrix calculation to the service layer, which writes the calculated rules to the database,
// and responds with a success message.
//
// Parameters:
// - c: Context object representing the HTTP request and response.
func (p *CalculateSecurityMatrixController) CalculateGoSecurityMatrix(c *gin.Context) {
	p._service.CalculateGoSecurityMatrix(c.Request.Context())

	c.JSON(http.StatusOK, gin.H{
		"message": "Security Matrix Rules generated and written successfully :)",
	})
}
