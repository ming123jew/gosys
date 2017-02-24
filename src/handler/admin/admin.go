package handler


import (

	"fmt"

)
//登陆结构
type AdminLogin struct {
	AdminBaseHandler
	Username string
	Password string
}



func (x *AdminLogin) Get() {

	fmt.Println( x.Ctx.Req() )
	x.Session.Set("test", "1")
	a := x.Session.Get("test").(string)
	fmt.Println(a)
	var params = make(map[string]interface{})
	x.HTML("administrator/login.html",params)
}

func (x *AdminLogin) Post()  {
	

}

func (x *AdminLogin)Del() {

	x.ResponseWriter.Write([]byte("ok"))
}

 





