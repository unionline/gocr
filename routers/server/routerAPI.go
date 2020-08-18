/*
@Time : 2020/4/20 14:54
@Author : FB
@File : routerLinux.go
@Software: GoLand
*/
package server

import (
	"github.com/gin-gonic/gin"
	"gocr/routers/api/v1"
)

func routerAPI(r *gin.Engine) {

	// 定义路由组
	// api 是有返回数据的接口
	api := r.Group("/api/")
	// users
	api.GET("users/", v1.UserInfo)
	api.POST("users/add", v1.CreateUser)

	// setting
	api.GET("setting/", v1.SettingInfo)
	api.POST("setting/add", v1.CreateSetting)

	// basic_setting
	api.GET("basic_setting/", v1.BasicSettingInfo)
	api.POST("basic_setting/add", v1.CreateBasicSetting)

	//// 暂时不提供
	// log
	// ocr_log

	// 文字识别
	api.POST("ocr/", v1.RecogniteWithUrl)
	api.GET("st/", v1.RecogniteWithUrl)

	// image 是只返回状态的接口
	// 上传模块
	image := r.Group("/upload/")
	image.POST("image/", v1.UploadImage)
	image.POST("images/", v1.UploadImages)

}
