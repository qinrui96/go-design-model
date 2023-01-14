package pkg

import "fmt"

type TiDB struct {
}

func (t *TiDB) WriteTiDB() {
	fmt.Println("write TiDB")
}

func NewTiDB() *TiDB {
	return &TiDB{}
}
