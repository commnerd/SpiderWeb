package message_bus

type Bus struct{
	Receivers map[string]Receiver
	Senders map[string]Sender
}

var instance *Bus

func init() {
	instance = New()
}

func New() *Bus {
	return &Bus{
		Receivers: make(map[string]Receiver),
		Senders: make(map[string]Sender),
	}
}

func Get() *Bus {
	return instance
}

func Register(item interface{}) {
	instance.Register(item)
}

func (b *Bus) Register(item interface{}) {
	r, isReceiver := item.(Receiver)
	s, isSender := item.(Sender)

	if isReceiver {
		b.Receivers[r.GetLabel()] = r
	}

	if isSender {
		b.Senders[s.GetLabel()] = s
	}
}