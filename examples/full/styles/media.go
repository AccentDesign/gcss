package styles

import (
	"fmt"
	"io"
	"slices"
)

type (
	Media struct {
		MediaType
		Query string
	}
	MediaType int
)

const (
	Mobile MediaType = iota
	Desktop
)

var (
	mobileMedia  = &Media{MediaType: Mobile, Query: "@media (max-width: 768px)"}
	desktopMedia = &Media{MediaType: Desktop, Query: "@media (min-width: 769px)"}
)

// CSS Writes the CSS for the media to the writer.
func (m *Media) CSS(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s{", m.Query); err != nil {
		return err
	}
	for _, style := range slices.Concat(
		m.Layout(),
	) {
		if err := style.CSS(w); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "}"); err != nil {
		return err
	}
	return nil
}
