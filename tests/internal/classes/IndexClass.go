package classes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shenyisyn/goft-gin/goft"
	"github.com/shenyisyn/goft-gin/tests/internal/Services"
	"github.com/shenyisyn/goft-gin/tests/internal/fairing"
)

type IndexClass struct {
	MyTest  *Services.TestService `inject:"-"`
	MyTest2 *Services.TestService
	Age     *goft.Value `prefix:"user.age"`
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}
func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	this.MyTest.Naming.ShowName()
	return "IndexClass"
}
func (this *IndexClass) Test(ctx *gin.Context) goft.Json {
	//fmt.Println("name is", ctx.PostForm("name"))
	fmt.Println(this.Age)
	return NewDataModel(101, "wfew")
}
func (this *IndexClass) TestUsers(ctx *gin.Context) goft.Query {
	this.MyTest2.Naming.ShowName()
	return goft.SimpleQuery("select * from users").WithMapping(map[string]string{
		"user_name": "uname",
	}).WithKey("result")
}
func (this *IndexClass) TestUserDetail(ctx *gin.Context) goft.Json {
	ret := goft.SimpleQuery("select * from users where user_id=?").
		WithArgs(ctx.Param("id")).WithMapping(map[string]string{
		"usr": "user",
	}).WithFirst().WithKey("result").Get()

	fmt.Printf("%T", ret.(gin.H)["result"].(map[string]interface{}))
	return ret
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.HandleWithFairing("GET", "/",
		this.GetIndex, fairing.NewIndexFairing()).
		Handle("GET", "/users", this.TestUsers).
		Handle("GET", "/users/:id", this.TestUserDetail).
		Handle("POST", "/test", this.Test)
}
func (this *IndexClass) Name() string {
	return "IndexClass"
}
