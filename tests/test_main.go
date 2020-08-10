package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/shenyisyn/goft-gin/tests/Configuration"
	. "github.com/shenyisyn/goft-gin/tests/classes"
	"github.com/shenyisyn/goft-gin/tests/fairing"
)

func main() {
	goft.Ignite().
		Config(Configuration.NewMyConfig()).
		Attach(fairing.NewGlobalFairing()).
		Mount("", NewIndexClass()). //控制器，挂载到v1
		Launch()
}
