package router

import (
	"github.com/lunny/tango"
	H "handler/home"
)

func HomeRouter(g *tango.Group)  {
	g.Route([]string{"GET:Get", "POST:Post"}, "/index", new(H.HomeHandler))
	g.Group("/User",func(g *tango.Group) {

		g.Route([]string{"GET:Get", "POST:Post"}, "/login", new(H.HomeHandler))
	})

}