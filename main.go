package main

import "fmt"

type Message struct {
	message string
}

func (m Message) getMessage() string {
	return m.message
}

func main() {
	message := Message{message: "Some test message"}

	fmt.Println(message.getMessage())
}
