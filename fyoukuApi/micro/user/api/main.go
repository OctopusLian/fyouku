package main

import (
	"context"
	"encoding/json"
	user "fyoukuApi/micro/user/proto"
	"log"
	"strings"
	"time"

	"github.com/micro/go-micro"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type User struct {
	Client user.UserService
}

func (u *User) LoginDo(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("收到User.LoginDo API请求")
	//接受参数
	mobile, ok := req.Post["mobile"]
	if !ok || len(mobile.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.user", "mobile为空")
	}
	password, ok := req.Post["password"]
	if !ok || len(password.Values) == 0 {
		return errors.BadRequest("go.micro.api.fyoukuApi.user", "password为空")
	}

	response, err := u.Client.LoginDo(ctx, &user.RequestLoginDo{
		Mobile:   strings.Join(mobile.Values, ""),
		Password: strings.Join(password.Values, ""),
	})
	if err != nil {
		return err
	}
	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]interface{}{
		"code":  response.Code,
		"msg":   response.Msg,
		"items": response.Items,
		"count": response.Count,
	})
	rsp.Body = string(b)
	return nil
}

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.api.fyoukuApi.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	service.Server().Handle(
		service.Server().NewHandler(
			&User{Client: user.NewUserService("go.micro.srv.fyoukuApi.user", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
