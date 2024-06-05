package main

import (
	"fmt"
	"github.com/AccentDesign/gcss/examples/full/styles"
	"net/http"
)

var (
	stylesheet = styles.NewStyleSheet()
	html       = `
		<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="/stylesheet.css">
		</head>
		<body>
			<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
			<button class="button button-primary">Click me</button>
		</body>
		</html>
	`
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprint(w, html); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/stylesheet.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		if err := stylesheet.CSS(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
