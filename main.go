package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/AccentDesign/gostyle/examples"
	"github.com/AccentDesign/gostyle/style"
)

func main() {
	defer func(start time.Time) {
		fmt.Println("CSS and JSON files created successfully!")
		fmt.Println("Execution time:", time.Since(start))
	}(time.Now())

	all := map[string][]style.Style{
		"alerts":  examples.Alerts,
		"badges":  examples.Badges,
		"buttons": examples.Buttons,
		"form":    examples.Form,
		"tables":  examples.Tables,
	}

	for name, styles := range all {
		exportCSS(styles, name)
		exportJSON(styles, name)
	}
}

func exportCSS(styles []style.Style, name string) {
	f, fErr := os.Create(fmt.Sprintf("docs/css/%s.css", name))
	if fErr != nil {
		panic(fErr)
	}
	for _, s := range styles {
		if err := s.CSS(f); err != nil {
			panic(err)
		}
		if _, err := fmt.Fprintln(f, ""); err != nil {
			panic(err)
		}
	}
}

func exportJSON(styles []style.Style, name string) {
	j, err := json.MarshalIndent(styles, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	f, fErr := os.Create(fmt.Sprintf("json/%s.json", name))
	if fErr != nil {
		panic(fErr)
	}
	if _, err := fmt.Fprintln(f, string(j)); err != nil {
		panic(err)
	}
}
