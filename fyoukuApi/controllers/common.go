package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/astaxie/beego"
)

// Operations about Users
type CommonController struct {
	beego.Controller
}

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Items interface{} `json:"items"`
	Count int64       `json:"count"`
}

//返回成功的格式
func ReturnSuccess(code int, msg interface{}, items interface{}, count int64) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg, Items: items, Count: count}
	return
}

//返回失败的格式
func ReturnError(code int, msg interface{}) (json *JsonStruct) {
	json = &JsonStruct{Code: code, Msg: msg}
	return
}

//用户密码加密
func MD5V(password string) string {
	h := md5.New()
	h.Write([]byte(password + beego.AppConfig.String("md5code")))
	return hex.EncodeToString(h.Sum(nil))
}

//格式化时间
func DateFormat(times int64) string {
	video_time := time.Unix(times, 0)
	return video_time.Format("2006-01-02")
}
