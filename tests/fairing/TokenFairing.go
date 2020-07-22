package fairing

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TokenFairing struct{}

func NewTokenFairing() *TokenFairing {
	return &TokenFairing{}
}
func (this *TokenFairing) OnRequest(ctx *gin.Context) error {
	if ctx.Query("token") == "" {
		return fmt.Errorf("token required")
	}
	return nil
}
func (this *TokenFairing) OnResponse(ret interface{}) (interface{}, error) {
	return ret, nil
}
