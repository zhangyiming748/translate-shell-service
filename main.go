package main

import (
	"translate-shell-service/bootstrap"
	"translate-shell-service/storage"
	"translate-shell-service/util"
"log"
	"github.com/gin-gonic/gin"
)

func init() {
	var baseDir string
	if util.IsRunningInContainer() {
		log.Println("运行在容器中")
		baseDir = "/"
	} else {
		log.Println("运行在主机中")
		baseDir = "."
	}
	util.SetLog(baseDir)
	storage.SetSqlite(baseDir)
	new(storage.Cache).Sync()
}
func main() {
	// gin服务
	gin.SetMode(gin.DebugMode)
	engine := gin.New()
	bootstrap.InitService(engine)
	// 启动http服务
	engine.Run(":6380")
}
