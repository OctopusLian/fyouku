package main

import (
	"context"
	"fmt"
	userProto "fyoukuapi/micro/user/proto"

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

	user := userProto.NewUserService("go.micro.srv.fyoukuApi.user", service.Client())

	rsp, err := user.LoginDo(context.TODO(), &userProto.RequestLoginDo{
		Mobile:   "18600001111",
		Password: "111111",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
