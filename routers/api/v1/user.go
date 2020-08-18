package v1

import (
	"github.com/gin-gonic/gin"
	"gocr/services/user_service"
	"net/http"
)

var userService *user_service.UserService

func init() {
	userService = user_service.NewUserService()
}

// UserInfo ...
func UserInfo(ctx *gin.Context) {
	userInfo, err := userService.GetUserInfo(1)
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
		"data":    userInfo,
	})
	return
}

// UserInfo ...
func CreateUser(ctx *gin.Context) {

	//username := ctx.Params.ByName("username")
	//password := ctx.Params.ByName("password")

	username := ctx.Query("username")
	password := ctx.Query("password")

	err := userService.CreateUser(username, password)
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
