package handler
//admin base
//主要定义结构类型，并分配每个接口类型的角色权限
import (
	. "common"
	"github.com/mikespook/gorbac"
	"encoding/json"
	"github.com/lunny/tango"

	//"log"
	"model"
	"strings"
)

const (
	SESSION_NAME_ADMIN	= "gosys_session_name_admin"
	USER_AUTH_TYPE 		=   2			// 默认认证类型 0:不认证 1 session登录认证 2 实时认证
	NOT_AUTH_MODULE         = "public,user,index"	// 默认无需认证模块
	NOT_AUTH_ACTION         = "login"		// 默认无需认证操作



)

type AdminBaseHandler struct {
	BaseHandler
}

func CheckPermission(role string)(bool,string)  {
	permissionId := ActionName+":"+MethodName
	p :=  gorbac.NewStdPermission(permissionId)
	res := Rbac.IsGranted(role,p, nil)
	if !res{
		a := map[string]interface{}{
			"result":"you not have permission!",
			"status":200,
		}
		str, _ := json.Marshal(a)
		return false,string(str)

	}
	return true,"you have permission!"
}

func Check()  (bool,string){
	//action not need to verify
	for _, nap := range strings.Split(NOT_AUTH_ACTION, ",") {
		//log.Println(nap,":",params[2])
		if MethodName == nap {
			return true,"this method is not need verify."
		}
	}
	//认证类型
	if  USER_AUTH_TYPE > 0 {
		admin_user_session := Session.Get(SESSION_NAME_ADMIN)

		if admin_user_session != nil {
			admin_user := admin_user_session.(model.User)
			b,s := CheckPermission(admin_user.Roleid)
			//log.Println(b,s,admin_user.Roleid)
			return b,s
		}else{

			return false,"you have to login at first."
		}

	}else{
		return false,"check USER_AUTH_TYPE"
	}
}

func AdminMiddlerWare() tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {
			b,_ := Check()
			if b==true{
				ctx.Next()
			}else{
				//a := map[string]interface{}{
				//	"info":  info,
				//	"status":200,
				//}
				//ctx.ServeJson(a)
				ctx.Redirect("/admin/AdminLogin/login")
			}
		}


	}
}
