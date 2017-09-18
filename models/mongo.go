package models

import (
	"github.com/astaxie/beego"
	mgo "gopkg.in/mgo.v2"
)

// ConnectMgo ...
func ConnectMgo() (*mgo.Session, error) {
	dbURL := "mongodb://localhost:27017"
	sess, err := mgo.Dial(dbURL)
	CheckErr(err, "Error Connect Mongo")

	sess.SetMode(mgo.Monotonic, true)

	return sess, err
}

// GetCollection ...
func GetCollection(name string) (*mgo.Session, *mgo.Collection) {
	conn, err := ConnectMgo()
	CheckErr(err, "Error Connect Mongo from get collection")
	coll := conn.DB("matakaki").C(name)

	return conn, coll
}

// CheckErr ...
func CheckErr(err error, msg string) {
	if err != nil {
		beego.Warning(msg)
		beego.Warning(err)
	}
}
