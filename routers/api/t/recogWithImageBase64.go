/*
@Time : 2020/4/29 14:27
@Author : FB
@File : recogWithImageBase64.go
@Software: GoLand
*/
package t

import (
	"github.com/gin-gonic/gin"
	"gocr/env"
	"gocr/services/common_service"
	"gocr/services/ocr_service"
	"net/http"
)

func RecogniteWithImageBase64(ctx *gin.Context) {

	data := ctx.Request.FormValue("data")
	ip := common_service.GetClientIP(ctx)

	if data == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Request Paramter don't allow null",
		})
		return
	}

	data, err := ocr_service.GetResultWithImageBase64(data, env.TYPE_OCR_GENERAL_BASIC, ip)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data":    data,
	})
	return
}
