/*
@Time : 2020/3/20 14:42
@Author : FB
@File : custom.go
@Software: GoLand
*/
package utils

import "strings"

func ValidTemplateName(tmplName string) (err error) {

	seq := GetPathSeparator()
	//arr := strings.Split(tmplName, seq)
	//name := arr[len(arr)-1:]
	path_dir := strings.Join(
		[]string{
			"views",
			"template",
		}, seq)

	err = PathExists(path_dir + seq + tmplName)
	if err != nil {
		return
	}
	return
}
