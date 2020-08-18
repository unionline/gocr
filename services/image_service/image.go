/*
@Time : 2020/4/17 9:27
@Author : FB
@File : deleteImage.go
@Software: GoLand
*/
package image_service

import (
	"gocr/models"
	"gocr/resposity"
)

var repo *resposity.ImageRepo
var item_ *models.Image

func init() {
	repo = resposity.NewImageRepo()
	item_ = models.NewImage()
}

func CreateImage(item *models.Image) (err error) {
	err = repo.CreateImage(item)
	return
}

func GetImageList() (data []models.Image, err error) {
	data, err = repo.GetImageList()
	return
}

func UpdateImage(item *models.Image) (err error) {

	item_.XID = item.XID
	item_.Etc = item.Etc

	err = repo.UpdateImage(item_)

	return
}

func DeleteImage(xid string) (err error) {

	itemDel := models.NewImage()
	itemDel.XID = xid

	err = repo.DeleteImage(itemDel)

	return
}
