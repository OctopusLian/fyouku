package controllers

import (
	"fyouku/models"

	"github.com/astaxie/beego"
)

type AliyunController struct {
	beego.Controller
}

//获取上传凭证
func (this *AliyunController) CreateUploadVideo() {
	title := this.GetString("title")
	desc := this.GetString("desc")
	fileName := this.GetString("fileName")
	coverUrl := this.GetString("coverUrl")
	tags := this.GetString("tags")
	req := models.CreateUploadVideo(title, desc, fileName, coverUrl, tags)
	this.Ctx.WriteString(req)
}

//刷新上传凭证
func (this *AliyunController) RefreshUploadVideo() {
	videoId := this.GetString("videoId")

	req := models.RefreshUploadVideo(videoId)
	this.Ctx.WriteString(req)
}
