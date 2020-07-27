package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/shenyisyn/goft-gin/tests/fairing"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}
func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	//goft.Throw("aaa",500,ctx)
	//fmt.Println(ctx.Request.URL.Query())
	return "IndexClass"
}
func (this *IndexClass) Test(ctx *gin.Context) goft.Json {
	fmt.Println("name is", ctx.PostForm("name"))
	return NewDataModel(101, "wfew")
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.HandleWithFairing("GET", "/",
		this.GetIndex, fairing.NewIndexFairing()).
		Handle("POST", "/test", this.Test)
}
func (this *IndexClass) Name() string {
	return "IndexClass"
}
