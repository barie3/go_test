package routes

import (
	"html/template"
	"net/http"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/root.html"))
	t.Execute(w, nil)
}
