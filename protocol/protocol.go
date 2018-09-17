package protocol

/**
	Protocol definition:
	CMD_NUMBER: 16 bytes
	DATA: n bytes
 */

const(
	TEXT_MESSAGE_CMD = 1
	SEND_FILE_CMD = 2
)

func PrepareCommand(command int64, data []byte) (Command, error) {

	return nil, nil
}