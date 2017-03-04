package routes

import (
	"enquete/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users/register", controllers.UserRegisterHandler)

	return r
}
