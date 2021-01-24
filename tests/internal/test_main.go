package main

import (
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/shenyisyn/goft-gin/tests/internal/Configuration"
	"github.com/shenyisyn/goft-gin/tests/internal/classes"
	"github.com/shenyisyn/goft-gin/tests/internal/fairing"
)

func main() {

	goft.Ignite().
		Config(Configuration.NewMyConfig()).
		Attach(fairing.NewGlobalFairing()).
		Mount("", classes.NewIndexClass()). //控制器，挂载到v1
		Launch()

}
