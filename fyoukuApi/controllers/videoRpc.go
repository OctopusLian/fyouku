package controllers

import (
	"context"
	videoRpcProto "fyoukuApi/micro/video/proto"
	"fyoukuApi/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type VideoRpcController struct {
	beego.Controller
}

func (this *VideoRpcController) ChannelAdvert(ctx context.Context, req *videoRpcProto.RequestChannelAdvert, res *videoRpcProto.ResponseChannelAdvert) error {
	var (
		channelAdvertDatas []*videoRpcProto.ChannelAdvertData
		num                int64
		videos             []models.Advert
		err                error
	)

	channelId, _ := strconv.Atoi(req.ChannelId)

	if channelId == 0 {
		res.Code = 4001
		res.Msg = "必须指定频道"
		goto ERR
	}
	num, videos, err = models.GetChannelAdvert(channelId)

	if err == nil {
		for _, v := range videos {
			var channelAdvertData videoRpcProto.ChannelAdvertData
			channelAdvertData.Id = int64(v.Id)
			channelAdvertData.Title = v.Title
			channelAdvertData.SubTitle = v.SubTitle
			channelAdvertData.Img = v.Img
			channelAdvertData.Url = v.Url
			channelAdvertData.AddTime = v.AddTime
			channelAdvertDatas = append(channelAdvertDatas, &channelAdvertData)
		}
		res.Code = 0
		res.Msg = "success"
		res.Items = channelAdvertDatas
		res.Count = num
		return nil
	} else {
		res.Code = 4004
		res.Msg = "请求数据失败，请稍后重试~"
		goto ERR
	}
ERR:
	res.Items = channelAdvertDatas
	res.Count = 0
	return nil
}

func (this *VideoRpcController) ChannelHotList(ctx context.Context, req *videoRpcProto.RequestChannelHotList, res *videoRpcProto.ResponseChannelHotList) error {
	var (
		channelHotListDatas []*videoRpcProto.ChannelHotListData
		num                 int64
		videos              []models.VideoData
		err                 error
	)

	channelId, _ := strconv.Atoi(req.ChannelId)

	if channelId == 0 {
		res.Code = 4001
		res.Msg = "必须指定频道"
		goto ERR
	}
	num, videos, err = models.GetChannelHotList(channelId)

	if err == nil {
		for _, v := range videos {
			var channelHotListData videoRpcProto.ChannelHotListData
			channelHotListData.Id = int64(v.Id)
			channelHotListData.Title = v.Title
			channelHotListData.SubTitle = v.SubTitle
			channelHotListData.Img = v.Img
			channelHotListData.Img1 = v.Img1
			channelHotListData.IsEnd = int64(v.IsEnd)
			channelHotListData.AddTime = v.AddTime
			channelHotListData.EpisodesCount = int64(v.EpisodesCount)
			channelHotListData.Comment = int64(v.Comment)
			channelHotListDatas = append(channelHotListDatas, &channelHotListData)
		}

		res.Code = 0
		res.Msg = "success"
		res.Items = channelHotListDatas
		res.Count = num
		return nil
	} else {
		res.Code = 4004
		res.Msg = "没有相关内容"
		goto ERR
	}
ERR:
	res.Items = channelHotListDatas
	res.Count = 0
	return nil
}
