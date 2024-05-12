package main

import (
	"fmt"
	"github.com/AccentDesign/gcss"
	"github.com/AccentDesign/gcss/examples/styles"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	exportCSS()
}

func exportCSS() {
	all := map[string][]gcss.Style{
		"alerts":  styles.Alerts,
		"badges":  styles.Badges,
		"buttons": styles.Buttons,
		"form":    styles.Form,
		"tables":  styles.Tables,
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	dir := filepath.Dir(filename)

	for name, style := range all {
		cssPath := filepath.Join(dir, fmt.Sprintf("static/css/%s.css", name))
		fmt.Println("Exporting:", cssPath)
		f, err := os.Create(cssPath)
		if err != nil {
			panic(err)
		}
		for _, s := range style {
			if err := s.CSS(f); err != nil {
				panic(err)
			}
		}
	}
}
