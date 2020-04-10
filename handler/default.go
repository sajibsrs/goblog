// This file handles default routes of the application
// and serves templates accordingly

package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/default/layout.html"))
	tmpl.ExecuteTemplate(w, "layout", "content")
}