/*
@Time : 2020/4/14 22:53
@Author : Justin
@Description :
@File : update.go
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gocr/models"
)

type Update struct {
}

func (*Update) UpdateImage(db *gorm.DB, model *models.Image) (err error) {
	err = db.Model(&model).Updates(models.Image{Etc: model.Etc}).Error
	return
}
