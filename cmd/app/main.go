package main

import (
	"fmt"
	"time"

	"concurrent-bus/internal/app"
	"concurrent-bus/internal/infra"
)

func main() {
	now := time.Now().Format(time.StampMilli)
	fmt.Println(now, "main program started...")
	bus := infra.NewBus()

	cmd1 := app.Cmd1{}

	err := bus.Handle(cmd1)
	if err != nil {
		fmt.Printf("error from cmd1: %s\n", err)
	}

	fmt.Println()
	fmt.Println("reset of the application...")
}
