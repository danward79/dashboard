package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/danward79/dashboard/controllers"
)

var env *controllers.Env

func init() {
	log.Println("Dashboard Web App")

	dbPath := flag.String("d", "", "path to database")
	flag.Parse()

	if *dbPath == "" {
		flag.PrintDefaults()
		log.Println("database source required")
		os.Exit(1)
	}

	env = controllers.Init(*dbPath)
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
