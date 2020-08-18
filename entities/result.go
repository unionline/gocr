/*
@Time : 2020/4/14 20:06
@Author : Justin
@Description :
@File : result.go
@Software: GoLand
*/
package entities

//{
//"log_id": 2471272194,
//"words_result_num": 2,
//"words_result":
//[
//{"words": " TSINGTAO"},
//{"words": "青島睥酒"}
//]
//}

type Result struct {
	LogId            int64   `json:"log_id"`
	ErrorCode        int     `json:"error_code"`
	ErrorMsg         string  `json:"error_msg"`
	Words_result_num int     `json:"words_result_num"`
	Words_result     []words `json:"words_result"`
}

type words struct {
	Words string `json:"words"`
}
