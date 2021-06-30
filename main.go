package main

import (
	"github.com/jjonline/golang-backend/app/console/command"
	_ "go.uber.org/automaxprocs"
)

func main() {
	command.Start()
}
