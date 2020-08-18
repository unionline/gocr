/*
@Time : 2020/4/14 23:37
@Author : Justin
@Description :
@File : log
@Software: GoLand
*/
package ocr_log_service

import (
	"gocr/models"
	"gocr/resposity"
)

var repo *resposity.OcrLogRepo

func init() {
	repo = resposity.NewOcrLogRepo()
}

func GetOcrLogInfo(id uint) (resp models.OcrLog, err error) {
	resp, err = repo.GetOcrLogInfo(id)
	return
}

func GetOcrLogInfoByImageId(id string) (resp models.OcrLog, err error) {

	resp, err = repo.GetOcrLogInfoImageId(id)
	return
}

func CreateOcrLog(item *models.OcrLog) (err error) {

	err = repo.CreateOcrLog(item)
	return
}
