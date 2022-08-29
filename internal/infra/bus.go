package infra

import (
	"fmt"
	"reflect"
	"sync"

	"concurrent-bus/internal/app"
)

type bus struct {
	// domain bus to transmit domain events
	domainBus chan interface{}
	errChan   chan error

	eventCmds []interface{}

	hndlr1 app.Hndlr1
	hndlr2 app.Hndlr2
}

func NewBus() bus {
	domainBus := make(chan interface{})
	errChan := make(chan error)

	return bus{
		domainBus,
		errChan,
		[]interface{}{},
		app.Hndlr1{},
		app.Hndlr2{},
	}
}

func (b *bus) Handle(cmd interface{}) error {
	go b.handle(cmd)

	for {
		select {
		case err := <-b.errChan:
			if err != nil {
				return err
			}
			return b.handleDomainEvents()
		case domainEvent := <-b.domainBus:
			b.registerEventAsCommand(domainEvent)
		}
	}
}

func (b *bus) handleDomainEvents() error {
	var wg sync.WaitGroup
	wg.Add(len(b.eventCmds))
	for _, cmd := range b.eventCmds {
		go func(c interface{}) {
			defer wg.Done()
			b.Handle(c)
		}(cmd)
	}
	wg.Wait()
	return nil
}

func (b *bus) registerEventAsCommand(domainEvent interface{}) {
	eventType := getType(domainEvent)
	fmt.Println("registering event with type ", eventType)

	b.eventCmds = append(b.eventCmds, app.Cmd2{})
	b.eventCmds = append(b.eventCmds, app.Cmd2{})
	b.eventCmds = append(b.eventCmds, app.Cmd2{})
	// b.eventCmds = append(b.eventCmds, app.Cmd1{})
}

func (b *bus) handle(cmd interface{}) {
	cmdType := getType(cmd)
	switch cmdType {
	case "Cmd1":
		b.errChan <- b.hndlr1.Handle(cmd.(app.Cmd1), b.domainBus)
		return
	case "Cmd2":
		b.errChan <- b.hndlr2.Handle(cmd.(app.Cmd2))
		return
	}

	b.errChan <- fmt.Errorf("unable to find handler for cmd %s", cmdType)
}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
