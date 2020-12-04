package message_bus

type bus struct{
	channels map[string]Channel
}

var instance *bus

func init() {
	instance = &bus{
		channels: make(map[string]Channel, 0),
	}
}