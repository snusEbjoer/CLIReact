package input

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/pkg/types"
	"github.com/snusEbjoer/cli-react/pkg/utils"
	"github.com/snusEbjoer/cli-react/state"
)

type Input struct {
	View     string
	State    *state.State
	Enveloup *enveloup.Enveloup
	Events   []types.Event
}

func New(stateValue any, key string, envelope *enveloup.Enveloup) *Input {
	input := &Input{
		View:     "",
		State:    state.New(stateValue, key),
		Enveloup: envelope,
	}
	input.init()
	return input
}

func (i *Input) GetKey() string {
	return i.State.Key
}
func (i *Input) GetState() *state.State {
	return i.State
}

func (i *Input) init() {
	al := "abcdefghijklmnopqrstuvwxyz"
	for _, char := range al {
		keyEvent := utils.RuneKey(string(char))

		i.State.AddHandler(keyEvent, func(a ...any) {
			i.State.SetState(i.State.Curr.(string) + string(char))
		})

		i.Events = append(i.Events, keyEvent)
	}

	i.State.AddHandler(utils.NdaKey(keyboard.KeyBackspace2), func(a ...any) {
		curr := []rune(i.State.Curr.(string))
		if len(curr) == 0 {
			return
		}
		i.State.SetState(string(curr[0 : len(curr)-1]))
	})

	i.Events = append(i.Events, utils.NdaKey(keyboard.KeyBackspace2))
	i.Enveloup.Subscribe(i.Events, i.State)
}

func (i *Input) Render() string {
	return fmt.Sprintf("Input: %s ", i.State.Curr)
}
