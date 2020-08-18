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

type LogRepo struct {
}

func NewLogRepo() *LogRepo {
	return &LogRepo{}
}

func (*LogRepo) GetLogInfo(id uint) (log models.Log, err error) {
	log, err = Query.GetLogInfo(db.Db, id)
	return
}

func (*LogRepo) CreateLog(log *models.Log) (err error) {
	err = Insert.CreateLog(db.Db, log)
	return
}
