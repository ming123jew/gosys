package handler

import (


)
import (
	. "common"
	"fmt"
)

const (


)


type AdminHandler struct {
	BaseHandler

}



type AdminLogin struct {
	AdminHandler
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





