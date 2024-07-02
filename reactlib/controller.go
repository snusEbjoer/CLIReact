package reactlib

import (
	"sync"

	"github.com/eiannone/keyboard"
)

type Controller struct {
	ControllerEvents map[Event]struct{}
	Subs             []Sub
	EventsChan       chan Event
	CurrFocusedState *State
	mu               *sync.Mutex
}

func NewController() *Controller {
	c := &Controller{
		ControllerEvents: make(map[Event]struct{}),
		Subs:             []Sub{},
		EventsChan:       make(chan Event),
		mu:               &sync.Mutex{},
	}
	c.run()
	return c
}

func (c *Controller) Init() {
	c.CurrFocusedState.AddHandler(NdaKey(keyboard.KeyArrowRight),
		func(a ...any) {
			if c.CurrFocusedState.Curr == len(c.Subs)-1 {
				c.CurrFocusedState.SetState(0)
			} else {
				c.CurrFocusedState.SetState(c.CurrFocusedState.Curr.(int) + 1)
			}
		},
		[]any{},
	)
	c.CurrFocusedState.AddHandler(NdaKey(keyboard.KeyArrowLeft),
		func(a ...any) {
			if c.CurrFocusedState.Curr == 0 {
				c.CurrFocusedState.SetState(len(c.Subs) - 1)
			} else {
				c.CurrFocusedState.SetState(c.CurrFocusedState.Curr.(int) - 1)
			}
		},
		[]any{},
	)
	c.AddEvent(NdaKey(keyboard.KeyArrowRight))
	c.AddEvent(NdaKey(keyboard.KeyArrowLeft))
}

func (c *Controller) AddEvent(event Event) {
	c.ControllerEvents[event] = struct{}{}
}
func (c *Controller) IsInControllerEvents(event Event) bool {
	_, ok := c.ControllerEvents[event]
	return ok
}
func (c *Controller) Subscribe(sub Sub) {
	c.Subs = append(c.Subs, sub)
}
func (c *Controller) isInSubEvents(event Event, sub Sub) bool {
	_, ok := sub.Events[event]
	return ok
}

func (c *Controller) run() {
	go func() {
		for {
			for e := range c.EventsChan {
				if c.IsInControllerEvents(e) {
					c.CurrFocusedState.Events <- e
				} else {
					sub := c.Subs[c.CurrFocusedState.Curr.(int)]
					if c.isInSubEvents(e, sub) {
						sub.Chan <- e
					}

				}
			}
		}
	}()
}
