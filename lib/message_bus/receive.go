package message_bus

func Receive(channel string, params ...interface{}) interface{} {


	targetChannel, err := GetChannel(channel)

	if err != nil {
		panic("Something went wrong in message_bus receive.go")
	}

	return targetChannel.Receive(params...)
}