/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-23 12:43:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-23 19:46:19
 */
package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Barrage struct {
	Id          int
	Content     string `gorm:"comment:'姓名'"`
	CurrentTime int    `gorm:"comment:'视频当前播放时间'"`
	AddTime     int64  `gorm:"comment:'评论时间'"`
	UserId      int    `gorm:"comment:'评论用户id'"`
	Status      int    `gorm:"comment:'状态0=未审核 1=审核通过'"`
	EpisodesId  int    `gorm:"comment:'评论视频'"`
	VideoId     int    `gorm:"comment:'所属视频'"`
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

func BarrageList(episodesId int, startTime int, endTime int) (int64, []BarrageData, error) {
	o := orm.NewOrm()
	var barrages []BarrageData
	num, err := o.Raw("SELECT id,content,`current_time` FROM barrage WHERE status=1 AND episodes_id=? AND `current_time`>=? AND `current_time`<? ORDER BY `current_time` ASC", episodesId, startTime, endTime).QueryRows(&barrages)
	return num, barrages, err
}

func SaveBarrage(episodesId int, videoId int, currentTime int, userId int, content string) error {
	o := orm.NewOrm()
	var barrage Barrage
	barrage.Content = content
	barrage.CurrentTime = currentTime
	barrage.AddTime = time.Now().Unix()
	barrage.UserId = userId
	barrage.Status = 1
	barrage.EpisodesId = episodesId
	barrage.VideoId = videoId
	_, err := o.Insert(&barrage)
	return err
}
