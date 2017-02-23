package common

import (
	"common/orm"
	"common/conf"
	"github.com/go-xorm/xorm"
	cfg "github.com/Unknwon/goconfig"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/session"
	"github.com/tango-contrib/binding"
	"github.com/lunny/tango"
	"github.com/tango-contrib/rbac"
)


var	Orm  *xorm.Engine
var 	Cfg  *cfg.ConfigFile
var 	Renders *renders.Renders



func init()  {

	cfg := conf.NewCfg()
	var server=cfg.MustValue("db","server","127.0.0.1")
	var username=cfg.MustValue("db","username","username")
	var password=cfg.MustValue("db","password","")
	var db_name=cfg.MustValue("db","db_name","db_name")
	//var show_sql=Cfg.MustValue("db","show_sql","show_sql")
	var charset=cfg.MustValue("db","charset","utf8")
	var port=cfg.MustValue("db","port","3306")
	orm := orm.NewOrm(server,username,password,db_name,charset,port)
	Cfg = cfg
	Orm = orm

	Renders = renders.New(renders.Options{
		Reload: true,
		Directory: "./templates",
	})


}

func MiddlerWare() tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {


		}
		ctx.Next()
	}
}

//操作基本类型
type BaseHandler struct {
	renders.Renderer
	session.Session
	binding.Binder
	tango.Ctx
	rbac.Manager
}

func (x *BaseHandler)HTML(name string,T ...map[string]interface{})  {

	sys_params := map[string]interface{}{

		"map_appkey":Cfg.MustValue("map","map_appkey",""),
		"map_url":Cfg.MustValue("map","map_url",""),
		"static_url":Cfg.MustValue("common","static_url",""),

	}
	T2 := make(map[string]interface{})
	for _,v := range T{
		T2 = v
	}

	x.Render(name,renders.T{
		"C": sys_params, //common
		"P": T2,		 //params
	})
}