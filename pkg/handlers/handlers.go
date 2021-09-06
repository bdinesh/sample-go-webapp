package handlers

import (
	"github.com/bdinesh/sample-go-webapp/pkg/render"
	"net/http"
)

// Home handles the home page request
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home")
}

// About handles the about page request
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about")
}
