// This file handles user routes of the application
// and serves templates accordingly

package handler

import (
	"html/template"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"template/default/layout.html",
		"template/user/detail.html",
		"template/default/navigation.html"}

	tmpl := template.Must(template.ParseFiles(files...))
	_ = tmpl.ExecuteTemplate(w, "layout", r)
}
