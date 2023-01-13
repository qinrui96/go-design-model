package main

import "fmt"

func main() {
	logger := NewLogger()
	ret := logger.Add(3, 5)
	fmt.Println(ret)
}
