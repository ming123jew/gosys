package main

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/session"
	"github.com/tango-contrib/binding"
	"os"
	. "common"

	A "handler/admin"
	//H "handler/home"
	"time"
)

func main()  {
	//初始化tango
	Tg := tango.Classic()
	//模板
	Tg.Use(Renders)
	//SESSION
	Tg.Use(session.New(session.Options{
		MaxAge:time.Minute * 20,
		//Store: redistore.New(Options{
		//	Host:    "127.0.0.1",
		//	DbIndex: 0,
		//	MaxAge:  30 * time.Minute,
		//},
	}))
	//binding 自动提取请求参数到结构体的映射和要求检查
	Tg.Use(binding.Bind())

	//静态文件服务器
	Tg.Use(tango.Static(tango.StaticOptions{RootPath:"static"}))

	//路由
	Tg.Group("/admin", func(g *tango.Group) {

		g.Route([]string{"GET:Get","POST:Post"},"/index",new(A.AdminMain))
		g.Route([]string{"GET:Get","POST:Post"},"/login",new(A.AdminLogin))

	})
	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host",""))
	//Tg.RunTLS("F:\\go\\ca\\ca.crt","F:\\go\\ca\\ca.key")
	Tg.Run()
}