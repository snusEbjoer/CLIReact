package reactlib

import (
	"github.com/eiannone/keyboard"
)

func RuneKey(char string) Event {
	return Event{Key: 0, Char: char}
}

func NdaKey(key keyboard.Key) Event {
	return Event{Key: key, Char: "\x00"}
}
