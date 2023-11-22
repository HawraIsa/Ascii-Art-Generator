package main

import (
	"asciiartwebstylize"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

type PageData struct {
	Result string
}

/* ------------------------------------------------------------ MAIN FUNCTION --------------------------------------------- */
func main() {
	// handle stylesheet
	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Serving http://localhost:8080/")
	fmt.Println("Opening in Browser...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server ..", err)
	}
}

/* -------------------------------------------------------- INDEX HANDLER FUNCTION --------------------------------------------- */

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// if index.html not found
	if r.URL.Path != "/" {
		tmpl, err := template.ParseFiles("templates/404.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
			http.ServeFile(w, r, "templates/500.html")
			return
		}

		w.WriteHeader(404)
		tmpl.Execute(w, nil)
		return
	}
	// if form parameter method is GET, display home page
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles("templates/index.html")

		if err != nil {
			tmpl, _ := template.ParseFiles("templates/500.html")
			w.WriteHeader(500)
			tmpl.Execute(w, nil)
			return
		}
		// otherwise, write the status as OK(200) and display the index page
		w.WriteHeader(200)
		tmpl.Execute(w, nil)

	} else { // if method is not GET
		fmt.Println("Method Not Allowed:", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "<h1>405 Method Not Allowed</h1>")
		fmt.Fprintf(w, "<p>Sorry, the requested HTTP method is not allowed.</p>")
	}
}

/* -------------------------------------------------------- ASCII HANDLER FUNCTION --------------------------------------------- */

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii-art" {
		tmpl, err := template.ParseFiles("templates/404.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.WriteHeader(500)
			http.ServeFile(w, r, "templates/500.html")
			return
		}
		w.WriteHeader(404)
		tmpl.Execute(w, nil)
		return
	}

	// if form parameter method is POST, get required values from form
	if r.Method == "POST" {
		text := r.FormValue("text")
		banner := r.FormValue("banner")
		// validate font & text
		if !asciiartwebstylize.Validate(text) || !asciiartwebstylize.Validatefont(banner) {
			tmpl, err := template.ParseFiles("templates/400.html")

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.WriteHeader(500)
				http.ServeFile(w, r, "templates/500.html")
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			w.WriteHeader(400)
			tmpl.Execute(w, nil)
			return
		}

		tmpl, err := template.ParseFiles("templates/ascii-art")
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
		// generate ascii
		result := asciiartwebstylize.Matching1(lines, banner)

		// append ascii to the struct then display ascii art page
		data := PageData{Result: result}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		fs := http.FileServer(http.Dir("./assets/"))
		http.Handle("/assets/", http.StripPrefix("/assets/", fs))
		tmpl, err := template.ParseFiles("templates/404.html")
		if err != nil {
			http.ServeFile(w, r, "templates/500.html")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(404)
		w.WriteHeader(http.StatusNotFound)
		tmpl.Execute(w, nil)
	}
}

// w.WriteHeader(404)
// tmpl, _ := template.ParseFiles("templates/404.html")
// http.ServeFile(w, r,"./assets/style.css")
// tmpl.Execute(w, nil)
// return
