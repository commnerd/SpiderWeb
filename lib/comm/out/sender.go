package out

type sender struct{}

type Sender interface{}

func NewSender() Sender {
	return &sender{}
}