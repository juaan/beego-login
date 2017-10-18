package models

import "gopkg.in/mgo.v2/bson"

type (
	// Resp ...
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
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Name      string        `json:"name" bson:"name"`
		Email     string        `json:"email" bson:"email"`
		Password  string        `json:"password" bson:"password"`
		FavTracks Songs         `json:"fav_tracks" bson:"fav_tracks"`
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

	// Song ...
	Song struct {
		ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
		SongID    string        `json:"song_id" bson:"song_id"`
		SongTitle string        `json:"song_title" bson:"song_title"`
		SongImage string        `json:"song_image" bson:"song_image"`
		SongURI   string        `json:"song_uri" bson:"song_uri"`
		Artist    string        `json:"artist" bson:"artist"`
		ArtistAva string        `json:"artist_ava" bson:"artist_ava"`
	}
	// SongBody ...
	SongBody struct {
		Song Song
	}
	// Songs ...
	Songs []Song
)
