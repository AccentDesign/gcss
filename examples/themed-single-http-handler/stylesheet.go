package main

import (
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/props"
	"io"
)

type (
	Stylesheet struct {
		Dark  *Theme
		Light *Theme
	}
)

func (s *Stylesheet) WriteCSS(w io.Writer) error {
	for _, style := range s.styles() {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	if err := s.Dark.WriteCSS(w); err != nil {
		return err
	}
	if err := s.Light.WriteCSS(w); err != nil {
		return err
	}
	return nil
}

func (s *Stylesheet) styles() Styles {
	return Styles{
		{
			Selector: "html",
			Props: gcss.Props{
				FontFamily: props.FontFamilySans,
			},
		},
	}
}
