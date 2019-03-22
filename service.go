package main

type SericeStatus int

const (
	ServiceStatusInit ServiceStatus = iota
	ServiceStatusRunning
	ServiceStatusDead
	ServiceStatusKilled
)

type Service interface{
	GetServiceLabel(): string
	GetStatus(): ServiceStatus
	Run()
}