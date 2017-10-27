package controllers

import (
	"beego-login/helpers"
	"beego-login/models"
	"encoding/json"
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

// GetUserSongs ...
func (c *UserController) GetUserSongs() {
	resp := models.GetAllUser()

	err := c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Error(err)
	}
}

// AddSong ...
func (c *UserController) AddSong() {
	var songBd models.Song
	songBody := c.Ctx.Input.RequestBody
	// songBodyString := c.
	// beego.Warning(songBodytest)
	beego.Debug(string(songBody))
	err := json.Unmarshal(songBody, &songBd)
	if err != nil {
		beego.Error(err)
	} else {
		// Add song to DB
		models.InsertSong(songBd)
	}

	user := c.Ctx.Input.Param(":id")
	// Add song to user
	resp := models.AddSong(user, songBd.SongID)

	beego.Debug(songBd)
	beego.Debug(user)
	// beego.Debug(resp)

	err = c.Ctx.Output.JSON(resp, false, false)
	if err != nil {
		beego.Error(err)
	}
}

// FilterUser ...
var FilterUser = func(c *context.Context) {
	// beego.Debug(string(c.Input.RequestBody))
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
