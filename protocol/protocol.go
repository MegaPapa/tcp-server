package protocol

import "errors"

/**
	Protocol definition:
	CMD_NUMBER: 16 bytes
	--------------------------
	DATA_SIZE(bytes): 16 bytes
	DATA: n bytes
	--------------------------
	ARGS_LENGTH: 16 bytes
	CMD_ARGS: m bytes
	--------------------------
 */

const(
	PING_CMD = 1
	SEND_FILE_CMD = 2
	LS_CMD = 3
	CD_CMD = 4
	LOAD_FILE_CMD = 5
	EXIT_CMD = 6
)

func PrepareCommand(command string, args []byte, data []byte) (Command, error) {
	executableCommand := ServerCommand{nil, len(data), data, len(args), args}
	switch command {
		case "ping":
			executableCommand.CmdType = PING_CMD
			break;
		case "send":
			executableCommand.CmdType = SEND_FILE_CMD
			break;
		case "ls":
			executableCommand.CmdType = LS_CMD
			break;
		case "cd":
			executableCommand.CmdType = CD_CMD
			break;
		case "load":
			executableCommand.CmdType = LOAD_FILE_CMD
			break;
		case "exit":
			executableCommand.CmdType = EXIT_CMD
			break;
		default:
			return nil, errors.New("Command not found: " + command)
	}
	return &executableCommand, nil
}