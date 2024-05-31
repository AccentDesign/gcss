package main

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"net/http"
)

type Stylesheet []gcss.Style

var (
	styles = Stylesheet{
		{
			Selector: "html",
			Props: gcss.Props{
				FontFamily: props.FontFamilySans,
			},
		},
		{
			Selector: ".button",
			Props: gcss.Props{
				BackgroundColor: variables.Zinc800,
				Border: props.Border{
					Width: props.UnitPx(1),
					Style: props.BorderStyleSolid,
					Color: variables.Zinc900.Alpha(128),
				},
				BorderRadius:  variables.Size1H,
				Color:         variables.White,
				FontSize:      variables.Size4,
				PaddingBottom: variables.Size3,
				PaddingLeft:   variables.Size5,
				PaddingRight:  variables.Size5,
				PaddingTop:    variables.Size3,
			},
		},
		{
			Selector: ".button:hover",
			Props: gcss.Props{
				BackgroundColor: variables.Zinc900,
			},
		},
	}

	html = `
		<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="/stylesheet.css">
		</head>
		<body>
			<button class="button">Click me</button>
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
		for _, style := range styles {
			if err := style.CSS(w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	})

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
