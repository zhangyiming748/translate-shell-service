package logic

import (
	"log"
	"os/exec"
	"strings"
	"time"
)

func Trans(src string) (dst string) {
	dst = TransByBing(src)
	return dst
}

func TransByGoogle(src, proxy string) (dst string) {
	cmd := exec.Command("trans", "-brief", "-engine", "google", "-proxy", proxy, ":zh-CN", src)
	output, err := cmd.CombinedOutput()
	result := string(output)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return src
	}
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("google查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		time.Sleep(3 * time.Second)
		TransByGoogle(src, proxy)
	}
	return result
}
func TransByBing(src string) (dst string) {
	cmd := exec.Command("trans", "-brief", "-engine", "bing", ":zh-CN", src)
	log.Printf("查询命令:%s\n", cmd.String())
	output, err := cmd.CombinedOutput()
	result := string(output)
	result = strings.Replace(result, "\\r\\n", "", 1)
	result = strings.Replace(result, "\n", "", 1)
	result = strings.Replace(result, "\r\n", "", 1)
	if result == "" {
		return src
	}
	if err != nil || strings.Contains(string(output), "u001b") || strings.Contains(string(output), "Didyoumean") || strings.Contains(string(output), "Connectiontimedout") {
		log.Printf("bing查询命令执行出错\t命令原文:%v\t错误原文:%v\n", cmd.String(), err.Error())
		time.Sleep(3 * time.Second)
		TransByBing(src)
	}
	log.Printf("%s执行查询命令后的译文:%s\n", src, result)
	return result
}
