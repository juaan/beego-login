package models

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

// InsertSong ...
func InsertSong(song Song) {
	conn, coll := GetCollection("song")

	defer conn.Close()
	song.ID = bson.NewObjectId()
	err := coll.Insert(&song)
	if err != nil {
		beego.Error(err)
	}
}
