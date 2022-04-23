package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type UserInfo struct {
	Id      int
	Name    string
	AddTime int64
	Avatar  string
}

//判断用户名密码是否正确
func IsMobileLogin(mobile string, password string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/login/do")
	//req := httplib.Post(beego.AppConfig.String("microApi") + "/fyoukuapi/user/user/LoginDo")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//保存用户
func UserSave(mobile string, password string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/register/save")
	req.Param("mobile", mobile)
	req.Param("password", password)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//保存用户
func SendMessageDo(uids string, content string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/send/message")
	req.Param("uids", uids)
	req.Param("content", content)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
