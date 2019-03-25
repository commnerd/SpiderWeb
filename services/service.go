package services

type ServiceStatus int

const (
	ServiceStatusInit ServiceStatus = iota
	ServiceStatusRunning
	ServiceStatusDead
	ServiceStatusKilled
)

type Service interface{
	GetServiceLabel() string
	GetStatus() ServiceStatus
	Run()
}

type Node interface {
	GetRegistrar() Node
	GetAddress() string
	GetRole() string
	RegisterService(Service)
}

type ServiceStruct struct{
	Node Node
	Label string
	Status ServiceStatus
}

func Bootstrap(node Node) {
	service := NewTunnel(node)
	node.RegisterService(service)
}