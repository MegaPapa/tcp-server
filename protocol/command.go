package protocol

import "net"

type Command interface {
	ExecuteCommand(connection net.Conn)
}

type ServerCommand struct {
	CmdType int64
	CmdData []byte
}

func (cmd *ServerCommand) ExecuteCommand(connection net.Conn) {

}
