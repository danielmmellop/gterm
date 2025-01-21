package command

import (
	"gterm/internal/directory"
)

func init() {
	commands["mkdir"] = func(input CommandInput) CommandOutput {
		err := directory.Create(input.Params)
		return CommandOutput{Result: []string{}, Error: err}
	}
}
