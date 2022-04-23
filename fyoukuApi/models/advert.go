/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-04-23 12:43:49
 * @LastEditors: neozhang
 * @LastEditTime: 2022-04-23 19:47:07
 */
package models

import (
	"github.com/astaxie/beego/orm"
)

//广告
type Advert struct {
	Id        int
	Title     string `gorm:"comment:'广告标题'"`
	SubTitle  string `gorm:"comment:'广告副标题'"`
	AddTime   int64  `gorm:"comment:'添加时间'"`
	Img       string `gorm:"comment:'广告图片'"`
	Url       string `gorm:"comment:'跳转地址'"`
	ChannelId int    `gorm:"comment:'所属频道'"`
	Sort      string `gorm:"comment:'排序'"`
	Status    bool   `gorm:"comment:'0=关闭 1=开启'"`
}

func init() {
	orm.RegisterModel(new(Advert))
}

func GetChannelAdvert(channelId int) (int64, []Advert, error) {
	o := orm.NewOrm()
	var adverts []Advert
	num, err := o.Raw("SELECT id, title, sub_title,img,add_time,url FROM advert WHERE status=1 AND channel_id=? ORDER BY sort DESC LIMIT 1", channelId).QueryRows(&adverts)
	return num, adverts, err
}
