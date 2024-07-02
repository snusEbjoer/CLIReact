package react

import "github.com/snusEbjoer/cli-react/reactlib"

type React struct {
	enveloup   *reactlib.Enveloup
	controller *reactlib.Controller
}

func New() *React { // ёбаный пиздец убейте меня...
	c := reactlib.NewController()
	r := &React{
		enveloup:   reactlib.NewEnvelope(c.EventsChan),
		controller: c,
	}
	c.CurrFocusedState = r.UseState(0, "controller")
	r.controller = c
	r.controller.Init()
	return r
}

func (r *React) GetController() *reactlib.Controller {
	return r.controller
}
func (r *React) UseState(value any, key string) *reactlib.State {
	return reactlib.NewState(value, key, r.controller)
}

func UseEffect(fn func(...any), args []any, dep ...*reactlib.State) {
	go func() {
		for {
			for _, s := range dep {
				if s.Changed {
					fn(args...)
					s.Changed = false
				}
			}
		}
	}()
}
