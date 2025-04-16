package render

import (
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles(
		"./templates/"+templateName+".html",
		"./templates/base.layout.tmpl.html",
	)
	if err := parsedTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
