package controllers

import (
	"fmt"
	"log"
	"net/http"
)

//Chart ...
func (e *Env) Chart(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path
	if f == "/" {
		f = "chart"
	}

	pageExecute(w, r, f)

	c, err := e.DB.Chart(7, 4)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Println(string(c))

}
