package rbac

import (

	"github.com/mikespook/gorbac"

)


func NewRbac()  *gorbac.RBAC {

	goRBAC := gorbac.New()

	//用户管理权限
	perm_adminuser_index := gorbac.NewStdPermission("adminuser:index")
	perm_adminuser_list := gorbac.NewStdPermission("adminuser:list")
	perm_adminuser_add := gorbac.NewStdPermission("adminuser:add")
	perm_adminuser_update := gorbac.NewStdPermission("adminUser:update")
	perm_adminuser_del := gorbac.NewStdPermission("adminUser:del")

	//后台主页权限
	perm_adminmain_index := gorbac.NewStdPermission("adminmain:index")
	perm_adminmain_list := gorbac.NewStdPermission("adminmain:list")
	perm_adminmain_add := gorbac.NewStdPermission("adminmain:add")
	perm_adminmain_update := gorbac.NewStdPermission("adminmain:update")
	perm_adminmain_del := gorbac.NewStdPermission("adminmain:del")

	role_admin := gorbac.NewStdRole("ROLE_ADMIN") //超级管理员
	role_admin.Assign(perm_adminuser_index)
	role_admin.Assign(perm_adminuser_list)
	role_admin.Assign(perm_adminuser_add)
	role_admin.Assign(perm_adminuser_update)
	role_admin.Assign(perm_adminuser_del)

	role_admin.Assign(perm_adminmain_index)
	role_admin.Assign(perm_adminmain_list)
	role_admin.Assign(perm_adminmain_add)
	role_admin.Assign(perm_adminmain_update)
	role_admin.Assign(perm_adminmain_del)

	goRBAC.Add(role_admin)
	return goRBAC
}

