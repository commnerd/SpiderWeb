package service

type mockManager struct{}

func MockManager() Manager {
	return &mockManager{}
}
