package command

import (
	"gterm/internal/file"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		err := file.Move(input.Params, false)
		return CommandOutput{Result: []string{""}, Error: err}
	}
	commands["copy"] = action
	commands["cp"] = action
}
