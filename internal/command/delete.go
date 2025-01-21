package command

import (
	"gterm/internal/directory"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		err := directory.Delete(input.Params)
		return CommandOutput{Result: []string{}, Error: err}
	}

	commands["rm"] = action
	commands["remove"] = action
	commands["del"] = action
	commands["delete"] = action
}
