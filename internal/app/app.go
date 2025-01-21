package app

import (
	"gterm/internal/os"
	"gterm/internal/theme"
	"gterm/internal/ui/display"
	"gterm/internal/ui/footer"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// TODO: Keep focus only on the InputCmd
type App struct {
	name   string
	Height float32
	Width  float32
}

type AppConfig struct {
	Name   string
	Height float32
	Width  float32
}

func New(config AppConfig) App {
	os.SetCurrentFolder()
	return App{
		name:   config.Name,
		Height: config.Height,
		Width:  config.Width,
	}
}

func (gTermApp *App) ShowAndRun() {
	application := app.New()
	application.Settings().SetTheme(theme.New())

	mainWindow := application.NewWindow(gTermApp.name)
	footer := footer.New()
	display := display.New(gTermApp.Width, gTermApp.Height, func() {
		mainWindow.Canvas().Focus(footer.Input)
	})
	footer.Input.SetTarget(display.Grid)
	mainWindow.SetContent(container.NewVBox(display.Container, footer.Container))

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(gTermApp.Width, gTermApp.Height))
	mainWindow.Canvas().Focus(footer.Input)
	mainWindow.ShowAndRun()
}
