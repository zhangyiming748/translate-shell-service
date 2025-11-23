package bootstrap

import (
	"github.com/gin-gonic/gin"
	"translate-shell-service/controller"
)

func InitService(engine *gin.Engine) {
	routeGroup := engine.Group("/api")
	{
		c := new(controller.TranslateServiceController)
		routeGroup.GET("/v1/health", c.GetAlive)
		routeGroup.POST("/v1/translate", c.PostTranslate)
	}
}
