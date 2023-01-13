package sender

import "fmt"

type Sender interface {
	Send(msg string)
}

type PhoneSender struct {
}

func NewPhoneSender() PhoneSender {
	return PhoneSender{}
}

func (p PhoneSender) Send(msg string) {
	fmt.Println(msg, " sent by phone")
}

type MessageSender struct {
}

func (m MessageSender) Send(msg string) {
	fmt.Println(msg, " sent by message")
}

func NewMessageSender() MessageSender {
	return MessageSender{}
}
