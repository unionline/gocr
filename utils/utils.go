/*
@Time : 2020/3/19 16:55
@Author : FB
@File : utils.go
@Software: GoLand
*/
package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func PathExists(path string) (err error) {
	_, err = os.Stat(path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		return
	}
	return
}

// 判定目录是否存在，不存在，则创建，否则不创建
func AutoCreateDir(path string) (ok bool) {
	err := PathExists(path)
	if err != nil {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			ok = false
			return
		} else {
			ok = true
			return
		}
	}

	ok = true
	return

}

func GetSystem() string {
	return runtime.GOOS
}

func IsLinuxPlatform() bool {
	return strings.EqualFold(GetSystem(), "linux")
}

func IsWindowsPlatform() bool {
	return strings.EqualFold(GetSystem(), "windows")
}

func GetPathSeparator() string {

	return string(os.PathSeparator)

}

// 在指定路径创建文件,并写入内容
func WriterFileConent(path string, data []byte) error {
	dir, _ := filepath.Split(path)
	DirIsExist(dir)
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	} else {
		_, err = f.Write(data)
		return err
	}
	return err
}

// 判断目录是否存在,不存在则新建一个
func DirIsExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return CreateDir(path)
	}
	return err
}

// 新建目录
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetFileNameSuffix(filename string) string {

	pos := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename
	}
	return filename[pos:]
}

func AddFileNameSuffix(filename string, format string) string {

	pos := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename + format
	}
	if strings.Index(format, ".") == 0 {
		format = format[1:len(format)]
	}

	return filename[:pos+1] + format
}

func GetImageId(filename string) string {
	// 如果不是图片格式，根本不可能上传成功
	// ./views/resource/upload/images/abc123md5123.png -> abc123md5123
	pos := strings.LastIndex(filename, "/")
	pos2 := strings.LastIndex(filename, ".")
	if pos < 1 {
		return filename[:pos2]
	}

	return filename[pos+1 : pos2]
}
