package enveloup

import (
	"log"
	"sync"

	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/pkg/types"
	"github.com/snusEbjoer/cli-react/state"
)

type Sub struct {
	Events []types.Event
	Chan   chan types.Event
}

type Enveloup struct {
	events       chan types.Event
	subs         map[string]Sub
	controls     bool
	controlsChan chan types.Event
	currFocused  string
	mu           *sync.Mutex
}

func New() *Enveloup {
	return &Enveloup{
		events:       make(chan types.Event),
		subs:         make(map[string]Sub),
		controlsChan: make(chan types.Event),
		mu:           &sync.Mutex{},
		controls:     false,
	}
}
func (e *Enveloup) ToggleControls() {
	e.controls = !e.controls
}

func (e *Enveloup) SetFocused(focused string) {
	e.currFocused = focused
}
func (e *Enveloup) GetControls() bool {
	return e.controls
}

func (e *Enveloup) Subscribe(events []types.Event, state *state.State) {
	e.mu.Lock()
	e.subs[state.Key] = Sub{
		Events: events,
		Chan:   state.Events,
	}
	e.mu.Unlock()
}

func (e *Enveloup) SubscribeToController(events []types.Event, state *state.State) { // rewrite as Subcribe()
	go func() {
		for {
			v, ok := <-e.controlsChan
			if ok {
				for _, event := range events {
					if v == event {
						state.Events <- v
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
			if key == keyboard.KeyEsc { // fix later
				e.ToggleControls()
				go func() {
					e.controlsChan <- types.Event{Key: key, Char: string(r)}
				}()
				continue
			}
			currEvent := types.Event{Key: key, Char: string(r)}
			if e.controls {
				go func() {
					e.controlsChan <- currEvent
				}()
			} else {
				go func() {
					v, ok := e.subs[e.currFocused]
					if ok {
						for _, event := range v.Events {
							if event == currEvent {
								v.Chan <- types.Event{Key: key, Char: string(r)}
							}
						}
					}
				}()
			}
		}
	}()
}
