package styles

import (
	"bytes"
	"github.com/AccentDesign/gcss"
	"io"
	"slices"
	"sync"
)

type (
	Styles     []gcss.Style
	StyleSheet struct {
		Media  []*Media
		Themes []*Theme
		css    bytes.Buffer
		mutex  sync.Mutex
	}
)

// CSS Writes the CSS for the stylesheet to the writer.
func (ss *StyleSheet) CSS(w io.Writer) error {
	ss.mutex.Lock()
	defer ss.mutex.Unlock()

	if ss.css.Len() > 0 {
		_, err := w.Write(ss.css.Bytes())
		return err
	}

	for _, style := range slices.Concat(
		ss.Resets(),
		ss.Layout(),
		ss.Buttons(),
	) {
		if err := style.CSS(&ss.css); err != nil {
			return err
		}
	}

	for _, media := range ss.Media {
		if err := media.CSS(&ss.css); err != nil {
			return err
		}
	}

	for _, theme := range ss.Themes {
		if err := theme.CSS(&ss.css); err != nil {
			return err
		}
	}

	_, err := w.Write(ss.css.Bytes())
	return err
}

// NewStyleSheet returns a new stylesheet. It includes the media queries and themes.
func NewStyleSheet() *StyleSheet {
	return &StyleSheet{
		Media: []*Media{
			{MediaType: Mobile, Query: "@media (max-width: 768px)"},
			{MediaType: Desktop, Query: "@media (min-width: 769px)"},
		},
		Themes: []*Theme{
			lightTheme,
			darkTheme,
		},
	}
}
