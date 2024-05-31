package main

import (
	"fmt"
	"github.com/AccentDesign/gcss/variables"
	"net/http"
)

var (
	darkTheme = &Theme{
		Background:        variables.Zinc800,
		Foreground:        variables.White,
		Primary:           variables.White,
		PrimaryForeground: variables.Zinc800,
	}
	lightTheme = &Theme{
		Background:        variables.White,
		Foreground:        variables.Zinc800,
		Primary:           variables.Zinc800,
		PrimaryForeground: variables.White,
	}
	html = `
		<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="/dark.css" media="(prefers-color-scheme: dark)">
			<link rel="stylesheet" href="/light.css" media="(prefers-color-scheme: light)">
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

	http.HandleFunc("/dark.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		if err := darkTheme.WriteCSS(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/light.css", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
		if err := lightTheme.WriteCSS(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
