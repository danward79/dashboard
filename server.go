package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/danward79/dashboard/controllers"
)

//Env ...
type Env struct {
	*controllers.Env
}

var env *Env

func init() {
	log.Println("Dashboard Web App")

	env = controllers.Init("data.db")

	fmt.Println(env.DB)
}

func main() {

	//Routing
	r := mux.NewRouter()
	r.HandleFunc("/chart", env.Chart)
	r.HandleFunc("/", controllers.Default)
	r.NotFoundHandler = http.HandlerFunc(controllers.Default)

	//Middleware and static file handling
	n := negroni.Classic()
	n.UseHandler(r)

	//Run the server
	n.Run(":3000")
}
