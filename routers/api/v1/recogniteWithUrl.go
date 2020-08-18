/*
@Time : 2020/4/14 16:07
@Author : FB
@File : regWithUrl.go
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/env"
	"gocr/services/common_service"
	"gocr/services/ocr_service"
	"net/http"
)

func RecogniteWithUrl(ctx *gin.Context) {

	img_url := ctx.Request.FormValue("img_url")
	ip := common_service.GetClientIP(ctx)

	if img_url == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Request Paramter don't allow null",
		})
		return
	}

	data, err := ocr_service.GetResultWithUrl(img_url, env.TYPE_OCR_GENERAL_BASIC, ip)
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
