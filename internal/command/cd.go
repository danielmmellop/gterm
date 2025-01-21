package command

import (
	"gterm/internal/directory"
)

func init() {
	commands["cd"] = func(input CommandInput) CommandOutput {
		err := directory.Change(input.Params)
		return CommandOutput{Result: []string{}, Error: err}
	}
}
