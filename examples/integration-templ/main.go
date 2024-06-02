package main

import (
	"context"
	"os"
)

func main() {
	file, err := os.Create("home.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	home := Page()
	home.Render(context.Background(), file)
}
