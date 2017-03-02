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
	mgov2 "gopkg.in/mgo.v2"
	"common/mgo"
	myRbac "common/rbac"
	"regexp"
	"strings"
	"github.com/mikespook/gorbac"
	"time"
	"log"

)

const(
	ADMIN_FLAG = "admin"			//后台入口标示
)

var	Orm  *xorm.Engine
var 	Cfg  *cfg.ConfigFile
var 	Renders *renders.Renders
var 	Mgo *mgov2.Session
var	Rbac *gorbac.RBAC
var 	Sessions *session.Sessions
var	Session  *session.Session

var	ActionName string
var	MethodName string
var	UriArray   []string




func init()  {

	sessions := session.New(session.Options{
		MaxAge:time.Minute * 20,
		//Store: redistore.New(Options{
		//	Host:    "127.0.0.1",
		//	DbIndex: 0,
		//	MaxAge:  30 * time.Minute,
		//},
	})
	Sessions = sessions

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

	Rbac = myRbac.NewRbac()

	mgo := mgo.NewMgo("127.0.0.1:27017,127.0.0.1:27018,127.0.0.1:27019")
	Mgo = mgo

	Renders = renders.New(renders.Options{
		Reload: true,
		Directory: "./templates",
	})



}

func MiddlerWare(s *session.Sessions) tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {
			action := ctx.Route().Method().String()
			reg := regexp.MustCompile(`([^<func(*handler.].*[^) Value>])`)
			ActionName = strings.ToLower(reg.FindAllString(action,-1)[0])			//操作名称
			UriArray = strings.Split(strings.ToLower(strings.Split(ctx.Req().RequestURI, "?")[0]), "/")  //地址分割

			if len(UriArray)>3{
				MethodName = UriArray[3]
			}else{
				MethodName = "index"
			}

			//后台地址需要验证路由操作
			if UriArray[2]==ADMIN_FLAG{
				if ActionName != UriArray[3]{
					panic("check route and action!")
				}
			}

			//fmt.Println("common.go:",ActionName,MethodName)
			ss := s.Session(ctx.Req(), ctx.ResponseWriter)
			Session = ss
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
func (x *BaseHandler)HTML2(name string,names []string,T ...map[string]interface{})  {
	sys_params := map[string]interface{}{

		"map_appkey":Cfg.MustValue("map","map_appkey",""),
		"map_url":Cfg.MustValue("map","map_url",""),
		"static_url":Cfg.MustValue("common","static_url",""),

	}
	T2 := make(map[string]interface{})
	for _,v := range T{
		T2 = v
	}

	_,err :=x.Renderer.Template(name).ParseFiles(names...)
	if err != nil {
		log.Println("common.go:153",err)
	}

	err = x.Renderer.Template(name).Execute(x.Ctx.ResponseWriter,renders.T{
		"C": sys_params, //common
		"P": T2,		 //params
	})
	if err != nil {
		log.Println("common.go:161",err)
	}
}
