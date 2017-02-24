package handler

import "log"

type AdminMain struct {
	AdminBaseHandler
}

func (x *AdminMain)Get()  {

	log.Println(x.Session.Get("aa"))
	var params = make(map[string]interface{})
	x.HTML("administrator/index.html",params)
}

func (x *AdminMain)Post() ()  {

}

