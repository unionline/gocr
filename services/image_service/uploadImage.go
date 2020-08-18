/*
@Time : 2020/4/14 16:03
@Author : FB
@File : loadImage.go
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

func UploadImage(data []byte, header *multipart.FileHeader) (img_path string, err error) {

	item := models.Image{}

	images_dir := config.Setting.Path.ImagesUploadDir

	md5_str := utils.MD5Encode(string(data))
	item, err = CheckImageExisted(md5_str)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}

	if item.ID > 0 {
		img_path = images_dir + item.FileName
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
	item.Md5 = md5_str
	item.From = "gocr"

	img_path = images_dir + item.FileName
	err = utils.WriterFileConent(img_path, data)
	if err != nil {
		return
	}

	err = CreateImage(&item)
	if err != nil {
		return
	}

	return
}
