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

type Service interface{
	GetLabel() string
	GetNode() Node
	Run()
}

type Registry interface{
	GetEventChannel() chan ServiceNotification
	Add(Service)
	Remove(Service)
}

type Node interface {
	GetRegistrar() Node
	GetAddress() string
	GetRoleLabel() string
	GetServiceRegistry() Registry
}

func Bootstrap(node Node) {
	tunnel := NewTunnel(node)
	node.GetServiceRegistry().Add(tunnel)
}
