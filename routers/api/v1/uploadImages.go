/*
@Time : 2020/4/18 14:10
@Author : FB
@File : uploadImages
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/services/image_service"
	"gocr/services/ocr_service"
	"io/ioutil"
	"log"
	"net/http"
)

func UploadImages(ctx *gin.Context) {

	// 图片
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	img_path_arr := make([]string, len(files))
	for i, file := range files {
		log.Println(file.Filename)

		src, err := file.Open()
		if err != nil {
			return
		}
		defer src.Close()

		data, err := ioutil.ReadAll(src)
		if err != nil {
			return
		}

		existed, img_path, err := image_service.SaveImageRecord(data, file)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
			return
		}

		img_path_arr[i] = img_path

		if existed {
			continue
		}

		var dst = img_path // 上传文件到指定的路径
		ctx.SaveUploadedFile(file, dst)
	}

	ip := ctx.ClientIP()

	// 暂时采用url模式
	// 在服务器上才有效
	ret, err := ocr_service.GetResultWithImagesUseUrl(img_path_arr, ip)
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
