package models

import "gopkg.in/mgo.v2/bson"

// Resp ...
type (
	Resp struct {
		Data interface{}
		Err  string
	}

	// Req ...
	Req struct {
		Name     string
		Email    string
		Password string
	}

	// User ...
	User struct {
		ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name     string        `json:"name" bson:"name"`
		Email    string        `json:"email" bson:"email"`
		Password string        `json:"password" bson:"password"`
	}

	// Login ...
	Login struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password" bson:"password"`
	}
	// LoginRes ...
	LoginRes struct {
		Token string
	}
)
