package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}
func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	//goft.Throw("aaa",500,ctx)
	fmt.Println(ctx.Get("name"))
	return "index"
}
func (this *IndexClass) Test(ctx *gin.Context) goft.Json {
	fmt.Println("name is", ctx.PostForm("name"))
	return NewDataModel(101, "wfew")
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.GetIndex).
		Handle("POST", "/test", this.Test)
}
func (this *IndexClass) Name() string {
	return "IndexClass"
}
