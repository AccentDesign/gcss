package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AccentDesign/gostyle/examples"
	"github.com/AccentDesign/gostyle/style"
)

func main() {
	defer func(start time.Time) {
		fmt.Println("CSS file created successfully!")
		fmt.Println("Execution time:", time.Since(start))
	}(time.Now())

	all := map[string][]style.Style{
		"buttons": examples.Buttons,
		"form":    examples.Form,
		"tables":  examples.Tables,
	}

	// Print CSS examples
	for name, styles := range all {
		f, fErr := os.Create(fmt.Sprintf("css/%s.css", name))
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
}
