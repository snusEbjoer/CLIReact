package component

import (
	"fmt"
	"log"
	"strconv"

	"github.com/snusEbjoer/cli-react/react"
	"github.com/snusEbjoer/cli-react/reactlib"
)

type Component struct {
	View  string
	State *reactlib.State
	React *react.React
}

const RuneKey = 0

func New(stateValue any, key string, react *react.React) *Component {
	c := &Component{
		View:  "",
		State: react.UseState(stateValue, key),
		React: react,
	}
	c.init()
	return c
}
func (c *Component) GetKey() string {
	return c.State.Key
}
func (c *Component) GetState() *reactlib.State {
	return c.State
}

func (c *Component) init() {
	c.State.AddHandler(reactlib.RuneKey("r"), func(a ...any) {
		c.State.SetState(c.State.Curr.(int) + 1)
		c.View = strconv.Itoa(c.State.Curr.(int))
	})

	c.State.AddHandler(reactlib.RuneKey("q"), func(a ...any) {
		log.Fatal("exit")
	})
}

func (c *Component) Render() string {
	return fmt.Sprintf("Count: %d ", c.State.Curr)
}
