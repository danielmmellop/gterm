package input

import (
	"gterm/internal/command"
	"gterm/internal/ui/grid"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Input struct {
	widget.Entry
	target *grid.Grid
}

func New() *Input {
	input := &Input{}
	input.ExtendBaseWidget(input)
	input.MultiLine = false

	input.OnSubmitted = func(cmd string) {
		if input.target == nil {
			return
		}

		newCommand := command.NewCommand(cmd)
		newCommand.Run()
		if len(newCommand.Result) == 2 && newCommand.Result[1] == command.CLEAR_IDENTIFY {
			input.target.Clean()
		} else {
			input.target.UpdateContentText(newCommand.Result)
		}

		input.SetText("")
	}

	return input
}

func (input *Input) SetTarget(target *grid.Grid) {
	input.target = target
}

func (input *Input) KeyUp(k *fyne.KeyEvent) {
	if k.Name == "Up" {
		input.SetText(command.History[len(command.History)-1])
	}
}
