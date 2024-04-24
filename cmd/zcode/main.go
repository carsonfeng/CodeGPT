package main

import (
	"github.com/carsonfeng/ZCode/cmd"

	"github.com/appleboy/graceful"
)

func main() {
	m := graceful.NewManager()
	cmd.Execute(m.ShutdownContext())
}
