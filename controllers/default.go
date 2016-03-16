package controllers

import "net/http"

//Default handler
func Default(w http.ResponseWriter, r *http.Request) {
	f := r.URL.Path
	if f == "/" {
		f = "index"
	}

	pageExecute(w, r, f)
}
