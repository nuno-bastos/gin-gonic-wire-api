package server

import (
	"github.com/gin-gonic/gin"

	controller "github.com/nuno-bastos/gin-gonic-wire-api/api/controller"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func StartServer(calculateGoSecurityMatrixController *controller.CalculateSecurityMatrixController) *ServerHTTP {
	engine := gin.New()

	engine.ForwardedByClientIP = true
	engine.SetTrustedProxies([]string{"localhost"})
	engine.Use(gin.Logger())

	api := engine.Group("/github.com/nuno-bastos/gin-gonic-wire-api") // base api URL
	api.POST("/CalculateGoSecurityMatrix", calculateGoSecurityMatrixController.CalculateGoSecurityMatrix)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":8080")
}

func (sh *ServerHTTP) GetEngine() *gin.Engine {
	return sh.engine
}
