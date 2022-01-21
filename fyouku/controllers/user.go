package controllers

import (
	"fmt"
	"fyouku/models"
	"fyouku/utils"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

//用户相关页面
//登录页
func (this *UserController) Login() {
	fmt.Println(1)
	this.TplName = "login.html"
}

//登录页 mini
func (this *UserController) MiniLogin() {
	this.TplName = "mini_login.html"
}

//登录接口
func (this *UserController) LoginDo() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")
	var title string

	if mobile == "" {
		title = utils.ReturnError(5000, "手机号不能为空")
	}
	isorno, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if !isorno {
		title = utils.ReturnError(5000, "手机号格式不正确")
	}
	if password == "" {
		title = utils.ReturnError(5000, "密码不能为空")
	}

	if title == "" {
		title = models.IsMobileLogin(mobile, password)
	}
	this.Ctx.WriteString(title)
}

//注册
func (this *UserController) Register() {
	this.TplName = "register.html"
}

//注册接口
func (this *UserController) RegisterSave() {
	mobile := this.GetString("mobile")
	password := this.GetString("password")
	var title string

	if mobile == "" {
		title = utils.ReturnError(5000, "手机号不能为空")
	}
	isorno, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, mobile)
	if !isorno {
		title = utils.ReturnError(5000, "手机号格式不正确")
	}
	if password == "" {
		title = utils.ReturnError(5000, "密码不能为空")
	}

	if title == "" {
		title = models.UserSave(mobile, password)
	}
	this.Ctx.WriteString(title)
}

//个人中心-视频管理
func (this *UserController) UserVideo() {
	this.TplName = "ucenter_video.html"
}

//获取我的视频接口
func (this *UserController) GetMyVideos() {
	uid, _ := this.GetInt("uid")
	var title string

	if uid == 0 {
		title = utils.ReturnError(5000, "请先登录")
	}

	if title == "" {
		title = models.GetMyVideos(uid)
	}
	this.Ctx.WriteString(title)
}

//发送消息页面
func (this *UserController) SendMessage() {
	this.TplName = "sendMessage.html"
}

//发送消息接口
func (this *UserController) SendMessageDo() {
	uids := this.GetString("uids")
	content := this.GetString("content")
	var title string

	if uids == "" {
		title = utils.ReturnError(5000, "请设置接收人")
	}
	if content == "" {
		title = utils.ReturnError(5000, "请设置发送内容")
	}

	if title == "" {
		title = models.SendMessageDo(uids, content)
	}
	this.Ctx.WriteString(title)
}

//视频上传页面
func (this *UserController) Upload() {
	this.TplName = "ucenter_video_upload.html"
}

//视频基础信息保存
func (this *UserController) UploadInfoDo() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")
	aliyunVideoId := this.GetString("aliyunVideoId")

	var returnTitle string

	if uid == 0 {
		returnTitle = utils.ReturnError(5000, "请先登录")
	}
	if playUrl == "" {
		returnTitle = utils.ReturnError(5000, "视频地址必须为空")
	}

	if returnTitle == "" {
		returnTitle = models.SaveVideoInfo(uid, playUrl, title, subTitle, channelId, typeId, regionId, aliyunVideoId)
	}
	this.Ctx.WriteString(returnTitle)
}

//上传视频文件
func (this *UserController) UploadVideo() {
	var (
		err   error
		title string
	)
	r := *this.Ctx.Request
	//获取表单提交的数据
	uid := r.FormValue("uid")
	//获取文件流
	file, header, _ := r.FormFile("file")
	//转换文件流为二进制
	b, _ := ioutil.ReadAll(file)

	//生成文件名
	filename := strings.Split(header.Filename, ".")
	filename[0] = utils.GetVideoName(uid)
	//文件保存的位置
	var fileDir = "/Users/zhja/goRoot/src/fyouku/static/video/" + filename[0] + "." + filename[1]
	//播放地址
	var playUrl = "/static/video/" + filename[0] + "." + filename[1]
	err = ioutil.WriteFile(fileDir, b, 0777)
	if err == nil {
		title = utils.ReturnSuccess(0, playUrl, nil, 1)
	} else {
		title = utils.ReturnError(5000, "上传失败，请联系客服")
	}
	this.Ctx.WriteString(title)
}
