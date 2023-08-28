package main

import (
	"linklifyBackend/internal/pkg/app"
)

func main() {
	app := app.New()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
