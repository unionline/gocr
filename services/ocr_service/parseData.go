/*
@Time : 2020/4/14 16:30
@Author : FB
@File : parseData.go
@Software: GoLand
*/
package ocr_service

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"gocr/entities"
	"gocr/env"
	"gocr/models"
	"gocr/services/ocr_log_service"
	"strings"
)

func parseToResult(input string) (data entities.Result, err error) {

	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	json.Unmarshal([]byte(input), &data)

	return
}

func parseData(data entities.Result, type_data int) (ret string, err error) {

	if type_data == env.TYPE_OCR_GENERAL_BASIC {
		ret = getWords(data)
	} else {
		fmt.Println("unsupport")
	}

	return
}

func getWords(input entities.Result) (ret string) {

	words_arr := make([]string, input.Words_result_num)
	for i, v := range input.Words_result {
		words_arr[i] = v.Words
	}

	ret = strings.Join(words_arr, "\n")

	return
}

func SaveLogOfOcr(item *models.OcrLog) (err error) {

	err = ocr_log_service.CreateOcrLog(item)
	return
}
