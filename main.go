package main

import (
	"github.com/snusEbjoer/cli-react/component"
	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/input"
	"github.com/snusEbjoer/cli-react/screen"
)

func main() {
	e := enveloup.New() // SINGLETON
	c := component.New(0, "count", e)
	c2 := component.New(0, "count2", e)
	input := input.New("", "input", e)
	s := screen.New([]screen.Renderer{c, c2, input}, e)
	s.ScreenContoller()
	e.Run()
	s.Render()
	select {}
}
