package main

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"html/template"
	"net/http"
	"strings"
)

type Stylesheet []gcss.Style

func (ss Stylesheet) CSS() (template.CSS, error) {
	var sb strings.Builder
	for _, style := range ss {
		if err := style.CSS(&sb); err != nil {
			return "", err
		}
	}
	return template.CSS(sb.String()), nil
}

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
)

var homeTemplate *template.Template

func main() {
	var err error
	homeTemplate, err = template.New("home.gohtml").Funcs(template.FuncMap{
		"css": func() template.CSS {
			// todo: handle error
			css, _ := styles.CSS()
			return css
		},
	}).ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := homeTemplate.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server started at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
