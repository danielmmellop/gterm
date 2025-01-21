package command

import (
	"fmt"
	"strings"
)

var commands = make(map[string]func(CommandInput) CommandOutput)
var History = []string{}

type CommandInput struct {
	Args   []string
	Params []string
}

type CommandOutput struct {
	Result []string
	Error  error
}

type Command struct {
	Name       string
	Raw        string   // The entire command string
	Parameters []string // cp . /source (param1) /dest (param2)
	Arguments  []string // cp -r (arg1) . (param1) /dest (param2)
	Result     []string
	Error      error
	isEmpty    bool
}

func NewCommand(cmd string) *Command {
	if len(cmd) == 0 {
		return &Command{isEmpty: true}
	}
	result := []string{}

	arguments, name, parameters, err := parse(cmd)
	if err != nil {
		result = []string{err.Error()}
	}

	History = append(History, cmd)

	return &Command{
		Name:       name,
		Raw:        cmd,
		Parameters: parameters,
		Arguments:  arguments,
		Result:     result,
	}
}

func (cmd *Command) Run() {
	if cmd.isEmpty {
		return
	}

	if len(cmd.Result) == 0 {
		command := commands[cmd.Name]
		output := command(CommandInput{cmd.Arguments, cmd.Parameters})
		if output.Error != nil {
			cmd.Result = []string{output.Error.Error()}
		} else {
			cmd.Result = output.Result
		}
	}
	cmd.Result = formatResult(cmd.Raw, cmd.Result)
}

func parse(rawCommand string) ([]string, string, []string, error) {
	var arguments, parameters []string

	split := strings.Split(strings.TrimSpace(rawCommand), " ")

	name := split[0]

	for _, v := range split[1:] {
		if strings.HasPrefix(v, "-") {
			arguments = append(arguments, v)
		} else {
			parameters = append(parameters, v)
		}
	}

	_, ok := commands[name]
	if !ok {
		return arguments, name, parameters, fmt.Errorf("command not found: %s", name)
	}

	return arguments, name, parameters, nil
}

func formatResult(rawCommand string, result []string) []string {
	if len(result) > 0 && len(result[0]) == 0 {
		return []string{"> " + rawCommand}
	}

	return append([]string{"> " + rawCommand}, result...)
}
