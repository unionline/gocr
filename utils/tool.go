package utils

import (
	"strconv"
	"strings"
)

func ConvertHumanUnit(num float64) string {

	var ret string
	if num >= 1024*1024 {
		ret = Float2Str(num/1024, 2) + " MB"
	} else if num >= 1024 {
		ret = Float2Str(num/1024, 2) + " KB"
	} else if num < 1024 {
		ret = Float2Str(num/1024, 2) + " Byte"
	} else {
		return "1 GB"
	}

	return ret
}

func Str2Float(str string) (f float64) {
	f, _ = strconv.ParseFloat(str, 64)
	return
}

func Float2Str(f float64, prec int) (str string) {
	str = strconv.FormatFloat(f, 'f', prec, 64)
	return
}

func Str2FloatPrec(strFloat string, prec int) (str string) {

	if prec == -1 {
		return Float2Str(Str2Float(strFloat), -1)
	}

	pos := strings.Index(strFloat, ".")
	pos2 := pos + prec + 1
	len_str := len(strFloat)

	if pos == -1 {
		str = strFloat
	} else if pos2 <= len_str {
		str = strFloat[:pos+prec+1]
	} else {
		str = strFloat[:len_str]
	}

	return
}
