package controller

import (
	"fmt"
	"os"
	"translate-shell-service/geo"
	"translate-shell-service/logic"

	"github.com/gin-gonic/gin"
)

type TranslateServiceController struct{}

/*
curl --location --request GET 'http://127.0.0.1:8192/api/v1/s1/gethello?user=<user>' \
--header 'User-Agent: Apifox/1.0.0 (https://www.apifox.cn)'
*/
func (tsc TranslateServiceController) GetAlive(ctx *gin.Context) {
	user := ctx.Query("user")
	ctx.String(200, fmt.Sprintf("Hello, %s!", user))
}

// 结构体必须大写 否则找不到
type RequestBody struct {
	Src   string `json:"src"` // 原文
	Proxy string `json:"proxy,omitempty"` // 本地运行时可选使用代理
	Abracadarbra string `json:"abracadarbra,omitempty"` // 设置一个keyword防止服务被滥用
}
type ResponseBody struct {
	Dst string `json:"dst"` // 译文
	Msg   geo.IpInfo `json:"msg,omitempty"` // 目前设置为返回请求客户端的IP地址
}

/*
 */
func (tsc TranslateServiceController) PostTranslate(ctx *gin.Context) {
	fmt.Println("get src")
	var requestBody RequestBody
	if err := ctx.BindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v\n",requestBody)
	if !isAuthorizedClient(){
		ctx.JSON(403, gin.H{"error": "unauthorized"})
	}
	var rep ResponseBody
	rep.Dst = logic.Trans(requestBody.Src, requestBody.Proxy)
	rep.Msg = geo.GetIPInfo(ctx.ClientIP())
	ctx.JSON(200, rep)
}
func isAuthorizedClient()bool{
	abracadarbra := os.Getenv("abracadabra")
	return abracadarbra == "abracadabra"
}