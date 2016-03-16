package controllers

import (
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

//pageExecute page handling
func pageExecute(w http.ResponseWriter, r *http.Request, file string) {

	tp := path.Join("views", file+".html")

	fi, err := os.Stat(tp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	if fi.IsDir() {
		http.NotFound(w, r)
		return
	}

	templateView(w, tp)

}

//templateView
func templateView(w http.ResponseWriter, tmpl string) {

	lp := path.Join("views", "layout.html")

	t, err := template.ParseFiles(lp, tmpl)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
		return
	}

	t.ExecuteTemplate(w, "layout", nil)
}
