package in

type mockServer struct{
	MsgHandler Receiver
}

type MockServer interface{
	Server
}
func ServerMock() MockServer {
	return &mockServer{}
}

func (s *mockServer) Serve() {

}