//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/qinrui/design-model/bridge-model/pkg/level"
	"github.com/qinrui/design-model/bridge-model/pkg/sender"
)

var MageSet = wire.NewSet(sender.NewMessageSender, sender.NewPhoneSender)

func NewWarningSender() level.ILevel {

	wire.Build(MageSet, wire.Bind(new(sender.Sender), new(sender.PhoneSender)), level.NewWarning)
	return &level.Log{}
}
