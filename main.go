package main

import (
	"github.com/jjonline/golang-backend/app/console/command"
	"github.com/jjonline/golang-backend/client/initializer"
	"github.com/jjonline/golang-backend/conf"
	_ "go.uber.org/automaxprocs"
)

func init() {
	// step1、init config
	conf.Init()

	// step2、init global client handle
	initializer.Init()
}

func main() {
	command.Start()
}
