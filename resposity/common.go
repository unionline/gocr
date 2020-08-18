/*
@Time : 2020/4/14 23:00
@Author : Justin
@Description :
@File : common.go
@Software: GoLand
*/
package resposity

import (
	"gocr/models/crud"
)

var (
	Insert *crud.Insert
	Query  *crud.Read
	Update *crud.Update
	Delete *crud.Delete
)
