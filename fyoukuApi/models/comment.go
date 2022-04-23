/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-23 12:43:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-23 19:48:10
 */
package models

import (
	"encoding/json"
	"fyoukuApi/services/mq"
	"time"

	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id          int
	Content     string `gorm:"comment:'评论内容'"`
	AddTime     int64  `gorm:"comment:'评论时间'"`
	UserId      int    `gorm:"comment:'评论用户'"`
	Stamp       int    `gorm:"comment:'状态0=未审核 1=审核通过'"`
	Status      int    `gorm:"comment:'盖章1=热评2=公告'"`
	PraiseCount int    `gorm:"comment:'点赞数'"`
	EpisodesId  int    `gorm:"comment:'评论视频'"`
	VideoId     int    `gorm:"comment:'所属视频'"`
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId int, offset int, limit int) (int64, []Comment, error) {
	o := orm.NewOrm()
	var comments []Comment
	num, _ := o.Raw("SELECT id FROM comment WHERE status=1 AND episodes_id=?", episodesId).QueryRows(&comments)
	_, err := o.Raw("SELECT id,content,add_time,user_id,stamp,praise_count,episodes_id FROM comment WHERE status=1 AND episodes_id=? ORDER BY add_time DESC LIMIT ?,?", episodesId, offset, limit).QueryRows(&comments)
	return num, comments, err
}

func SaveComment(content string, uid int, episodesId int, videoId int) error {
	o := orm.NewOrm()
	var comment Comment
	comment.Content = content
	comment.UserId = uid
	comment.EpisodesId = episodesId
	comment.VideoId = videoId
	comment.Stamp = 0
	comment.Status = 1
	comment.AddTime = time.Now().Unix()
	_, err := o.Insert(&comment)
	if err == nil {
		//修改视频的总评论数
		o.Raw("UPDATE video SET comment=comment+1 WHERE id=?", videoId).Exec()
		//修改视频剧集的评论数
		o.Raw("UPDATE video_episodes SET comment=comment+1 WHERE id=?", episodesId).Exec()

		//更新redis排行榜 - 通过MQ来实现
		//创建一个简单模式的MQ
		//把要传递的数据转换为json字符串
		videoObj := map[string]int{
			"VideoId": videoId,
		}
		videoJson, _ := json.Marshal(videoObj)
		mq.Publish("", "fyouku_top", string(videoJson))

		//延迟增加评论数
		videoCountObj := map[string]int{
			"VideoId":    videoId,
			"EpisodesId": episodesId,
		}
		videoCountJson, _ := json.Marshal(videoCountObj)
		mq.PublishDlx("fyouku.comment.count", string(videoCountJson))
	}
	return err
}
