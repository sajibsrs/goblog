// This file handles default routes of the application
// and serves templates accordingly

package handler

import (
	"html/template"
	"log"
	"net/http"
)

// Index handles default request
func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"template/default/layout.html",
		"template/default/content.html",
		"template/default/navigation.html"}
	tmpl := template.Must(template.ParseFiles(files...))
	err := tmpl.ExecuteTemplate(w, "layout", r)
	if err != nil {
		log.Println("Unable to excute template", err)
	}
}
