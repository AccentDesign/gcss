package styles

import (
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"github.com/AccentDesign/gcss/variables"
	"io"
	"slices"
)

type (
	Media struct {
		Query       string
		Padding     props.Unit
		VerticalGap props.Unit
	}
)

var (
	mobileMedia = &Media{
		Query:       "@media (max-width: 768px)",
		Padding:     variables.Size8,
		VerticalGap: variables.Size6,
	}
	desktopMedia = &Media{
		Query:       "@media (min-width: 769px)",
		Padding:     variables.Size16,
		VerticalGap: variables.Size8,
	}
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
