package screen

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/pkg/types"
	utils "github.com/snusEbjoer/cli-react/pkg/utils"
	"github.com/snusEbjoer/cli-react/state"
)

type Renderer interface {
	Render() string
	GetKey() string
	GetState() *state.State
}

type Screen struct {
	Views    []Renderer
	curr     int
	Controls *enveloup.Enveloup
}

func New(views []Renderer, controller *enveloup.Enveloup) *Screen {
	return &Screen{
		views,
		0, // fix later
		controller,
	}
}
func (s *Screen) Show() string {
	view := ""
	for _, v := range s.Views {
		view += v.Render()
	}
	return view
}
func (s *Screen) GetStates() []*state.State {
	res := make([]*state.State, 0, len(s.Views))
	for _, st := range s.Views {
		res = append(res, st.GetState())
	}
	return res
}
func (s *Screen) ScreenContoller() {
	controllerState := state.New(0, "controller")
	s.Controls.SetFocused(s.Views[0].GetKey())
	controllerState.AddHandler(utils.NdaKey(keyboard.KeyArrowRight), func(a ...any) {
		if s.curr == len(s.Views)-1 {
			s.curr = 0
			s.Controls.SetFocused(s.Views[s.curr].GetKey())
			return
		}
		s.curr++
		s.Controls.SetFocused(s.Views[s.curr].GetKey())
	}, []any{})
	controllerState.AddHandler(utils.NdaKey(keyboard.KeyArrowLeft), func(a ...any) {
		if s.curr == 0 {
			s.curr = len(s.Views) - 1
			s.Controls.SetFocused(s.Views[s.curr].GetKey())
			return
		}
		s.curr--
		s.Controls.SetFocused(s.Views[s.curr].GetKey())
	}, []any{})
	s.Controls.Subscribe([]types.Event{utils.NdaKey(keyboard.KeyArrowRight), utils.NdaKey(keyboard.KeyArrowLeft)}, controllerState)
}

func (s *Screen) Render() {
	fmt.Printf(s.Show() + "\r")
	state.UseEffect(func(a ...any) {
		fmt.Printf(s.Show() + "\r")
	}, []any{}, s.GetStates())
}
