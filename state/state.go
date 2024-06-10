package state

import (
	"sync"

	"github.com/snusEbjoer/cli-react/pkg/types"
)

type State struct {
	Prev     any
	Curr     any
	Key      string
	Events   chan types.Event
	changed  bool
	handlers map[types.Event]Handler
	mu       sync.RWMutex
}

type Handler struct {
	fn   func(...any)
	args []any
}

func New(value any, key string) *State {
	s := &State{
		Curr:     value,
		Events:   make(chan types.Event),
		Key:      key,
		handlers: make(map[types.Event]Handler),
		changed:  false,
	}
	s.ListenEvents()
	return s
}

func UseEffect(fn func(...any), args []any, dep []*State) {
	go func() {
		for {
			for _, s := range dep {
				if s.changed {
					fn(args...)
					s.changed = false
				}
			}
		}
	}()
}

func (s *State) SetState(value any) {
	s.Prev = s.Curr
	s.Curr = value
	s.changed = true
}

func (s *State) AddHandler(event types.Event, fn func(...any), args []any) {
	s.mu.Lock()
	s.handlers[event] = Handler{fn, args}
	s.mu.Unlock()
}

func (s *State) ListenEvents() {
	go func() {
		for {
			v, open := <-s.Events
			if open {
				v, ok := s.handlers[v]
				if ok {
					v.fn(v.args...)
				}
			}
		}
	}()

}
