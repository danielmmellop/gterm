package command

import "gterm/internal/os"

func init() {
	commands["pwd"] = func(input CommandInput) CommandOutput {
		return CommandOutput{Result: []string{os.CurrentPath}, Error: nil}
	}
}
