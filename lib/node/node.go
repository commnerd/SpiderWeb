package node

import (
	"../keys"
	"../id"
)

type Node struct{
	Id id.Id
	Mask id.Mask
	Ip string
	ChildCount uint64
	SshPort int
	SshKeys keys.Keys
	registrants map[id.Id]interface{}
	children map[id.Id]interface{}
}