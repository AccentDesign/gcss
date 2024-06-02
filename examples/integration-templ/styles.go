package main

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/variables"
	"github.com/a-h/templ"
)

var (
	buttonStyles = Stylesheet{
		{
			Selector: ".button",
			Props: gcss.Props{
				BackgroundColor: variables.Zinc800,
				BorderRadius:    variables.Size1H,
				Color:           variables.White,
				FontSize:        variables.Size4,
				PaddingBottom:   variables.Size3,
				PaddingLeft:     variables.Size5,
				PaddingRight:    variables.Size5,
				PaddingTop:      variables.Size3,
			},
		},
		{
			Selector: ".button:hover",
			Props: gcss.Props{
				BackgroundColor: variables.Zinc900,
			},
		},
	}
	buttonStylesHandle = templ.NewOnceHandle()
)
