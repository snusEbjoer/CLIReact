package enveloup

import (
	"log"
	"sync"

	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/pkg/types"
	"github.com/snusEbjoer/cli-react/state"
)

type StateValue interface {
	GetChan() chan types.Event
	GetKey() string
}

type Sub struct {
	Events []types.Event
	Chan   chan types.Event
}

type Enveloup struct {
	events      chan types.Event
	subs        map[string]Sub
	controls    bool
	currFocused string
	mu          *sync.Mutex
}

func New() *Enveloup {
	return &Enveloup{
		events:   make(chan types.Event),
		subs:     make(map[string]Sub),
		mu:       &sync.Mutex{},
		controls: false,
	}
}
func (e *Enveloup) ToggleControls() {
	e.controls = !e.controls
}

func (e *Enveloup) SetFocused(focused string) {
	e.currFocused = focused
}

func (e *Enveloup) Subscribe(events []types.Event, state *state.State) {
	e.mu.Lock()
	e.subs[state.Key] = Sub{
		Events: events,
		Chan:   state.Events,
	}
	e.mu.Unlock()
	go func() {
		for {
			curr, ok := e.subs[e.currFocused]
			v, open := <-e.events
			if v.Key == keyboard.KeyEsc {
				e.ToggleControls()
				continue
			}
			if e.controls {
				c, ok := e.subs["controller"]
				if ok {
					c.Chan <- v
					continue
				}
			}
			if ok && open {
				for _, event := range curr.Events {
					if v == event {
						curr.Chan <- v
					}
				}
			}
		}
	}()
}

func (e *Enveloup) Run() {
	go func() {
		err := keyboard.Open()
		if err != nil {
			log.Fatal(err)
		}
		for {
			r, key, err := keyboard.GetKey()
			if err != nil {
				log.Print("err: ", err)
				continue
			}
			go func() {
				_, ok := e.subs[e.currFocused]
				if ok {
					e.events <- types.Event{Key: key, Char: string(r)}
				}
			}()
		}
	}()
}