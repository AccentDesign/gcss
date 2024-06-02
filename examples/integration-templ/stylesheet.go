package main

import (
	"context"
	"github.com/AccentDesign/gcss"
	"io"
)

// Stylesheet is a collection of styles.
type Stylesheet []gcss.Style

// Render writes the CSS representation of the stylesheet to the writer.
// this is to ensure that it implements a templ.Component.
func (ss Stylesheet) Render(ctx context.Context, w io.Writer) error {
	//TODO: Get a CSP nonce from the context.
	if _, err := io.WriteString(w, "<style type=\"text/css\">\n"); err != nil {
		return err
	}
	for _, s := range ss {
		if err := s.CSS(w); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "\n</style>"); err != nil {
		return err
	}
	return nil
}
