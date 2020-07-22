package goft

import "github.com/gin-gonic/gin"

type Fairing interface {
	OnRequest(*gin.Context) error
	OnResponse(result interface{}) (interface{}, error)
}
