package command

import (
	"gterm/internal/file"
)

func init() {
	action := func(input CommandInput) CommandOutput {
		results, err := file.Read(input.Params)
		return CommandOutput{Result: results, Error: err}
	}
	commands["cat"] = action
	commands["dog"] = action // If you are a dog fan :)
}
