package logic

import (
	"log"
	"testing"
	"translate-shell-service/storage"
	"translate-shell-service/util"
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
func TestTranslate(t *testing.T) {
	dst := TransByBing("hello")
	t.Log(dst)
}
