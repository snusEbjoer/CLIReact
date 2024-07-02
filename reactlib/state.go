package reactlib

type State struct {
	Prev           any
	Curr           any
	Key            string
	Events         chan Event
	Changed        bool
	SubbedOnEvents map[Event]struct{}
	handlers       map[Event]Handler
	controller     *Controller
}

type Handler struct {
	fn   func(...any)
	args []any
}

func NewState(value any, key string, controller *Controller) *State {
	s := &State{
		Curr:           value,
		Events:         make(chan Event),
		Key:            key,
		handlers:       make(map[Event]Handler),
		Changed:        false,
		SubbedOnEvents: make(map[Event]struct{}),
		controller:     controller,
	}
	s.controller.Subscribe(Sub{
		Events: s.SubbedOnEvents,
		Chan:   s.Events,
	})
	s.listenEvents()
	return s
}

func (s *State) SetState(value any) {
	s.Prev = s.Curr
	s.Curr = value
	s.Changed = true
}

func (s *State) AddHandler(event Event, fn func(...any), args ...any) {
	s.handlers[event] = Handler{fn, args}
	s.SubbedOnEvents[event] = struct{}{}
}

func (s *State) listenEvents() {
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
