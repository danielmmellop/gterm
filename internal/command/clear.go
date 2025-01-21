package command

import (
	"fmt"
	"time"
)

var CLEAR_IDENTIFY = fmt.Sprintf("clear%vclear", time.Now().UnixNano())

func init() {
	commands["clear"] = func(input CommandInput) CommandOutput {
		return CommandOutput{Result: []string{CLEAR_IDENTIFY}, Error: nil}
	}
}
