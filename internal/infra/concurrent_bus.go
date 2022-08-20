package infra

import (
	"fmt"
)

type ConcurrentBus struct {
	bus     *Bus
	cmds    []interface{}
	errChan chan error
}

func (b *ConcurrentBus) Handle() {
	for _, cmd := range b.cmds {
		go b.handle(cmd)
	}

	for i := 0; i < len(b.cmds); i++ {
		if err := <-b.errChan; err != nil {
			fmt.Println("error: ", err)
		}
	}
}

func (b *ConcurrentBus) handle(cmd interface{}) {
	b.errChan <- b.bus.Handle(cmd)
}

func NewConcurrentBus(bus *Bus, cmds []interface{}) *ConcurrentBus {
	errChan := make(chan error, len(cmds))
	return &ConcurrentBus{bus, cmds, errChan}
}
