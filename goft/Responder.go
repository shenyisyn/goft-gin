package goft

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

var responderList []Responder
var once_resp_list sync.Once

func get_responder_list() []Responder {
	once_resp_list.Do(func() {
		responderList = []Responder{(StringResponder)(nil),
			(JsonResponder)(nil),
			(ViewResponder)(nil),
			(SqlResponder)(nil),
			(SqlQueryResponder)(nil),
			(VoidResponder)(nil),
		}
	})
	return responderList
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	for _, r := range get_responder_list() {
		r_ref := reflect.TypeOf(r)
		if h_ref.Type().ConvertibleTo(r_ref) {
			return h_ref.Convert(r_ref).Interface().(Responder).RespondTo()
		}
	}
	return nil
}

type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(200, getFairingHandler().handlerFairing(this, context).(string))
	}
}

type Json interface{}
type JsonResponder func(*gin.Context) Json

func (this JsonResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, getFairingHandler().handlerFairing(this, context))
	}
}

type SqlQueryResponder func(*gin.Context) Query

func (this SqlQueryResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getQuery := getFairingHandler().handlerFairing(this, context).(Query)
		ret, err := queryForMapsByInterface(getQuery)
		if err != nil {
			panic(err)
		}
		context.JSON(200, ret)
	}
}

type SqlResponder func(*gin.Context) SimpleQuery

func (this SqlResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getSql := getFairingHandler().handlerFairing(this, context).(SimpleQuery)
		ret, err := queryForMaps(string(getSql), nil, []interface{}{}...)
		if err != nil {
			panic(err)
		}
		context.JSON(200, ret)
	}
}

type Void struct{}
type VoidResponder func(ctx *gin.Context) Void

func (this VoidResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		getFairingHandler().handlerFairing(this, context)
	}
}

// Deprecated: 暂时不提供View的解析
type View string
type ViewResponder func(*gin.Context) View

func (this ViewResponder) RespondTo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.HTML(200, string(this(context))+".html", context.Keys)
	}
}
