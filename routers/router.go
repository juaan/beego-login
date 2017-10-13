package routers

import (
	"beego-login/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/login",
			&controllers.LoginController{},
			"post:Post",
		),
		beego.NSRouter("/register",
			&controllers.UserController{},
			"post:Post",
		),
		beego.NSRouter("/user",
			&controllers.UserController{},
			"get:Get",
		),
		beego.NSRouter("/user/:id",
			&controllers.UserController{},
			"put:AddSong",
		),
	)

	beego.AddNamespace(ns)
	// adding middleware to auth jwts
	beego.InsertFilter("/v1/user/*", beego.BeforeRouter, controllers.FilterUser)

}
