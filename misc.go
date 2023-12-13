package xutil

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"time"
)

var rRand *rand.Rand

func init() {
	rRand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 获取len长度的随机字符串
func GetRandomString(len int) string {
	b := bytes.Buffer{}
	for i := 0; i < len; i++ {
		b.WriteByte(byte(rRand.Intn(256)))
	}
	return hex.EncodeToString(b.Bytes())
}

// 获取[min,max]范围随机整数
func GetRandomInt(min, max int) int {
	return rRand.Intn(max-min+1) + min
}

// 获取当前时间
func GetCurrentDateTime() string {
	return time.Now().Format("2006-03-02 15:04:05")
}
