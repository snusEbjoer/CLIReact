package reactlib

import "github.com/eiannone/keyboard"

type Event struct {
	Key  keyboard.Key
	Char string
}

type Sub struct {
	Events map[Event]struct{}
	Chan   chan Event
}
