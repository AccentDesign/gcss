package main

import (
	"fmt"
	"github.com/AccentDesign/gcss/variables"
	"net/http"
)

var (
	stylesheet = &Stylesheet{
		Dark: &Theme{
			MediaQuery:        "@media (prefers-color-scheme: dark)",
			Background:        variables.Zinc800,
			Foreground:        variables.White,
			Primary:           variables.White,
			PrimaryForeground: variables.Zinc800,
		},
		Light: &Theme{
			MediaQuery:        "@media (prefers-color-scheme: light)",
			Background:        variables.White,
			Foreground:        variables.Zinc800,
			Primary:           variables.Zinc800,
			PrimaryForeground: variables.White,
		},
	}
	html = `
		<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="/stylesheet.css">
		</head>
		<body>
			<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
			<button class="button-primary">Click me</button>
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
		if err := stylesheet.WriteCSS(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
