package styles

import (
	"bytes"
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
	var buf bytes.Buffer
	for _, style := range slices.Concat(
		m.Layout(),
	) {
		if err := style.CSS(&buf); err != nil {
			return err
		}
	}
	if buf.Len() > 0 {
		_, err := fmt.Fprintf(w, "%s{%s}", m.Query, buf.String())
		return err
	}
	return nil
}
