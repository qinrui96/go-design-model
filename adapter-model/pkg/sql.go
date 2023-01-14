package pkg

import "fmt"

type Sql struct {
}

func (s *Sql) WriteSql() {
	fmt.Println("write sql")
}

func NewSql() *Sql {
	return &Sql{}
}
