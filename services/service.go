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

type Service struct{
	Node Node
	Label string
	Index int
}

type Node interface {
	GetRegistrar() Node
	GetServiceChannel() chan ServiceNotification
	GetAddress() string
	GetRoleLabel() string
	RegisterService(*Service)
}

func Bootstrap(node Node) {
	tunnel := Service(NewTunnel(node))
	node.RegisterService(&tunnel)
}
