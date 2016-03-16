package controllers

import 	"net/http"


//Chart ...
func (e *Env) Chart(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path
	if f == "/" {
		f = "chart"
	}

	pageExecute(w, r, f)
}
