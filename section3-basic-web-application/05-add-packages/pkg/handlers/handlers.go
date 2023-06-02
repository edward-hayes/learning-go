package handlers

import (
	"net/http"

	"github.com/edward-hayes/go-course/pkg/render"
)

// home is the home page template
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// about is the about page template
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
