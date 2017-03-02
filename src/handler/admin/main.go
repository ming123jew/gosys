package handler

import "log"

type AdminMain struct {
	AdminBaseHandler
}

func (x *AdminMain)Get()  {

	log.Println()
	var params = make(map[string]interface{})

	x.HTML2("administrator/index.html",[]string{"templates/administrator/menu.html","templates/administrator/a.html"},params)
	//x.HTML("administrator/index.html",params)
}

func (x *AdminMain)Post() ()  {

}

