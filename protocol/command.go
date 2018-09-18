package protocol

import (
	"net"
	"strings"
)

type Command interface {
	WriteCommand(connection net.Conn)
}

type ServerCommand struct {
	CmdType int
	DataLength int
	CmdData []byte
	ArgsLength int
	CmdArgs []byte
}

func (cmd *ServerCommand) WriteCommand(connection net.Conn) {

}

/**
	Translate user input to tokens, which will be used in protocol compression
	Typical input structure: command arg1, arg2, ...
 */
func TokenizeCommand(input string) (command string, args []string) {
	input = strings.TrimSpace(input)
	tokens := strings.Split(input, " ")
	command = tokens[0]
	argsCount := len(tokens) - 1
	args = make([]string, argsCount)
	for i := 1; i < len(tokens); i++ {
		args[i - 1] = tokens[i]
	}
	return command, args
}
