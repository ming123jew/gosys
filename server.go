package main

import (
	"github.com/lunny/tango"
	"github.com/tango-contrib/session"
	"github.com/tango-contrib/binding"
	TgRbac "github.com/tango-contrib/rbac"
	"os"
	. "common"
	//H "handler/home"
	"time"
	"common/rbac"
	"common/router"
)

func main()  {
	//初始化tango
	Tg := tango.Classic()
	//模板
	Tg.Use(Renders)
	//SESSION
	sessions := session.New(session.Options{
		MaxAge:time.Minute * 20,
		//Store: redistore.New(Options{
		//	Host:    "127.0.0.1",
		//	DbIndex: 0,
		//	MaxAge:  30 * time.Minute,
		//},
		})
	Tg.Use(sessions)
	//binding 自动提取请求参数到结构体的映射和要求检查
	Tg.Use(binding.Bind())

	//轻量级rbac
	Rbac := rbac.NewRbac()
	Tg.Use(TgRbac.RBAC(Rbac, sessions))

	Tg.Use(MiddlerWare())
	//静态文件服务器
	Tg.Use(tango.Static(tango.StaticOptions{RootPath:"static"}))

	//路由
	Tg.Group("/admin", func(g *tango.Group) {
		router.AdminRouter(g)

	})

	Tg.Group("/home", func(g *tango.Group) {
		router.HomeRouter(g)

	})
	//设置访问端口
	os.Setenv("PORT",Cfg.MustValue("common","http_port","8000"))
	os.Setenv("HOST",Cfg.MustValue("common","http_host",""))
	//Tg.RunTLS("F:\\go\\ca\\ca.crt","F:\\go\\ca\\ca.key")
	Tg.Run()
}