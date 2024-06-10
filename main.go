package main

import (
	"github.com/snusEbjoer/cli-react/component"
	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/screen"
)

func main() {
	e := enveloup.New()
	c := component.New(0, "count", e)
	c.Init()
	s := screen.New([]screen.Renderer{c}, e)
	s.ScreenContoller()
	s.Render()
	e.Run()
	select {}
}
