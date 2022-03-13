package controllers

import (
	"fmt"
	"fyouku/models"
	"fyouku/utils"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router /index [get]
func (this *MainController) Get() {
	fmt.Println(1)
	//前后端分离
	//获取顶部视频推荐广告，调用接口
	advertVideos := models.GetChannelAdvert(1)
	this.Data["advertVideos"] = advertVideos
	//获取正在热播，调用接口
	hotVideos := models.GetChannelHotList(1)
	this.Data["hotVideos"] = hotVideos
	//获取日漫、国漫推荐，调用接口
	japanVideos := models.GetChannelRegionRecommend(1, 1)
	chinaVideos := models.GetChannelRegionRecommend(1, 2)
	this.Data["japanVideos"] = japanVideos
	this.Data["chinaVideos"] = chinaVideos
	//获取少女推荐，调用接口
	girlVideos := models.GetChannelTypeRecommend(1, 1)
	this.Data["girlVideos"] = girlVideos

	//获取动漫排行榜
	comicTop := models.GetChannelTop(1)
	this.Data["comicTop"] = comicTop
	//获取少女排行榜
	girlTop := models.GetTypeTop(1)
	this.Data["girlTop"] = girlTop

	this.TplName = "index.html"
}

//更多页面
//频道页
func (this *MainController) Channel() {
	//接受浏览器参数
	//地区
	regionId, _ := this.GetInt("regionId")
	this.Data["regionId"] = regionId
	//类型
	typeId, _ := this.GetInt("typeId")
	this.Data["typeId"] = typeId
	//状态
	end := this.GetString("end")
	this.Data["end"] = end
	//排序
	sort := this.GetString("sort")
	if sort == "" {
		sort = "episodesUpdateTime"
	}
	this.Data["sort"] = sort
	//获得地区列表
	regions := models.GetChannelRegion(1)
	this.Data["regions"] = regions
	//获得类型列表
	types := models.GetChannelType(1)
	this.Data["types"] = types

	//根据接受到的参数获取视频信息
	videos, _ := models.GetChannelVideoList(1, regionId, typeId, end, sort, 1)
	this.Data["videos"] = videos
	this.TplName = "channel.html"
}

//频道JS加载数据接口
func (this *MainController) ChannelVideoData() {
	//接受浏览器参数
	//地区
	regionId, _ := this.GetInt("regionId")
	//类型
	typeId, _ := this.GetInt("typeId")
	//状态
	end := this.GetString("end")
	//排序
	sort := this.GetString("sort")
	if sort == "" {
		sort = "episodesUpdateTime"
	}
	//页数
	page, _ := this.GetInt("page")

	//根据接受到的参数获取视频信息
	videos, count := models.GetChannelVideoList(1, regionId, typeId, end, sort, page)
	title := utils.ReturnSuccess(0, "success", videos, count)
	this.Ctx.WriteString(title)
}

//搜索页
func (this *MainController) Search() {
	keyword := this.GetString("keyword")
	this.Data["keyword"] = keyword
	this.TplName = "search.html"
}

//搜索接口
func (this *MainController) SearchData() {
	//接受浏览器参数
	//搜索关键词
	keyword := this.GetString("keyword")
	//页数
	page, _ := this.GetInt("page")

	//根据接受到的参数获取视频信息
	videos, count := models.GetSearchVideoList(keyword, page)
	title := utils.ReturnSuccess(0, "success", videos, count)
	this.Ctx.WriteString(title)
}

//排行榜
func (this *MainController) Top() {
	//获取动漫排行榜
	comicTop := models.GetChannelTop(1)
	this.Data["comicTop"] = comicTop
	//获取少女排行榜
	girlTop := models.GetTypeTop(1)
	this.Data["girlTop"] = girlTop

	this.TplName = "top.html"
}

//详情页
func (this *MainController) Show() {
	//获取视频ID
	videoId, _ := this.GetInt("id")
	this.Data["videoId"] = videoId
	episodesId, _ := this.GetInt("episodesId")
	//获取视频信息
	videoInfo := models.GetVideoInfo(videoId)
	this.Data["videoInfo"] = videoInfo
	//获取视频集数列表
	episodesList := models.GetVideoEpisodesList(videoId)
	this.Data["episodesList"] = episodesList

	var episodes models.Episodes
	if len(episodesList) > 0 {
		for i, v := range episodesList {
			if episodesId != 0 && episodesId == v.Id {
				episodes.Id = v.Id
				episodes.Title = v.Title
				episodes.AddTime = v.AddTime
				episodes.Num = v.Num
				episodes.PlayUrl = v.PlayUrl
				episodes.Comment = v.Comment
				episodes.AliyunVideoId = v.AliyunVideoId
			} else if episodesId == 0 && i == 0 {
				episodes.Id = v.Id
				episodes.Title = v.Title
				episodes.AddTime = v.AddTime
				episodes.Num = v.Num
				episodes.PlayUrl = v.PlayUrl
				episodes.Comment = v.Comment
				episodes.AliyunVideoId = v.AliyunVideoId
			}
		}
	}
	this.Data["episodes"] = episodes

	//获取播放凭证
	if episodes.AliyunVideoId != "" {
		this.Data["PlayAuth"] = models.GetPlayAuth(episodes.AliyunVideoId)
	} else {
		this.Data["PlayAuth"] = ""
	}

	//获取国漫推荐
	chinaVideos := models.GetChannelRegionRecommend(1, 2)
	this.Data["chinaVideos"] = chinaVideos
	//获取日漫推荐
	japanVideos := models.GetChannelRegionRecommend(1, 1)
	this.Data["japanVideos"] = japanVideos
	//获取少女推荐
	girlVideos := models.GetChannelTypeRecommend(1, 1)
	this.Data["girlVideos"] = girlVideos

	this.TplName = "show.html"
}

//获取评论信息
func (this *MainController) GetCommentList() {
	//集数ID
	episodesId, _ := this.GetInt("episodesId")
	//获取页码
	page, _ := this.GetInt("page")

	//获取评论ID
	title := models.GetCommentList(episodesId, page)
	this.Ctx.WriteString(title)
}

//保存评论
func (this *MainController) SaveComment() {
	//集数ID
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	//获取用户ID
	uid, _ := this.GetInt("uid")
	//内容
	content := this.GetString("content")

	//获取评论ID
	title := models.SaveComment(content, uid, episodesId, videoId)
	this.Ctx.WriteString(title)
}

//保存弹幕
func (this *MainController) SaveBarrage() {
	//集数ID
	episodesId, _ := this.GetInt("episodesId")
	videoId, _ := this.GetInt("videoId")
	//获取用户ID
	uid, _ := this.GetInt("uid")
	//内容
	content := this.GetString("content")
	//视频当前播放时间
	currentTime, _ := this.GetInt("currentTime")

	//获取评论ID
	title := models.SaveBarrage(content, currentTime, uid, episodesId, videoId)
	this.Ctx.WriteString(title)
}
