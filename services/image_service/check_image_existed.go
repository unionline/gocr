/*
@Time : 2020/4/21 16:04
@Author : FB
@File : check_image_existed
@Software: GoLand
*/
package image_service

import "gocr/models"

func CheckImageExisted(md5 string) (item models.Image, err error) {
	item, err = repo.GetImageInfoByMd5(md5)
	return
}
