package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"sync"
)

var innerRouter *GoftTree // inner tree node . backup httpmethod and path
var innerRouter_once sync.Once

func getInnerRouter() *GoftTree {
	innerRouter_once.Do(func() {
		innerRouter = NewGoftTree()
	})
	return innerRouter
}

type Goft struct {
	*gin.Engine
	g            *gin.RouterGroup // 保存 group对象
	beanFactory  *BeanFactory
	exprData     map[string]interface{}
	currentGroup string // temp-var for group string
}

func Ignite() *Goft {
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory(),
		exprData: map[string]interface{}{},
	}
	g.Use(ErrorHandler()) //强迫加载的异常处理中间件
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	if config.Server.Html != "" {
		g.LoadHTMLGlob(config.Server.Html)
	}
	return g
}
func (this *Goft) Launch() {
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	this.Run(fmt.Sprintf(":%d", port))
}
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		getInnerRouter().addRoute(httpMethod, relativePath, h) // for future
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}
func (this *Goft) HandleWithFairing(httpMethod, relativePath string, handler interface{}, fairings ...Fairing) *Goft {
	if h := Convert(handler); h != nil {
		g := "/" + this.currentGroup
		if g == "/" {
			g = ""
		}
		getInnerRouter().addRoute(httpMethod, g+relativePath, fairings) //for future
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

// 注册中间件
func (this *Goft) Attach(f ...Fairing) *Goft {
	getFairingHandler().AddFairing(f...)
	return this
}

func (this *Goft) Beans(beans ...Bean) *Goft {
	for _, bean := range beans {
		this.exprData[bean.Name()] = bean
	}
	this.beanFactory.setBean(beans...)
	return this
}

func (this *Goft) Mount(group string, classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		this.currentGroup = group
		class.Build(this)
		this.beanFactory.inject(class)
		this.Beans(class)
	}
	return this
}

//0/3 * * * * *  //增加定时任务
func (this *Goft) Task(cron string, expr interface{}) *Goft {
	var err error
	if f, ok := expr.(func()); ok {
		_, err = getCronTask().AddFunc(cron, f)
	} else if exp, ok := expr.(Expr); ok {
		_, err = getCronTask().AddFunc(cron, func() {
			_, expErr := ExecExpr(exp, this.exprData)
			if expErr != nil {
				log.Println(expErr)
			}
		})
	}

	if err != nil {
		log.Println(err)
	}
	return this
}
