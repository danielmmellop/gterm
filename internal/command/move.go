package command

import (
	"gterm/internal/file"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		err := file.Move(input.Params, true)
		return CommandOutput{Result: []string{""}, Error: err}
	}
	commands["move"] = action
	commands["mv"] = action
}
