package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/shenyisyn/goft-gin/tests/Services"
	"github.com/shenyisyn/goft-gin/tests/fairing"
)

type IndexClass struct {
	MyTest *Services.TestService `inject:"-"`
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}
func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	this.MyTest.Naming.ShowName()
	return "IndexClass"
}
func (this *IndexClass) Test(ctx *gin.Context) goft.Json {
	fmt.Println("name is", ctx.PostForm("name"))
	return NewDataModel(101, "wfew")
}
func (this *IndexClass) TestSql(ctx *gin.Context) goft.SimpleQuery {
	return "select * from users"
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.HandleWithFairing("GET", "/",
		this.GetIndex, fairing.NewIndexFairing()).
		Handle("GET", "/sql", this.TestSql).
		Handle("POST", "/test", this.Test)
}
func (this *IndexClass) Name() string {
	return "IndexClass"
}
