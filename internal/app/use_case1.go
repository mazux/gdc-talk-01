package app

import (
	"concurrent-bus/internal/domain"
	"fmt"
	"time"
)

var counter int

func init() {
	counter = 2
}

type Cmd1 struct {
	// ... fields
}

type Hndlr1 struct {
	// ... dependencies
}

func (h *Hndlr1) Handle(cmd Cmd1, domainBus chan<- interface{}) error {
	now := time.Now().Format(time.StampMilli)
	fmt.Println(now, "Use-Case 1, handling command normally...")

	if counter > 0 {
		mutatedAgg1, err := domain.MutateAggregate1("foo", 5, domainBus)
		if err != nil {
			return err
		}

		// do something with the aggregate
		// memic the handling period using sleep
		fmt.Printf("domain event emittid from aggregate %v, while counter is %d", mutatedAgg1, counter)
		fmt.Println()
	}

	counter--
	time.Sleep(time.Second)

	return nil
}
