package main

import (
	"github.com/snusEbjoer/cli-react/component"
	"github.com/snusEbjoer/cli-react/input"
	"github.com/snusEbjoer/cli-react/react"
	"github.com/snusEbjoer/cli-react/screen"
)

func main() {
	r := react.New()
	c := component.New(0, "count", r)
	_ = input.New("", "input", r)
	screen.New([]screen.Renderer{c}, r).Render()
	select {}
}
