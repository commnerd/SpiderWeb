package services

type ServiceEvent int

const (
	ServiceInitialized ServiceEvent = iota
	ServiceRunning
	ServiceDied
	ServiceKilled
)

type ServiceNotification struct {
	Service Service
	Event ServiceEvent
}

type ServiceStruct struct {
	Node Node
	Label string
	Index int
}

type Service interface {
	GetNode() Node
	GetLabel() string
	GetIndex() int
	SetIndex(int)
	Run()
}

type Node interface {
	GetRegistrar() Node
	GetCommChannel() chan ServiceNotification
	GetAddress() string
	GetRole() string
	RegisterService(Service)
}

func Bootstrap(node Node) {
	tunnel := NewTunnel(node)
	node.RegisterService(tunnel)
}
