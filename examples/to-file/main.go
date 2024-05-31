package main

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"os"
)

type Stylesheet []gcss.Style

var styles = Stylesheet{
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

func main() {
	file, err := os.Create("stylesheet.css")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, style := range styles {
		if err := style.CSS(file); err != nil {
			panic(err)
		}
	}
}
