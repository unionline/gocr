/*
@Time : 2020/4/14 21:48
@Author : Justin
@Description :
@File : Image.go
@Software: GoLand
*/
package resposity

import (
	"gocr/initializers/db"
	"gocr/models"
)

type ImageRepo struct {
}

func NewImageRepo() *ImageRepo {
	return &ImageRepo{}
}

func (*ImageRepo) GetImageInfo(id uint) (item models.Image, err error) {
	item, err = Query.GetImageInfo(db.Db, id)
	return
}

func (*ImageRepo) GetImageInfoByMd5(md5 string) (item models.Image, err error) {
	item, err = Query.GetImageInfoByMd5(db.Db, md5)
	return
}

func (*ImageRepo) GetImageList() (items []models.Image, err error) {
	items, err = Query.GetImageList(db.Db)
	return
}

func (*ImageRepo) CreateImage(Image *models.Image) (err error) {
	tx := db.Db.Begin()
	err = Insert.CreateImage(tx, Image)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

func (*ImageRepo) DeleteImage(Image *models.Image) (err error) {
	tx := db.Db.Begin()
	err = Delete.DeleteImage(tx, Image)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}

func (*ImageRepo) UpdateImage(Image *models.Image) (err error) {
	tx := db.Db.Begin()
	err = Update.UpdateImage(tx, Image)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return
}
