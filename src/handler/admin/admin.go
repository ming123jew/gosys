package handler


import (

	"log"
	"model"


)
//登陆结构
type AdminLogin struct {
	AdminBaseHandler
	Username string
	Password string
}



func (x *AdminLogin) Get() {


	//x.Session.Set("test", "1")
	//a := x.Session.Get("test").(string)

	var params = make(map[string]interface{})

	x.HTML("administrator/login.html",params)
}

func (x *AdminLogin) Post()  {

	var form AdminLogin
	x.Bind(&form)
	user := model.ChatUser{Username:form.Username,Password:form.Password}

	has, err := user.Exist()
	if err != nil{
		x.Ctx.Write([]byte("账号或密码错误."))
	}

	//验证账号真实性
	if user.Id==0 {
		x.Ctx.Write([]byte("账号不存在."))
		has = false
	}

	//
	if has == true && user.Id>0{

		log.Println(user)

		x.Session.Set(SESSION_NAME_ADMIN,user)

		x.Ctx.Redirect("/admin/AdminMain/index")
	}else{
		res := map[string]interface{}{
			"info"		:	"not is admin user",
			"status"	:	200,
		}
		x.Ctx.ServeJson(res)
	}

}

func (x *AdminLogin)Del() {

	x.ResponseWriter.Write([]byte("ok"))
}

 





