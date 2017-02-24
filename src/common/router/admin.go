package router

import (
	"github.com/lunny/tango"
	A "handler/admin"
)

func AdminRouter(g *tango.Group)  {
	g.Use(A.AdminMiddlerWare())
	g.Group("/AdminMain",func(g *tango.Group) {
		g.Route([]string{"GET:Get","POST:Post"},"/index",new(A.AdminMain))
	})
	g.Group("/AdminLogin",func(g *tango.Group) {
		g.Route([]string{"GET:Get","POST:Post"},"/login",new(A.AdminLogin))
	})
	g.Group("/AdminLogin",func(g *tango.Group) {
		g.Route([]string{"GET:Get"},"/index",new(A.AdminUser))
	})


}