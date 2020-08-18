package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/services/setting_service"
	"net/http"
)

var settingService *setting_service.SettingService

func init() {
	settingService = setting_service.NewSettingService()
}

// SettingInfo ...
func SettingInfo(ctx *gin.Context) {
	settingInfo, err := settingService.GetSettingInfo(1)
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

// SettingInfo ...
func CreateSetting(ctx *gin.Context) {

	app_id := ctx.Request.FormValue("appId")
	app_key := ctx.Request.FormValue("apiKey")
	app_secret := ctx.Request.FormValue("secretKey")
	type_setting := ctx.Request.FormValue("typeSetting")

	err := settingService.CreateSetting(app_id, app_key, app_secret, type_setting)
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
		"data":    "",
	})
	return
}
