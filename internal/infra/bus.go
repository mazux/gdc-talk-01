package infra

import (
	"fmt"
	"reflect"

	"concurrent-bus/internal/app"
)

type Bus struct {
	hndlr1 app.Hndlr1
	hndlr2 app.Hndlr2
}

// typical approach
func (b *Bus) Handle(cmd interface{}) error {
	cmdType := getType(cmd)
	switch cmdType {
	case "Cmd1":
		return b.hndlr1.Handle(cmd.(app.Cmd1))
	case "Cmd2":
		return b.hndlr2.Handle(cmd.(app.Cmd2))
	}

	return fmt.Errorf("unable to find handler for cmd %s", cmdType)
}

func getType(v interface{}) string {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
