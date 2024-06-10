package types

import "github.com/eiannone/keyboard"

type Event struct {
	Key  keyboard.Key
	Char string
}
