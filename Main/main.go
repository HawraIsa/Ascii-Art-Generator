package main

import (
	"asciiartwebstylize"
	"os"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Result string
}

func main() {
	// handle function for the home page & ascii-art page

	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Listen & serve http://localhost:8080/")
	fmt.Println("Opening in Browser...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server ..", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// test if page file NOT FOUND
	if r.URL.Path != "/" {
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/404.html")
		return
	}
	// if form parameter method = GET
	if r.Method == "GET" { // display home page
		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
			http.ServeFile(w, r, "templates/500.html")
			return
		}
		w.WriteHeader(200)

		tmpl.Execute(w, nil)
	} else {
		fmt.Println("Method Not Allowed:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "<h1>405 Method Not Allowed</h1>")
		fmt.Fprintf(w, "<p>Sorry, the requested HTTP method is not allowed.</p>")
	}

}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// if form parameter method = POST
	if r.Method == "POST" {
		text := r.FormValue("text") // get values from form
		banner := r.FormValue("banner")

		if !asciiartwebstylize.Validate(text) || !asciiartwebstylize.Validatefont(banner) {
			w.WriteHeader(http.StatusBadRequest)
			w.WriteHeader(400)
			http.ServeFile(w, r, "templates/400.html")
			return
		}

		tmpl, err := template.ParseFiles("templates/ascii-art.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the server file for the banner name exists
		_, err = os.Stat("banners/" + banner + ".txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.ServeFile(w, r, "templates/500.html")
			return
		}

		// Replace with a new line only
		lines := strings.ReplaceAll(text, "\r\n", "\n")

		result := asciiartwebstylize.Matching1(lines, banner)

		// append
		data := PageData{Result: result}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.WriteHeader(404)
		http.ServeFile(w, r, "templates/404.html")
		return
	}
}
