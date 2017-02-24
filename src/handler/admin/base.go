package handler
//admin base
//主要定义结构类型，并分配每个接口类型的角色权限
import (
	. "common"

	//"strings"
	"github.com/mikespook/gorbac"
	"encoding/json"
	"github.com/lunny/tango"

	"fmt"
)

const (
	USER_AUTH_TYPE 		=   2			// 默认认证类型 0:不认证 1 session登录认证 2 实时认证
	NOT_AUTH_MODULE         = "public,user,index"	// 默认无需认证模块
	NOT_AUTH_ACTION         = "login"		// 默认无需认证操作

	ADMIN_FLAG = "admin"			//后台入口标示

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

func Check()  {
	//认证类型
	if  USER_AUTH_TYPE > 0 {

		b,s := CheckPermission("ROLE_ADMIN")
		fmt.Println(b,s)
	}
}

func AdminMiddlerWare() tango.HandlerFunc  {
	return func(ctx *tango.Context) {
		if action := ctx.Action(); action != nil {
			Check()
		}
		ctx.Next()
	}
}
