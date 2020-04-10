// This file handles default routes of the application
// and serves templates accordingly

package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpls := []string{
		"template/default/layout.html",
		"template/default/content.html",
		"template/default/navigation.html"}

	tmpl := template.Must(template.ParseFiles(tmpls...))
	tmpl.ExecuteTemplate(w, "layout", r)
}