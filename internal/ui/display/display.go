package display

import (
	"gterm/internal/ui/grid"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Display struct {
	Container *fyne.Container
	Grid      *grid.Grid
}

func New(minWidth, minHeight float32, customTapEvent func()) *Display {
	grid := grid.New(customTapEvent)
	textGridContainer := container.NewScroll(grid)
	textGridContainer.SetMinSize(fyne.NewSize(minWidth, minHeight))

	return &Display{
		Grid:      grid,
		Container: container.NewGridWithRows(1, textGridContainer),
	}

}
