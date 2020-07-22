package goft

import (
	"github.com/gin-gonic/gin"
)

const (
	HTTP_STATUS = "GOFT_STATUS"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				status := 400 //default status==400
				if value, exists := context.Get(HTTP_STATUS); exists {
					if v, ok := value.(int); ok {
						status = v
					}
				}
				context.AbortWithStatusJSON(status, gin.H{"error": e})
			}
		}()
		context.Next()
	}
}
func Throw(err string, code int, context *gin.Context) {
	context.Set(HTTP_STATUS, code)
	panic(err)
}
func Error(err error, msg ...string) {
	if err == nil {
		return
	} else {
		errMsg := err.Error()
		if len(msg) > 0 {
			errMsg = msg[0]
		}
		panic(errMsg)
	}
}
