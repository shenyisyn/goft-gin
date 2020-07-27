package fairing

import (
	"github.com/gin-gonic/gin"
)

type TestFairing struct{}

func NewTestFairing() *TestFairing {
	return &TestFairing{}
}

func (this *TestFairing) OnRequest(ctx *gin.Context) error {
	if name, exists := ctx.Get("name"); exists {
		name = "this is " + name.(string)
		ctx.Set("name", name)
	}
	return nil
}
func (this *TestFairing) OnResponse(ret interface{}) (interface{}, error) {
	if str, ok := ret.(string); ok {
		str = "test_" + str
		return str, nil
	}
	return ret, nil
}
