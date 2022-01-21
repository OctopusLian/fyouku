package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

func PageStart(page int, limit int) int {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}
	start := (page - 1) * limit
	return start
}

//字串截取
func SubString(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//用户注册MD5加密
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

type ReturnSuccessJson struct {
	//必须的大写开头
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

func ReturnSuccess(code int, msg string, items interface{}, count int64) string {
	jsonData := ReturnSuccessJson{Code: code, Msg: msg, Items: items, Count: count}
	if bytes, err := json.Marshal(jsonData); err == nil {
		return string(bytes)
	}
	return ""
}

type RetuenErrorJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func ReturnError(code int, err interface{}) string {
	var msg string
	switch err.(type) {
	case string:
		msg, _ = err.(string)
	default:
		msg = fmt.Sprintf("%s", err)
	}

	jsonData := RetuenErrorJson{Code: code, Msg: msg}
	if bytes, err := json.Marshal(jsonData); err == nil {
		return string(bytes)
	}
	return ""
}

func HttpGetApi(url string) string {
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//生成随记字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//视频文件名生成函数
func GetVideoName(uid string) string {
	//用户ID+精确到毫秒的时间戳
	h := md5.New()
	h.Write([]byte(uid + strconv.FormatInt(time.Now().UnixNano(), 10)))
	return hex.EncodeToString(h.Sum(nil))
}
