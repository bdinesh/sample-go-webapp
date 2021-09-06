package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders the given HTML template
func RenderTemplate(w http.ResponseWriter, templateName string) {
	filename := fmt.Sprintf("./templates/%s.html", templateName)
	parsedTemplate, _ := template.ParseFiles(filename)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
