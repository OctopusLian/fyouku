package models

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//保存弹幕
func SaveBarrage(content string, currentTime int, uid int, episodesId int, videoId int) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/barrage/save")
	req.Param("content", content)
	req.Param("currentTime", strconv.Itoa(currentTime))
	req.Param("uid", strconv.Itoa(uid))
	req.Param("episodesId", strconv.Itoa(episodesId))
	req.Param("videoId", strconv.Itoa(videoId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
