package reactlib

import (
	"log"
	"sync"

	"github.com/eiannone/keyboard"
)

type Enveloup struct {
	mu             *sync.Mutex
	controllerChan chan Event
}

func NewEnvelope(controllerChan chan Event) *Enveloup {
	e := &Enveloup{
		mu:             &sync.Mutex{},
		controllerChan: controllerChan,
	}
	err := keyboard.Open()
	if err != nil {
		log.Fatal("Can't open keyboard:", err)
	}
	e.run()
	return e
}

func (e *Enveloup) run() {
	go func() {
		for {
			r, key, err := keyboard.GetKey()
			if err != nil {
				log.Fatal("err: ", err)
			}
			currEvent := Event{Key: key, Char: string(r)}
			go func() {
				e.controllerChan <- currEvent
			}()
		}
	}()
}
