package controllers

import (
	"beego-login/models"

	"github.com/astaxie/beego"
)

// LoginController ...
type LoginController struct {
	beego.Controller
}

// Post ...
func (c *LoginController) Post() {
	var login models.Login
	login.Email = c.GetString("email")
	login.Password = c.GetString("password")

	resp := models.GetJWT(login)
	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Error(err)
	}
}
