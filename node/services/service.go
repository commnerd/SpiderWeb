package services

type Service interface{
    GetLabel() string
    Run(Node)
    IsRunning() boolean
}
