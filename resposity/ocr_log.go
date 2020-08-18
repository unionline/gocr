/*
@Time : 2020/4/14 21:48
@Author : Justin
@Description :
@File : log.go
@Software: GoLand
*/
package resposity

import (
	"gocr/initializers/db"
	"gocr/models"
)

type OcrLogRepo struct {
}

func NewOcrLogRepo() *OcrLogRepo {
	return &OcrLogRepo{}
}

func (*OcrLogRepo) GetOcrLogInfo(id uint) (log models.OcrLog, err error) {
	log, err = Query.GetOcrLogInfo(db.Db, id)
	return
}

func (*OcrLogRepo) GetOcrLogInfoImageId(id string) (log models.OcrLog, err error) {
	log, err = Query.GetOcrLogInfoByImageId(db.Db, id)
	return
}

func (*OcrLogRepo) CreateOcrLog(log *models.OcrLog) (err error) {
	err = Insert.CreateOcrLog(db.Db, log)
	return
}
