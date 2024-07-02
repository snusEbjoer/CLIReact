package screen

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/snusEbjoer/cli-react/react"
	"github.com/snusEbjoer/cli-react/reactlib"
)

type Renderer interface {
	Render() string
	GetState() *reactlib.State
}

type Screen struct {
	Views        []Renderer
	React        *react.React
	Controller   *reactlib.Controller
	State        *reactlib.State
	CustomStates []*reactlib.State
}

func New(views []Renderer, react *react.React) *Screen {
	s := &Screen{
		views,
		react,
		react.GetController(),
		react.UseState(0, "controller"),
		[]*reactlib.State{},
	}
	return s
}

func (s *Screen) Show() string {
	var focusedStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Height(1).Width(7).Foreground(lipgloss.Color("#7D56F4"))
	var defaultStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Height(1).Width(7)

	view := []string{}
	for i, v := range s.Views {
		if i == s.Controller.CurrFocusedState.Curr.(int) {
			view = append(view, focusedStyle.Render(v.Render()))
		} else {
			view = append(view, defaultStyle.Render(v.Render()))
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Center, view...)
}
func (s *Screen) GetStates() []*reactlib.State {
	res := make([]*reactlib.State, 0, len(s.Views))
	for _, st := range s.Views {
		res = append(res, st.GetState())
	}
	return res
}

func (s *Screen) Render() {
	states := []*reactlib.State{}
	states = append(states, s.GetStates()...)
	states = append(states, s.State)
	states = append(states, s.Controller.CurrFocusedState)
	states = append(states, s.CustomStates...)

	fmt.Printf("%s", s.Show())
	react.UseEffect(func(a ...any) {
		fmt.Print("\033[H\033[J")
		fmt.Printf("%s", s.Show())
	}, []any{}, states...)
}
