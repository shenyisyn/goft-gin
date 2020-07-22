package fairing

import (
	"github.com/gin-gonic/gin"
	"log"
)

type IndexFairing struct{}

func NewIndexFairing() *IndexFairing {
	return &IndexFairing{}
}

func (this *IndexFairing) OnRequest(ctx *gin.Context) error {
	ctx.Set("name", "shenyi")
	log.Println("index fairing")
	return nil
}
func (this *IndexFairing) OnResponse(ret interface{}) (interface{}, error) {
	if str, ok := ret.(string); ok {
		str = "this is " + str
		return str, nil
	}

	return ret, nil
}
