/*
@Time : 2020/4/14 21:16
@Author : Justin
@Description :
@File : global.go
@Software: GoLand
*/
package env

// 全局变量，保存和请求
var (
	REQUEST_URL_RES = "/res/"
	SAVE_PAHT_RES   = "./views/resource/"
)

const (
	// 通用
	TYPE_OCR_GENERAL_BASIC  int = iota // 通用文字识别
	TYPE_OCR_GENERAL                   // 含位值
	TYPE_OCR_accurate_basic            // 高精度
	TYPE_OCR_accurate                  // 高精度含位值

	// 卡证
	TYPE_OCR_idcard_front
	TYPE_OCR_idcard_back
	TYPE_OCR_bankcard

	// 汽车场景
	TYPE_OCR_license_plate

	// 教育场景
	TYPE_OCR_handwriting

	// 其它文字
	// 网络图片文字
	TYPE_OCR_webimage
	TYPE_OCR_form
	TYPE_OCR_numbers
	TYPE_OCR_qrcode
)

var TypeMap = map[int]string{
	// 通用
	TYPE_OCR_GENERAL_BASIC:  "general_basic",  // 通用文字识别
	TYPE_OCR_GENERAL:        "general",        // 含位值
	TYPE_OCR_accurate_basic: "accurate_basic", // 高精度
	TYPE_OCR_accurate:       "accurate",       // 高精度含位值

	// 卡证
	TYPE_OCR_idcard_front: "idcard?id_card_side=front",
	TYPE_OCR_idcard_back:  "idcard?id_card_side=back",
	TYPE_OCR_bankcard:     "bankcard",

	// 汽车场景
	TYPE_OCR_license_plate: "license_plate",

	// 教育场景
	TYPE_OCR_handwriting: "handwriting",

	// 其它文字
	// 网络图片文字
	TYPE_OCR_webimage: "webimage",
	TYPE_OCR_form:     "form",
	TYPE_OCR_numbers:  "numbers",
}
