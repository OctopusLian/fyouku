package controllers

import (
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type TestController struct {
	beego.Controller
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// @router /test/index [get]
func (this *TestController) Get() {
	this.TplName = "danmuDemo.html"
}

func (this *TestController) WsFunc() {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)
	if conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil); err != nil {
		goto ERR
	}

	go func() {
		for {
			if err = conn.WriteMessage(websocket.TextMessage, []byte("Hello")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if _, data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
