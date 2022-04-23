package main

import (
	"fmt"
	"fyoukuApi/controllers"
	"fyoukuApi/micro/user/proto"
	_ "fyoukuApi/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	beego.LoadAppConfig("ini", "../../conf/app.conf")
	defaultdb := beego.AppConfig.String("defaultdb")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", defaultdb, 30, 30)

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"http://127.0.0.1:2379"}
	})

	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.micro.srv.fyoukuApi.user"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	proto.RegisterUserServiceHandler(service.Server(), new(controllers.UserRpcController))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
