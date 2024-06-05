package styles

import (
	"bytes"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/variables"
	"io"
	"slices"
	"sync"
)

type (
	Styles     []gcss.Style
	StyleSheet struct {
		Themes []*Theme
		css    bytes.Buffer
		mutex  sync.Mutex
	}
)

// CSS writes the stylesheet to the writer.
// It will generate the CSS for the stylesheet if it has not been generated yet.
func (ss *StyleSheet) CSS(w io.Writer) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()

	if ss.css.Len() > 0 {
		_, err := w.Write(ss.css.Bytes())
		return err
	}

	// Generate the CSS for the global styles.
	for _, style := range slices.Concat(
		ss.Resets(),
		ss.Base(),
		ss.Buttons(),
	) {
		if err := style.CSS(&ss.css); err != nil {
			return err
		}
	}
	// Generate the CSS for the themes.
	for _, theme := range ss.Themes {
		if err := theme.CSS(&ss.css); err != nil {
			return err
		}
	}
	// Write the CSS to the writer.
	_, err := w.Write(ss.css.Bytes())
	return err
}

// NewStyleSheet returns a new stylesheet.
func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		Themes: []*Theme{
			{
				MediaQuery:        "@media (prefers-color-scheme: light)",
				Background:        variables.White,
				Foreground:        variables.Neutral900,
				Primary:           variables.Neutral900,
				PrimaryForeground: variables.White,
			},
			{
				MediaQuery:        "@media (prefers-color-scheme: dark)",
				Background:        variables.Neutral900,
				Foreground:        variables.Neutral100,
				Primary:           variables.White,
				PrimaryForeground: variables.Neutral900,
			},
		},
	}
}
