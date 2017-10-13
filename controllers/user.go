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

// AddSong ...
func (c *UserController) AddSong() {
	songID := c.Input().Get("song_id")
	user := c.Ctx.Input.Param(":id")
	resp := models.AddSong(user, songID)

	beego.Debug(songID)
	beego.Debug(user)
	// beego.Debug(resp)

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

	beego.Debug("user valid")

}
