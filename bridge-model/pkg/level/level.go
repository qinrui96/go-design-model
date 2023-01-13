package level

import (
	"fmt"
	"github.com/qinrui/design-model/bridge-model/pkg/sender"
)

type ILevel interface {
	ShowLevel(str string)
}

type Warning struct {
	sender sender.Sender
}

func NewWarning(sender sender.Sender) ILevel {
	return &Warning{sender: sender}
}

func (w *Warning) ShowLevel(str string) {
	fmt.Print("Level: warning, ")
	w.sender.Send(str)
}

type Log struct {
	sender sender.Sender
}

func NewLog(sender2 sender.Sender) ILevel {
	return &Log{sender: sender2}
}

func (l *Log) ShowLevel(str string) {
	fmt.Print("Level: log, ")
	l.sender.Send(str)
}
