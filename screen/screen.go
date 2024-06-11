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
	Views        []Renderer
	Controls     *enveloup.Enveloup
	State        *state.State
	CustomStates []*state.State
}

func New(views []Renderer, controller *enveloup.Enveloup) *Screen {
	return &Screen{
		views,
		controller,
		state.New(0, "controller"),
		[]*state.State{},
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
	s.Controls.SetFocused(s.Views[0].GetKey())

	s.State.AddHandler(utils.NdaKey(keyboard.KeyArrowRight), func(a ...any) {
		if s.State.Curr == len(s.Views)-1 {
			s.State.SetState(0)
			s.Controls.SetFocused(s.Views[s.State.Curr.(int)].GetKey())
			return
		}
		s.State.SetState(s.State.Curr.(int) + 1)
		s.Controls.SetFocused(s.Views[s.State.Curr.(int)].GetKey())
	}, []any{})

	s.State.AddHandler(utils.NdaKey(keyboard.KeyArrowLeft), func(a ...any) {
		if s.State.Curr == 0 {
			s.State.Curr = len(s.Views) - 1
			s.Controls.SetFocused(s.Views[s.State.Curr.(int)].GetKey())
			return
		}
		s.State.SetState(s.State.Curr.(int) - 1)
		s.Controls.SetFocused(s.Views[s.State.Curr.(int)].GetKey())
	}, []any{})
	s.State.AddHandler(utils.NdaKey(keyboard.KeyEsc), func(a ...any) {
		s.State.SetState(s.State.Curr)
	}, []any{})
	s.Controls.SubscribeToController([]types.Event{utils.NdaKey(keyboard.KeyArrowRight), utils.NdaKey(keyboard.KeyArrowLeft), utils.NdaKey(keyboard.KeyEsc)}, s.State)
}

func (s *Screen) Mode() string {
	idx := 0
	if s.Controls.GetControls() {
		idx = 1
	}
	return [...]string{"Input", "Normal"}[idx]
}
func (s *Screen) Render() {
	states := []*state.State{}
	states = append(states, s.GetStates()...)
	states = append(states, s.State)
	states = append(states, s.CustomStates...)

	fmt.Printf("%s %s %d \r", s.Show(), s.Mode(), s.State.Curr)
	state.UseEffect(func(a ...any) {
		fmt.Printf("%s %s %d \r", s.Show(), s.Mode(), s.State.Curr)
	}, []any{}, states)
}
