package enquete

import (
	"enquete/routes"
	"net/http"

	"github.com/codegangsta/negroni"
)

func init() {
	r := routes.Init()

	n := negroni.Classic()
	n.UseHandler(r)

	http.Handle("/", n)
}
