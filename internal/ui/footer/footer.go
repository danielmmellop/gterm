package footer

import (
	"gterm/internal/ui/input"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Footer struct {
	Container *fyne.Container
	Input     *input.Input
}

// TODO: Fix the padding between the elements
func New() Footer {
	input := input.New()

	return Footer{
		Container: container.NewVBox(input),
		Input:     input,
	}
}
