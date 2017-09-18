package controllers

import (
	"beego-login/helpers"
	"beego-login/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// UserController ...
type UserController struct {
	beego.Controller
}

// Post ...
func (c *UserController) Post() {
	var reqData models.Req

	reqData.Email = c.GetString("email")
	reqData.Name = c.GetString("name")
	reqData.Password = c.GetString("password")
	resp := models.InsertUser(reqData)

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Error(err)
	}

}

// Get ...
func (c *UserController) Get() {
	resp := models.GetAllUser()

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Error(err)
	}
}

// FilterUser ...
var FilterUser = func(c *context.Context) {

	ezT := helpers.EzToken{}
	authToken := strings.TrimSpace(c.Request.Header.Get("Authorization"))
	valid, err := ezT.ValidateToken(authToken)
	beego.Debug(valid)
	if !valid {
		c.ResponseWriter.WriteHeader(401)
		resp := models.Resp{
			Data: nil,
			Err:  err.Error(),
		}
		err := c.Output.JSON(resp, false, false)
		if err != nil {
			beego.Error(err)
		}
		return
	}

}
