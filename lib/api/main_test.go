package api

import (
	"github.com/google/uuid"
	"testing"
	"../id"
)

func TestMain(m *testing.M) {

	RunningNode = &mockNodeStruct{}

	HomeStruct.Address = id.Id(uuid.MustParse("322a1963-2b7f-43d4-b9cf-2fcea27c63da")).String()
	HomeStruct.Ip = "192.168.10.10"
	HomeStruct.Port = 42

	m.Run()
}