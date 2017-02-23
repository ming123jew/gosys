package handler

import (
	. "common"
	"github.com/tango-contrib/rbac"
)
type AdminMain struct {
	BaseHandler
	rbac.Perm `"Write" "Read" "Reset" "Delete"`
	rbac.Role
}

func (x *AdminMain)Get()  {
	var params = make(map[string]interface{})
	x.HTML("administrator/index.html",params)
}

func (x *AdminMain)Post() ()  {

}