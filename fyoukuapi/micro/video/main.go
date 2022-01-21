package main

import (
	"fmt"
	"fyoukuapi/controllers"
	"fyoukuapi/micro/video/proto"
	_ "fyoukuapi/routers"
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
		micro.Name("go.micro.srv.fyoukuApi.video"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	service.Init()
	proto.RegisterVideoServiceHandler(service.Server(), new(controllers.VideoRpcController))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
