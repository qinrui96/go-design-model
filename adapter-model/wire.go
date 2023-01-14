//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qinrui/go-design-model/adapter-model/pkg"
)

var MageDB = wire.NewSet(pkg.NewSql, pkg.NewTiDB)

func NewDbAdapter() *pkg.TiDBAdapter {
	wire.Build(MageDB, pkg.NewTiDBAdapter)
	return &pkg.TiDBAdapter{}
}
