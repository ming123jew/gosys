package handler

import "log"

type AdminMain struct {
	AdminBaseHandler
}

func (x *AdminMain)Get()  {

	log.Println(x.Session.Get(SESSION_NAME_ADMIN))
	var params = make(map[string]interface{})

	//x.Template("administrator/index.html").Execute(x.Ctx.ResponseWriter,params)
	/*
	_,err :=x.Renderer.Template("administrator/index.html").ParseFiles("templates/administrator/menu.html")
	if err != nil {
		log.Println(err)
	}
	err = x.Renderer.Template("administrator/index.html").Execute(x.Ctx.ResponseWriter,params)
	if err != nil {
		log.Println(err)
	}*/
	x.HTML("administrator/index.html",params)
}

func (x *AdminMain)Post() ()  {

}

