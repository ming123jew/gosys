package handler
import (


)
//用户操作结构
type AdminUser struct {
	AdminBaseHandler
	Username string
	Password string
}

//列表
func (x *AdminUser)List(){

	var params = make(map[string]interface{})
	x.HTML2("administrator/user/list.html",[]string{"templates/administrator/menu.html"},params)

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