package Configuration

import "github.com/shenyisyn/goft-gin/tests/Services"

type MyConfig struct {
}

func NewMyConfig() *MyConfig {
	return &MyConfig{}
}
func (this *MyConfig) Test() *Services.TestService {
	return Services.NewTestService("mytest")
}
func (this *MyConfig) Naming() *Services.NameService {
	return Services.NewNameService("shenyi")
}
