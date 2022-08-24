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
	for i := 1; i <= maxRetries; i++ {
		cbus.Retry(time.Duration(i) * time.Second)
	}

	fmt.Println()
	fmt.Println("reset of the application...")
}
