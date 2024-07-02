package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/react"
	"github.com/snusEbjoer/cli-react/reactlib"
)

type Input struct {
	View   string
	State  *reactlib.State
	React  *react.React
	Events []reactlib.Event
}

func New(stateValue any, key string, react *react.React) *Input {
	input := &Input{
		View:  "",
		State: react.UseState(stateValue, key),
		React: react,
	}
	input.init()
	return input
}

func (i *Input) GetKey() string {
	return i.State.Key
}
func (i *Input) GetState() *reactlib.State {
	return i.State
}

func (i *Input) init() {
	al := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range al {
		keyEvent := reactlib.RuneKey(string(char))

		i.State.AddHandler(keyEvent, func(a ...any) {
			i.State.SetState(i.State.Curr.(string) + string(char))
		})

		i.Events = append(i.Events, keyEvent)
	}

	i.State.AddHandler(reactlib.NdaKey(keyboard.KeyBackspace2), func(a ...any) {
		curr := []rune(i.State.Curr.(string))
		if len(curr) == 0 {
			return
		}
		i.State.SetState(string(curr[0 : len(curr)-1]))
	})

	i.Events = append(i.Events, reactlib.NdaKey(keyboard.KeyBackspace2))
}

func (i *Input) Render() string {
	return fmt.Sprintf("Input: %s ", i.State.Curr)
}
