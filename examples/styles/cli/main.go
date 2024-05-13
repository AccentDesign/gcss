package main

import (
	"flag"
	"fmt"
	"github.com/AccentDesign/gcss/examples/styles"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "export":
		commandExport()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: <command> [options]")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("  export")
}

func commandExport() {
	cmd := flag.NewFlagSet("export", flag.ExitOnError)
	help := cmd.Bool("help", false, "Display help for the export command")
	directory := cmd.String("directory", ".", "The directory to export the CSS files to.")

	if err := cmd.Parse(os.Args[2:]); err != nil {
		log.Fatalf("Error parsing flags: %v\n", err)
	}

	if *help {
		cmd.Usage()
		return
	}
	if *directory == "" {
		flag.Usage()
		log.Fatalln("The -directory flag is required.")
	}

	if err := checkDirectory(*directory); err != nil {
		log.Fatalf("Directory error: %v\n", err)
	}

	exportStyles(*directory)
}

func checkDirectory(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return fmt.Errorf("the directory does not exist: %s", directory)
	}
	return nil
}

func exportStyles(directory string) {
	for name, style := range styles.Styles {
		filePath := filepath.Join(directory, fmt.Sprintf("%s.css", name))
		file, err := os.Create(filePath)
		if err != nil {
			log.Printf("Failed to create file %s: %v", filePath, err)
			continue
		}
		defer file.Close()

		for _, s := range style {
			if err := s.CSS(file); err != nil {
				log.Printf("Failed to write CSS for style %s: %v", name, err)
				break
			}
		}
		log.Println("Exported", name, "to", file.Name())
	}
}
