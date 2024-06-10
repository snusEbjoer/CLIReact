package utils

import (
	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/pkg/types"
)

func RuneKey(char string) types.Event {
	return types.Event{Key: 0, Char: char}
}

func NdaKey(key keyboard.Key) types.Event {
	return types.Event{Key: key, Char: "\x00"}
}
