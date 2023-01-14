package main

import (
	"fmt"
	"github.com/qinrui/go-design-model/combine-model/pkg"
)

func main() {
	folder := pkg.NewFolder()
	fmt.Println(folder.Count())
}
