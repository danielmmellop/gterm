package command

import (
	"gterm/internal/os"
)

func init() {
	commands["env"] = func(input CommandInput) CommandOutput {
		results, err := os.Env(input.Params)
		return CommandOutput{Result: results, Error: err}
	}
}
