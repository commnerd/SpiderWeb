package message_bus

func Send(channel string, msg interface{}){
	targetChannel, err := GetChannel(channel)

	if err != nil {
		panic("Something went wrong in message_bus send.go")
	}

	targetChannel.Send(msg)
}