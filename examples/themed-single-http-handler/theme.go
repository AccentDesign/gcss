package main

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"io"
)

type (
	Styles []gcss.Style
	Theme  struct {
		MediaQuery        string
		Background        props.Color
		Foreground        props.Color
		Primary           props.Color
		PrimaryForeground props.Color
	}
)

func (t *Theme) WriteCSS(w io.Writer) error {
	if t.MediaQuery != "" {
		if _, err := fmt.Fprintf(w, "%s{", t.MediaQuery); err != nil {
			return err
		}
		for _, style := range t.styles() {
			if err := style.CSS(w); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprint(w, "}"); err != nil {
			return err
		}
	}
	return nil
}

func (t *Theme) styles() Styles {
	return Styles{
		{
			Selector: "body",
			Props: gcss.Props{
				BackgroundColor: t.Background,
				Color:           t.Foreground,
			},
		},
		{
			Selector: ".button-primary",
			Props: gcss.Props{
				BackgroundColor: t.Primary,
				Border: props.Border{
					Width: props.UnitPx(1),
					Style: props.BorderStyleSolid,
					Color: t.Primary,
				},
				BorderRadius:  variables.Size1H,
				Color:         t.PrimaryForeground,
				FontSize:      variables.Size4,
				PaddingBottom: variables.Size3,
				PaddingLeft:   variables.Size5,
				PaddingRight:  variables.Size5,
				PaddingTop:    variables.Size3,
			},
		},
		{
			Selector: ".button-primary:hover",
			Props: gcss.Props{
				BackgroundColor: t.Primary.Alpha(230),
			},
		},
	}
}
