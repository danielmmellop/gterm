package command

import (
	"gterm/internal/file"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		err := file.Create(input.Params)
		return CommandOutput{Result: []string{}, Error: err}
	}
	commands["touch"] = action
	commands["new"] = action
}
