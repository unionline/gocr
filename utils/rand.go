/*
@Time : 2020/4/14 23:51
@Author : Justin
@Description :
@File : rand
@Software: GoLand
*/
package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/rs/xid"
	"math/rand"
	"strconv"
	"time"
)

func GetSalt() string {
	rand.Seed(time.Now().UnixNano())
	num := rand.Int63n(1000000)
	num_str := strconv.FormatInt(num, 10)
	return num_str
}

// 获取一个XID
func GetXID() (id string) {
	guid := xid.New()
	id = guid.String()

	return
}

func GetRandSalt() string {

	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

}

func CreateMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) // 将[]byte转成16进制
	return md5str
}
