/*
 * @Author: zhangniannian
 * @Date: 2022-01-21 18:05:35
 * @LastEditors: zhangniannian
 * @LastEditTime: 2022-01-21 18:10:58
 * @Description: 请填写简介
 */
package main

import (
	"context"
	"fmt"
	videoProto "fyoukuapi/micro/video/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Registry(reg),
	)

	service.Init()

	video := videoProto.NewVideoService("go.micro.srv.fyoukuApi.video", service.Client())

	rsp, err := video.ChannelAdvert(context.TODO(), &videoProto.RequestChannelAdvert{
		ChannelId: "1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)

	rspHot, _ := video.ChannelHotList(context.TODO(), &videoProto.RequestChannelHotList{
		ChannelId: "1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rspHot)
}
