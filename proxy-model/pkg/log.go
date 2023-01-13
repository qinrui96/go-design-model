package pkg

import (
	"fmt"
	"log"
)

type Logger struct {
	algorithm *Algorithm
}

func (l *Logger) Add(a, b int) int {
	str := fmt.Sprintf("operand a= %d, operand b=%d", a, b)
	log.Println(str)
	return l.algorithm.Add(a, b)
}

func NewLogger(algorithm *Algorithm) *Logger {
	return &Logger{
		algorithm: algorithm,
	}
}
