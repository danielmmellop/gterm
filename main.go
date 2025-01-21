package main

import (
	"gterm/internal/app"
)

func main() {
	appConfig := app.AppConfig{
		Name:   "gTerm",
		Width:  1024,
		Height: 768,
	}

	gterm := app.New(appConfig)
	gterm.ShowAndRun()
}
