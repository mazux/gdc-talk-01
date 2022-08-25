package main

import (
	"fmt"
	"time"

	"concurrent-bus/internal/app"
	"concurrent-bus/internal/infra"
)

const maxRetries = 3

func main() {
	now := time.Now().Format(time.StampMilli)
	fmt.Println(now, "main program started...")

	bus := &infra.Bus{}
	cmds := []interface{}{
		app.Cmd1{},
		app.Cmd2{},
		app.Cmd1{},
		app.Cmd2{},
	}

	cbus := infra.NewConcurrentBus(bus, cmds)
	cbus.Handle()
	strategy := infra.LinearBackoffStrategy()
	cbus.RetryWithStrategy(strategy, maxRetries)

	fmt.Println()
	fmt.Println("reset of the application...")
}
