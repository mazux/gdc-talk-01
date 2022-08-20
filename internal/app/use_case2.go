package app

import (
	"fmt"
	"time"
)

type Cmd2 struct {
	// ... fields
}

type Hndlr2 struct {
	// ... dependencies
}

func (h *Hndlr2) Handle(cmd Cmd2) error {
	// memic the handling period using sleep
	time.Sleep(time.Second)

	now := time.Now().Format(time.StampMilli)
	fmt.Println(now, "Use-Case 2, handling command normally... with error")

	return fmt.Errorf("error returned from Hndlr2")
}
