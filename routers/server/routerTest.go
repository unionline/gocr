/*
@Time : 2020/4/20 14:54
@Author : FB
@File : routerLinux.go
@Software: GoLand
*/
package server

import (
	"github.com/gin-gonic/gin"
	"gocr/routers/api/t"
	"net/http"
)

func routerTest(r *gin.Engine) {

	// Test
	// 定义路由用来测试使用
	test := r.Group("/api/test")
	test.GET("get", func(ctx *gin.Context) {
		data := ctx.Request.FormValue("data")
		data2 := ctx.Query("data2")
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": gin.H{
				"data":  data,
				"data2": data2,
			},
		})
	})

	test.POST("post", func(ctx *gin.Context) {
		data := ctx.Request.FormValue("data")
		data2 := ctx.Query("data2")
		ctx.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"message": gin.H{
				"data":  data,
				"data2": data2,
			},
		})

	})

	// 测试 ocr image With base 64
	test.POST("ocr_b64", t.RecogniteWithImageBase64)
}
