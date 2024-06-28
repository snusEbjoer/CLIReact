package component

import (
	"fmt"
	"log"
	"strconv"

	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/pkg/types"
	"github.com/snusEbjoer/cli-react/pkg/utils"
	"github.com/snusEbjoer/cli-react/state"
)

type Component struct {
	View     string
	State    *state.State
	Enveloup *enveloup.Enveloup
}

const RuneKey = 0

func New(stateValue any, key string, envelope *enveloup.Enveloup) *Component {
	c := &Component{
		View:     "",
		State:    state.New(stateValue, key),
		Enveloup: envelope,
	}
	c.init()
	return c
}
func (c *Component) GetKey() string {
	return c.State.Key
}
func (c *Component) GetState() *state.State {
	return c.State
}

func (c *Component) init() {
	c.State.AddHandler(utils.RuneKey("r"), func(a ...any) {
		c.State.SetState(c.State.Curr.(int) + 1)
		c.View = strconv.Itoa(c.State.Curr.(int))
	})

	c.State.AddHandler(utils.RuneKey("q"), func(a ...any) {
		log.Fatal("exit")
	})
	c.Enveloup.Subscribe([]types.Event{utils.RuneKey("r"), utils.RuneKey("q")}, c.State) // TODO move events to envelope
}

func (c *Component) Render() string {
	return fmt.Sprintf("Count: %d ", c.State.Curr)
}
