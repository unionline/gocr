/*
@Time : 2020/4/22 13:38
@Author : FB
@File : basicBasicSetting.go
@Software: GoLand
*/
package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/services/basic_setting_service"
	"net/http"
)

func BasicSettingInfo(ctx *gin.Context) {

	app_id := ctx.Request.FormValue("app_id")
	version := ctx.Request.FormValue("v")

	if app_id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "请求参数错误！",
		})

		return
	}

	settingInfo, err := basic_setting_service.GetBasicSettingInfo(app_id, version)
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
		"data":    settingInfo,
	})
	return
}

func CreateBasicSetting(ctx *gin.Context) {

	app_id := ctx.Request.FormValue("app_id")
	version := ctx.Request.FormValue("version")
	app_name := ctx.Request.FormValue("app_name")
	dev_log := ctx.Request.FormValue("dev_log")
	about := ctx.Request.FormValue("about")

	err := basic_setting_service.CreateBasicSetting(app_id, version, app_name, dev_log, about)
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
	})
	return
}
