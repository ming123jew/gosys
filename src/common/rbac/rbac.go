package rbac

import (

	"github.com/mikespook/gorbac"

)


func NewRbac()  *gorbac.RBAC {

	goRBAC := gorbac.New()
	rA := gorbac.NewStdRole("PUBLIC_ROLE")	//公共权限
	pA := gorbac.NewStdPermission("read")
	rA.Assign(pA)
	goRBAC.Add(rA)

	rB := gorbac.NewStdRole("writer")
	pB := gorbac.NewStdPermission("write")
	rB.Assign(pB)
	goRBAC.Add(rB)
	return goRBAC
}

