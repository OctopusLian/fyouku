package models

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type PlayAuthJSONS struct {
	PlayAuth string
}

//获取上传凭证
func CreateUploadVideo(title string, desc string, fileName string, coverUrl string, tags string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/aliyun/create/upload/video")
	req.Param("title", title)
	req.Param("desc", desc)
	req.Param("fileName", fileName)
	req.Param("coverUrl", coverUrl)
	req.Param("tags", tags)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//刷新上传凭证
func RefreshUploadVideo(videoId string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/aliyun/refresh/upload/video")
	req.Param("videoId", videoId)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//获取播放凭证
func GetPlayAuth(videoId string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/aliyun/video/play/auth")
	req.Param("videoId", videoId)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	stb := &PlayAuthJSONS{}
	err = json.Unmarshal([]byte(str), &stb)
	if err != nil {
		fmt.Println(err)
	}

	return stb.PlayAuth
}
