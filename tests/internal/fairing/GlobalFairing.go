package fairing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type GlobalFairing struct {
	DB *gorm.DB `inject:"-"`
}

func NewGlobalFairing() *GlobalFairing {
	return &GlobalFairing{}
}

func (this *GlobalFairing) OnRequest(ctx *gin.Context) error {
	ctx.Set("name", " global name ")
	return nil
}
func (this *GlobalFairing) OnResponse(ret interface{}) (interface{}, error) {
	fmt.Println(this.DB)
	if str, ok := ret.(string); ok {
		str = str + "_global"
		return str, nil
	}
	return ret, nil
}
