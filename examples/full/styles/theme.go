package styles

import (
	"fmt"
	"github.com/AccentDesign/gcss/props"
	"io"
	"slices"
)

type Theme struct {
	MediaQuery        string
	Background        props.Color
	Foreground        props.Color
	Primary           props.Color
	PrimaryForeground props.Color
}

func (t *Theme) CSS(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s{", t.MediaQuery); err != nil {
		return err
	}
	for _, style := range slices.Concat(
		t.Base(),
		t.Buttons(),
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
