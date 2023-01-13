//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qinrui/design-model/proxy-model/pkg"
)

func NewLogger() *pkg.Logger {
	wire.Build(pkg.NewAlgorithm, pkg.NewLogger)
	return &pkg.Logger{}
}
