package models

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type Video struct {
	Id            int
	Title         string
	SubTitle      string
	AddTime       int64
	Img           string
	Img1          string
	EpisodesCount int
	IsEnd         int
}

type VideoApiData struct {
	Code  int
	Msg   string
	Items []Video
	Count int
}

type VideoInfo struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Img                string
	Img1               string
	EpisodesCount      int
	IsEnd              int
	ChannelId          int
	Status             int
	RegionId           int
	TypeId             int
	Sort               string
	EpisodesUpdateTime int
	Comment            int
}
type VideoInfoData struct {
	Code  int
	Msg   string
	Items VideoInfo
	Count int
}

type Episodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	PlayUrl       string
	Comment       int
	AliyunVideoId string
}
type EpisodesApiData struct {
	Code  int
	Msg   string
	Items []Episodes
	Count int
}

//获取正在热播
func GetChannelHotList(channelId int) []Video {
	var info []Video
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/channel/hot?channelId=" + strconv.Itoa(channelId))
<<<<<<< HEAD
	//req := httplib.Get(beego.AppConfig.String("microApi") + "/fyoukuapi/video/video/ChannelHotList?channelId=" + strconv.Itoa(channelId))
=======
	//req := httplib.Get(beego.AppConfig.String("microApi") + "/fyoukuApi/video/video/ChannelHotList?channelId=" + strconv.Itoa(channelId))
>>>>>>> master
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

//获取少女推荐
func GetChannelTypeRecommend(channelId int, typeId int) []Video {
	var info []Video
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/channel/recommend/type?channelId=" + strconv.Itoa(channelId) + "&typeId=" + strconv.Itoa(typeId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

//获取日漫、国漫推荐
func GetChannelRegionRecommend(channelId int, regionId int) []Video {
	var info []Video
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/channel/recommend/region?channelId=" + strconv.Itoa(channelId) + "&regionId=" + strconv.Itoa(regionId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

//频道页，根据接收到的参数获取视频信息
func GetChannelVideoList(channelId int, regionId int, typeId int, end string, sort string, page int) ([]Video, int64) {
	var info []Video
	var count int64
	url := beego.AppConfig.String("apiurl") + "/channel/video?channelId=" + strconv.Itoa(channelId)
	if regionId > 0 {
		url += "&regionId=" + strconv.Itoa(regionId)
	}
	if typeId > 0 {
		url += "&typeId=" + strconv.Itoa(typeId)
	}
	if end != "" {
		url += "&end=" + end
	}
	if sort != "" {
		url += "&sort=" + sort
	}
	//获取数据条数和起始位置
	limit := 12
	var offset = (page - 1) * limit
	url += "&offset=" + strconv.Itoa(offset) + "&limit=" + strconv.Itoa(limit)
	//fmt.Println(url)
	req := httplib.Get(url)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
		countString := strconv.Itoa(stb.Count)
		count, _ = strconv.ParseInt(countString, 10, 64)
	}

	return info, count
}

//获取搜索结果数据
func GetSearchVideoList(keyword string, page int) ([]Video, int64) {
	var info []Video
	var count int64
	limit := 12
	var offset = (page - 1) * limit

	req := httplib.Post(beego.AppConfig.String("apiurl") + "/video/search")
	req.Param("keyword", keyword)
	req.Param("offset", strconv.Itoa(offset))
	req.Param("limit", strconv.Itoa(limit))
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoApiData{}
	err = json.Unmarshal([]byte(str), &stb)
	fmt.Println(stb)
	if stb.Code == 0 {
		info = stb.Items
		countString := strconv.Itoa(stb.Count)
		count, _ = strconv.ParseInt(countString, 10, 64)
	}

	return info, count
}

//获取我的视频
func GetMyVideos(uid int) string {
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/user/video?uid=" + strconv.Itoa(uid) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}

//获取视频详情
func GetVideoInfo(videoId int) VideoInfo {
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/video/info?videoId=" + strconv.Itoa(videoId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &VideoInfoData{}
	err = json.Unmarshal([]byte(str), &stb)

	var videoInfo VideoInfo
	if stb.Code == 0 {
		videoInfo = stb.Items
	}

	return videoInfo
}

//获取集数列表
func GetVideoEpisodesList(videoId int) []Episodes {
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/video/episodes/list?videoId=" + strconv.Itoa(videoId) + "")
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &EpisodesApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	var episodesData []Episodes
	if stb.Code == 0 {
		episodesData = stb.Items
	}

	return episodesData
}

//视频广告相关
type Advert struct {
	Id       int
	Title    string
	SubTitle string
	AddTime  int64
	Img      string
	Url      string
}

type AdvertApiData struct {
	Code  int
	Msg   string
	Items []Advert
	Count string
}

func GetChannelAdvert(channelId int) []Advert {
	var info []Advert
	req := httplib.Get(beego.AppConfig.String("apiurl") + "/channel/advert?channelId=" + strconv.Itoa(channelId) + "")
<<<<<<< HEAD
	//req := httplib.Get(beego.AppConfig.String("microApi") + "/fyoukuapi/video/video/ChannelAdvert?channelId=" + strconv.Itoa(channelId) + "")
=======
	//req := httplib.Get(beego.AppConfig.String("microApi") + "/fyoukuApi/video/video/ChannelAdvert?channelId=" + strconv.Itoa(channelId) + "")
>>>>>>> master
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	stb := &AdvertApiData{}
	err = json.Unmarshal([]byte(str), &stb)

	if stb.Code == 0 {
		info = stb.Items
	}

	return info
}

//保存视频
func SaveVideoInfo(uid int, playUrl string, title string, subTitle string, channelId int, typeId int, regionId int, aliyunVideoId string) string {
	req := httplib.Post(beego.AppConfig.String("apiurl") + "/video/save")
	req.Param("playUrl", playUrl)
	req.Param("title", title)
	req.Param("subTitle", subTitle)
	req.Param("uid", strconv.Itoa(uid))
	req.Param("channelId", strconv.Itoa(channelId))
	req.Param("typeId", strconv.Itoa(typeId))
	req.Param("regionId", strconv.Itoa(regionId))
	req.Param("aliyunVideoId", aliyunVideoId)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	return str
}
