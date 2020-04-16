// This file handles helper and utility functions of the application

package helper

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// ProcessTemplates processes, parses and executes template files
//  w = response writer
//  name = template name
//  data = data to pass to the template
//  tmp = template names to process in "directory.template" format
// Default template directory is named template and extension should be omitted
func ProcessTemplates(w http.ResponseWriter, name string, data interface{} , tmp ...string) {
	var procTmp []string
	for _, t := range tmp {
		t = "template/" + strings.ReplaceAll(t, ".", "/") + ".html"
		procTmp = append(procTmp, t)
	}
	temps := template.Must(template.ParseFiles(procTmp...))
	err := temps.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Println("Unable to execute templates", err)
	}
}
