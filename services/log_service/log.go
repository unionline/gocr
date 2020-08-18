/*
@Time : 2020/4/14 23:37
@Author : Justin
@Description :
@File : log
@Software: GoLand
*/
package log_service

import (
	"gocr/models"
	"gocr/resposity"
)

var repo *resposity.LogRepo

func init() {
	repo = resposity.NewLogRepo()
}

func GetLogInfo(id uint) (resp models.Log, err error) {
	resp, err = repo.GetLogInfo(id)
	return
}

func CreateLog(item *models.Log) (err error) {

	err = repo.CreateLog(item)
	return
}
