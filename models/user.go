package models

import (
	"beego-login/helpers"
	"time"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

// InsertUser ...
func InsertUser(reqdata Req) (dataResp Resp) {

	// query to find if the email is already registered
	query := bson.M{"email": reqdata.Email}
	conn, coll := GetCollection("user")

	defer conn.Close()
	countCheck, err := coll.Find(query).Count()
	if err != nil {
		beego.Error(err)
		dataResp.Err = err.Error()
	}

	if countCheck > 0 {
		dataResp.Err = "Email already registered"
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(reqdata.Password), 5)
	if err != nil {
		beego.Error(err)
	}
	reqdata.Password = string(hash)

	err = coll.Insert(&reqdata)
	if err != nil {
		beego.Error(err)
		dataResp.Err = err.Error()
		return
	}

	dataResp.Data = reqdata
	return

}

// GetAllUser ...
func GetAllUser() (dataResp Resp) {
	conn, coll := GetCollection("user")
	defer conn.Close()
	var users []User
	err := coll.Find(nil).All(&users)
	if err != nil {
		dataResp.Err = err.Error()
		return
	}

	dataResp.Data = users
	return

}

// AddSong ...
func AddSong(id string, songID string) (dataResp Resp) {
	conn, coll := GetCollection("user")
	defer conn.Close()

	arr := []string{songID}

	n, err := coll.Find(bson.M{"fav_tracks": bson.M{"$in": arr}}).Count()
	if err != nil {
		dataResp.Err = err.Error()
		return
	}

	if n == 0 {
		where := bson.M{"_id": bson.ObjectIdHex(id)}
		cond := bson.M{"$push": bson.M{"fav_tracks": songID}}
		err = coll.Update(where, cond)
		if err != nil {
			dataResp.Err = err.Error()
			return
		}

		dataResp.Data = "Song has been added"
		return
	}

	dataResp.Err = "song already in your fav list"
	return

}

// GetJWT ...
func GetJWT(loginData Login) (dataResp Resp) {
	var user User
	var loginRes LoginRes

	conn, coll := GetCollection("user")
	defer conn.Close()
	err := coll.Find(bson.M{"email": loginData.Email}).One(&user)
	if err != nil {
		dataResp.Err = err.Error()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		dataResp.Err = "Wrong Password"
		return
	}
	ezT := helpers.EzToken{
		Username: user.Name,
		ID:       user.ID.String(),
		Expires:  time.Now().Unix() + 3600,
	}
	token, err := ezT.GetToken()
	if err != nil {
		dataResp.Err = "Failed Generating token"
		return
	}

	loginRes.Token = token
	dataResp.Data = loginRes

	return
}

// GetUserSongs ...
// func GetUserSongs(id string) (dataResp Resp) {
// 	conn, coll := GetCollection("user")
// 	defer conn.Close()
// 	var songs Songs
//
// 	err := coll.Find()
// 	return
// }
