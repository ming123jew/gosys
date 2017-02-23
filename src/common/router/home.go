package router

import (
	"github.com/lunny/tango"
	A "handler/admin"
)

func HomeRouter(g *tango.Group)  {

	g.Route([]string{"GET:Get","POST:Post"},"/index",new(A.AdminMain))
	g.Route([]string{"GET:Get","POST:Post"},"/login",new(A.AdminLogin))

}