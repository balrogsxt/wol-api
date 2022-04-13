package helper

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

func IsEmpty(data interface{}) bool {
	s := strings.Trim(fmt.Sprintf("%v",data)," ")
	if len(s) == 0 {
		return true
	}
	return false
}

func GetRandomText(size int,s string) string {
	if len(s) == 0 {
		s = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	arr := strings.Split(s,"")
	str := ""
	for i:=0;i<size;i++{
		str +=arr[rand.Intn(len(s))]
	}
	return str
}
func GetRandomNumber(size int) string {
	return GetRandomText(size,"0123456789")
}

func Md5String(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}