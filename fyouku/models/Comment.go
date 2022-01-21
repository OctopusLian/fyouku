package models

import (
	"fmt"
	"fyouku/utils"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//获取评论信息
func GetCommentList(episodesId int, page int) string {
	url := beego.AppConfig.String("apiurl") + "/comment/list?episodesId=" + strconv.Itoa(episodesId)

	//获取数据条数和起始位置
	limit := 1
	if page == 0 {
		page = 1
	}
	var offset = (page - 1) * limit
	url += "&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)
	fmt.Println(url)
	req := utils.HttpGetApi(url)
	//fmt.Println(req)

	return req
}

//保存评论
func SaveComment(content string, uid int, episodesId int, videoId int) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/comment/save")
	req.Param("content", content)
	req.Param("uid", strconv.Itoa(uid))
	req.Param("episodesId", strconv.Itoa(episodesId))
	req.Param("videoId", strconv.Itoa(videoId))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
