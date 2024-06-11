package main

import (
	"github.com/snusEbjoer/cli-react/component"
	"github.com/snusEbjoer/cli-react/enveloup"
	"github.com/snusEbjoer/cli-react/screen"
)

func main() {
	e := enveloup.New() // SINGLETON
	c := component.New(0, "count", e)
	c2 := component.New(0, "count2", e)
	c.Init()
	c2.Init()
	s := screen.New([]screen.Renderer{c, c2}, e)
	s.ScreenContoller()
	e.Run()
	s.Render()
	select {}
}
