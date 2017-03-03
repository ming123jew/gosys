package handler
import (


)
import (
	"model"
	"strconv"
)

//用户操作结构
type AdminUser struct {
	AdminBaseHandler
	Username string
	Password string
}

//列表
func (x *AdminUser)List(){

	p:=x.Ctx.Req().FormValue("page")
	page,err := strconv.Atoi(p)
	if err != nil{}

	list,_ := new(model.User).SelectByPage(1,page,true)

	x.Cache.Put("test", "Hello Tango!", 20)

	var params = make(map[string]interface{})
	params["list"] = list
	params["cache"] = x.Cache.Get("test")

	x.HTML2([]string{"templates/administrator/user/list.html","templates/administrator/menu.html","templates/administrator/header.html"},params)

}
//添加
func (x *AdminUser)Add()  {

}
//删除
func (x *AdminUser)Del()  {

}
//更新
func (x *AdminUser)Update()  {

}