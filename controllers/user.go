package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t := template.Must(template.ParseFiles("templates/users/register.html"))
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Fprintln(w, r.Form["nickname"])
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprintln(w, r.URL.Scheme)
		fmt.Fprintln(w, r.Form["url_long"])
		for k, v := range r.Form {
			fmt.Fprintln(w, k)
			fmt.Fprintln(w, strings.Join(v, ""))
		}
		fmt.Printf("%#v", w)
		fmt.Fprintln(w, "Hello astaxie!")
	}
}
