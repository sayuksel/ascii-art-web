package PKG

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
	"log"
)

// handling 500 errors
func Errors500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	fileNameError := filepath.Join("..", "template", "500.html")
	t, err := template.ParseFiles(fileNameError)
	if err != nil {
		fmt.Fprintf(w, "<h1>500</h1><br><h1>ERROR in reading the 500 HTML template</h1>")
		return
	}

	err = t.ExecuteTemplate(w, "500.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handling 400 errors
func Errors400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	fileNameError := filepath.Join("..", "template", "400.html")
	t, err := template.ParseFiles(fileNameError)
	if err != nil {
		fmt.Fprintf(w, "<h1>400</h1><br><h1>ERROR in reading the 400 HTML template</h1>")
		return
	}

	err = t.ExecuteTemplate(w, "400.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
