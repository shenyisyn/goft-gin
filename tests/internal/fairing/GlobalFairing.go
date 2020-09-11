package fairing

import (
	"github.com/gin-gonic/gin"
)

type GlobalFairing struct{}

func NewGlobalFairing() *GlobalFairing {
	return &GlobalFairing{}
}

func (this *GlobalFairing) OnRequest(ctx *gin.Context) error {
	ctx.Set("name", " global name ")
	return nil
}
func (this *GlobalFairing) OnResponse(ret interface{}) (interface{}, error) {
	if str, ok := ret.(string); ok {
		str = str + "_global"
		return str, nil
	}
	return ret, nil
}
