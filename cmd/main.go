package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"web/pkg"
)

func form(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Join("..", "template", "form.html")
	t, err := template.ParseFiles(fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		PKG.Errors500(w)
		return
	}

	error := t.ExecuteTemplate(w, "form.html", nil)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		PKG.Errors500(w)
	}
}

// running the server handling any errors and sending them to correct status code
func serverHandler(w http.ResponseWriter, r *http.Request) {

	art, err := PKG.Text(r.FormValue("text"), r.FormValue("art"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("err: %v\n", err)
		PKG.Errors500(w)
		return
	}

	switch r.URL.Path {
	case "/":
		form(w, r)
	case "/ascii-art":
		if art == "500" {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("err: %v\n", err)
			PKG.Errors500(w)
			return
		}
		fileresult := filepath.Join("..", "template", "result.html")
		f, err := template.ParseFiles(fileresult)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			PKG.Errors500(w)
			return
		}

		if art == "400" {
			w.WriteHeader(http.StatusBadRequest)
			PKG.Errors400(w)
			return
		}
		err = f.ExecuteTemplate(w, "result.html", nil)
		fmt.Fprintf(w, "<h1>This is ASCII web </h1>")
		fmt.Fprintf(w, art)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("err: %v\n", err)
			PKG.Errors500(w)
		}

	default:
		w.WriteHeader(http.StatusNotFound)
		fileNameError := filepath.Join("..", "template", "404.html")
		t, err := template.ParseFiles(fileNameError)
		if err != nil {
			fmt.Fprintf(w, "<h1>404</h1><br><h1>ERROR in reading the 404 HTML template</h1>")
			return
		}

		err = t.ExecuteTemplate(w, "404.html", nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Printf("err: %v\n", err)
			PKG.Errors500(w)
		}

	}
}

// running the server
func main() {
	http.HandleFunc("/", serverHandler)
	fmt.Println("Server is running at localhost:8080")
	// Create a file server for serving CSS files
	styles := http.FileServer(http.Dir("../stylesheets"))
	http.Handle("/stylesheets/", http.StripPrefix("/stylesheets/", styles))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
}
