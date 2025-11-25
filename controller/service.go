package controller

import (
	"fmt"
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
	Src   string `json:"src"`
	Proxy string `json:"proxy,omitempty"`
}
type ResponseBody struct {
	Dst string `json:"dst"`
	Msg   string `json:"msg,omitempty"`
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
	fmt.Println(requestBody)
	fmt.Println(requestBody.Src, requestBody.Proxy)
	var rep ResponseBody
	rep.Dst = logic.Trans(requestBody.Src, requestBody.Proxy)
	rep.Msg = ctx.ClientIP()
	//rep.Dst = fmt.Sprintf("我已经%d年没见过%s了", requestBody.Age, requestBody.Name)
	ctx.JSON(200, rep)
}
