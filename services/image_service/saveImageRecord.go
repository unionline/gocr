/*
@Time : 2020/4/21 22:59
@Author : Justin
@Description :
@File : saveImageRecord
@Software: GoLand
*/
package image_service

import (
	"github.com/jinzhu/gorm"
	"gocr/initializers/config"
	"gocr/models"
	"gocr/utils"
	"mime/multipart"
)

func SaveImageRecord(data []byte, header *multipart.FileHeader) (existed bool, img_path string, err error) {

	item := models.Image{}

	images_dir := config.Setting.Path.ImagesUploadDir

	md5_str := utils.MD5Encode(string(data))
	item, err = CheckImageExisted(md5_str)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if item.ID > 0 {
		img_path = images_dir + item.FileName
		existed = true
		return
	}

	filename := header.Filename
	size := header.Size

	item.Name = filename
	item.XID = utils.GetXID()
	item.Size = utils.ConvertHumanUnit(float64(size))

	newfilename := item.XID
	item.FileName = newfilename + utils.GetFileNameSuffix(filename)
	item.ThumbnailName = utils.AddFileNameSuffix(newfilename, "png")
	item.From = "gocr"
	item.Md5 = md5_str

	img_path = images_dir + item.FileName

	err = CreateImage(&item)
	if err != nil {
		return
	}

	return
}
