package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	. "github.com/shenyisyn/goft-gin/tests/classes"
	"github.com/shenyisyn/goft-gin/tests/fairing"
)

func main() {
	goft.Ignite().
		Attach(fairing.NewGlobalFairing()).
		Mount("v1", NewIndexClass()). //控制器，挂载到v1
		Launch()
}
