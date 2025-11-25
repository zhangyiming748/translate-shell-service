package util

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/zhangyiming748/lumberjack"
)

func SetLog(root string) {
	os.Mkdir(root, os.ModePerm)
	logFile := filepath.Join(root, "gin.log")
	// 创建一个用于写入文件的Logger实例
	fileLogger := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    1, // MB
		MaxBackups: 3,
		MaxAge:     28, // days
	}
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
