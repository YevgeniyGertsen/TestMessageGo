package usermessage

type Message struct {
	message string
}

func (m Message) getMessage() string {
	return m.message
}
