package service

type Manager interface{}

type manager struct{
	Services []Service
}

func NewManager() Manager {
	return &manager{}
}