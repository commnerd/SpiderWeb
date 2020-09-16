package msg

type Message int

const (
    JoinRequest Message = iota
)

func (m Message) String() string {
    return [...]string{
		"Join Request",
	}[m]
}