package app

import (
	"fmt"
	"time"
)

type Cmd1 struct {
	// ... fields
}

type Hndlr1 struct {
	// ... dependencies
}

func (h *Hndlr1) Handle(cmd Cmd1) error {
	// memic the handling period using sleep
	time.Sleep(time.Second)

	now := time.Now().Format(time.StampMilli)
	fmt.Println(now, "Use-Case 1, handling command normally...")

	return nil
}
