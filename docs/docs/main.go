package main

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/docs/docs/examples"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Println("CSS and JSON files created successfully!")
		fmt.Println("Execution time:", time.Since(start))
	}(time.Now())

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dir := filepath.Dir(filename)
	cssDir := filepath.Join(dir, "..", "css")

	all := map[string][]gcss.Style{
		"alerts":  examples.Alerts,
		"badges":  examples.Badges,
		"buttons": examples.Buttons,
		"form":    examples.Form,
		"tables":  examples.Tables,
	}

	for name, styles := range all {
		exportCSS(styles, name, cssDir)
	}
}

func exportCSS(styles []gcss.Style, name string, dir string) {
	cssPath := filepath.Join(dir, fmt.Sprintf("%s.css", name))
	fmt.Println("Exporting:", cssPath)
	f, err := os.Create(cssPath)
	if err != nil {
		panic(err)
	}
	for _, s := range styles {
		if err := s.CSS(f); err != nil {
			panic(err)
		}
		// only for visual separation
		if _, err := fmt.Fprintln(f, ""); err != nil {
			panic(err)
		}
	}
}
