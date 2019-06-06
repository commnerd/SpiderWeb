package services

type manager struct{}

type PortManager interface{
    NextOpen() int

}

func NewPortManager() manager {
    return manager{}
}
