/*
@Time : 2020/4/14 16:09
@Author : FB
@File : upload.go
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/services/image_service"
	"gocr/services/ocr_service"
	"io/ioutil"
	"net/http"
	"strconv"
)

func UploadImage(ctx *gin.Context) {

	r := ctx.Request

	image_type := r.FormValue("type")

	if image_type == "" {
		image_type = "0"
	}

	image_type_int, err := strconv.Atoi(image_type)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	// 图片
	file, header, err := r.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusRequestHeaderFieldsTooLarge, gin.H{
			"code":    http.StatusRequestHeaderFieldsTooLarge,
			"message": err.Error(),
		})
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	img_path, err := image_service.UploadImage(data, header)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	ip := ctx.ClientIP()

	// 识别上传图片文字
	// 图片很小时，可以，一般图片，没有返回值，原因不详
	//ret, err := ocr_service.GetResultWithImage(img_path, ip)
	//if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"code":    http.StatusInternalServerError,
	//		"message": err.Error(),
	//	})
	//	return
	//}

	// 暂时采用url模式
	// 在服务器上才有效
	//

	ret, err := ocr_service.GetResultWithImageUseUrl(img_path, image_type_int, ip)
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
		"data":    ret,
	})
	return
}
