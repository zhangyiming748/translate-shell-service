package logic

import (
	"log"
	"os/exec"
	"strings"
	"time"
	"translate-shell-service/storage"
)

func Trans(src, proxy string) (dst string) {
	cache := new(storage.Cache)
	cache.Src = src
	if ok, _ := cache.GetBySrc(src); ok {
		log.Printf("从缓存中获取:%v\n", cache.Dst)
		return cache.Dst
	} else {
		log.Printf("从trans命令中获取:%v\n", cache.Dst)
		if proxy != "" {
			cache.Dst = TransByGoogle(src, proxy)
		} else {
			cache.Dst = TransByBing(src)
		}
	}
	err := cache.Create()
	if err != nil {
		log.Fatalf("创建缓存出错:%v\n", err)
	}
	return cache.Dst
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
	return result
}
