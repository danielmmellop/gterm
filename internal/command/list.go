package command

import (
	"gterm/internal/directory"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		result, err := directory.List(input.Params)
		return CommandOutput{Result: result, Error: err}
	}

	commands["dir"] = action
	commands["ls"] = action
}
