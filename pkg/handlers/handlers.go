package handlers

import (
	"app/pkg/render"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
