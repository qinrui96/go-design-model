package main

import "github.com/qinrui/design-model/singleton/pkg"

func main() {
	singleton := pkg.GetSingleton()
	singleton.Show()
}
