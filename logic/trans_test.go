package logic

import (
	"log"
	"testing"
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
}
func TestTranslate(t *testing.T) {
	dst := TransByBing("hello")
	t.Log(dst)
}
