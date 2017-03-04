package models

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type User struct {
	NickName  string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	user := User{
		NickName:  "Tsukasa Kakubari",
		Age:       33,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	key := datastore.NewIncompleteKey(c, "users", nil)
	_, err := datastore.Put(c, key, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
