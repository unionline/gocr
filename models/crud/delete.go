/*
@Time : 2020/4/14 22:54
@Author : Justin
@Description :
@File : delete
@Software: GoLand
*/
package crud

import (
	"github.com/jinzhu/gorm"
	"gocr/models"
)

type Delete struct {
}

func (*Delete) DeleteImage(db *gorm.DB, model *models.Image) (err error) {
	err = db.Where(models.Image{XID: model.XID, ID: model.ID}).Delete(&model).Error
	return
}
