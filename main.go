package main

import (
	"net/http"
	"os"
	"translate-shell-service/bootstrap"
	"translate-shell-service/util"

	"github.com/gin-gonic/gin"
)

func testResponse(c *gin.Context) {
	c.JSON(http.StatusGatewayTimeout, gin.H{
		"code": http.StatusGatewayTimeout,
		"msg":  "timeout",
	})
}

func init() {
	os.Mkdir("/log", os.ModePerm)
	util.SetLog("/log/gin.log")
}
func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	bootstrap.InitService(engine)
	// 启动http服务
	engine.Run(":80")
}
