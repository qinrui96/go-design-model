package pkg

import "fmt"

type singleton struct {
}

var Singleton *singleton

func init() {
	Singleton = new(singleton)
}

func GetSingleton() *singleton {
	return Singleton
}

func (s *singleton) Show() {
	fmt.Println("i'm a singleton")
}
