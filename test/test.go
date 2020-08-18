package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gocr/utils"
	"net/url"
	"sync"
	"time"
	"unicode"
)

type NameMapper func(string) string

var (
	// 驼峰命名转换为大写加下划线链接 CamelCase —> CAMEL_CASE
	CamelCase NameMapper = func(raw string) string {
		newstr := make([]rune, 0, len(raw))
		for i, chr := range raw {
			if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
				if i > 0 {
					newstr = append(newstr, '_')
				}
			}
			newstr = append(newstr, unicode.ToUpper(chr))
		}
		return string(newstr)
	}
	// 下划线连接命名转换为小写加下划线链接 Title_Underscore —> title_underscore
	TitleUnderscore NameMapper = func(raw string) string {
		newstr := make([]rune, 0, len(raw))
		for i, chr := range raw {
			if isUpper := 'A' <= chr && chr <= 'Z'; isUpper {
				if i > 0 {
					newstr = append(newstr, '_')
				}
				chr -= 'A' - 'a'
			}
			newstr = append(newstr, chr)
		}
		return string(newstr)
	}
)

func Test() {
	fmt.Println("TES START")
	//jsonUtil()
	//filterArr()
	//testCrpyto()

	//a := ""
	//b:= "asd"
	//ok := ValidParameter(a,b)
	//fmt.Println(ok)

	//var ret string
	//var str string
	//str = "InitWaterFont"
	//ret = TitleUnderscore(str)
	//fmt.Println("ret=", ret)

	//var input = "token_filepath"
	//ret2 := CamelCase(input)
	//fmt.Println("ret2=", ret2)

	//testEscapeParamter()
	//TestSleep()
	QueueAdd("111", "222", "333")
	TestSleepChan()

	fmt.Println("TEST END")
}

// crpyto
func testCrpyto() {
	var output string

	//output = utils.AesEncrypt("qqq", "1234567890123456")
	output = utils.MD5Encode("1234567890123456")

	showVals("output", output)

}

func showVals(args ...interface{}) {

	for _, v := range args {
		fmt.Println(v)
	}

}

func testEscapeParamter() {
	str := "iVBORw0KGgoAAAANSUhEUgAAAGYAAAAUCAYAAAB/NUioAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAG8SURBVGhD7ZVtjsIgEIb3KG3iqRrjRTTxLqbxIsbewugVNut/VphSZqYDbaltcJcfT1JgmK8X6Nf3z1Nl0iMLkyhZmETJwiTKosLc6p0qyqO6CGvRXI/E5yIxEiALMwSLtxaf95St3ai0hGnUvtyo/RXPPdRpu1FV/WA2MF+8vjVuvT3N27O6dT7sCZftcYzLwdkUh8bZSDemizFivwFyt+s0BwfxoUG1RNXBehFitjBFuVOnO9hAsmyMkoEk8ekDH86nE7mLbYTg47Awwf33s6rwWKwVIdyYqDrsXO+QyMwWxo3dnN1HmtZrSAspXEqexR0jTGC/aSprjvHhaxgXJroObhPmDU8ZtgkI40vKFGpvGY8hzI0Qxr8fvvVJ7tH5YPC8o+vgNmH+pTB0fYC0hJEcQ9Pd3ERhfEmRQocay+2nCtM+Zb7bIcGFmFDHAk9ZvwD44c0Q5sXYn2aosXOFgcZSG+2D5IxzNPGoEGPr0P1y/eE2YbzCWEf2Da7qhhVNRZDmuDAaK7CF7l9BGI3xgWtzaz1hUJOlg2qR64CeWRuaV5iAMJl4pAMyjSzMImRhEiUL82fJwiTJU/0C4alQvTGYOO8AAAAASUVORK5CYII="
	output := url.QueryEscape(str)
	fmt.Println(output)

}

// add@20200325
func ValidParameter(args ...string) (valid bool) {

	for _, v := range args {
		if v == "" {
			return
		}
	}

	valid = true

	return
}

func jsonUtil() {

	src := `{"a":13,"b":         "nih*<>ao中文"}`
	dest := bytes.Buffer{}

	json.HTMLEscape(&dest, []byte(src))
	showVals(dest.String())

}

func filterArr() {
	component_name_need_arr := []string{
		"122",
		"22",
		"33",
	}

	component_name_arr_existed := []string{
		"122",
		"12",
	}

	// 如果重复去除，去除后为空，则返回已经创建
	var component_name_need_arr_filter []string
	for _, v := range component_name_need_arr {
		var flag = 0
		for _, vv := range component_name_arr_existed {
			if v == vv {
				break
			} else {
				flag++
			}
		}

		if flag == len(component_name_arr_existed) {
			component_name_need_arr_filter = append(component_name_need_arr_filter, v)
		}

	}
	fmt.Println("component_name_need_arr_filter=", component_name_need_arr_filter)
}

func TestSleep() {
	fmt.Println("begin", time.Now().Format("2006-01-02 03:04:05:06"))
	time.Sleep(time.Duration(5) * time.Second)
	// 测试使用简写也是可以达到效果 如果不是数字变量 则必须写上 time.Duration
	// time.Sleep(2 *time.Second)

	fmt.Println("end", time.Now().Format("2006-01-02 03:04:05:06"))
}

func TestSleepChan() {
	for i := 0; i < len(queue_imgs); i = i + 1 {
		time.Sleep(time.Duration(1) * time.Second)
		go func(n int) {
			defer wg.Done()
			fmt.Println("done", n, time.Now().Format("2006-01-02 03:04:05:06"))
		}(i)
	}

	wg.Wait()

}

func QueueAdd(imgs ...string) {
	for _, img := range imgs {
		queue_imgs = append(queue_imgs, img)
	}

	wg.Add(len(imgs))

}

var queue_imgs []string
var wg sync.WaitGroup
