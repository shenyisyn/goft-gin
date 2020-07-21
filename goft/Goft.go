package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Goft struct {
	*gin.Engine                  //我们把 engine放到 主类里
	g           *gin.RouterGroup //这里就是保存 group对象
	beanFactory *BeanFactory
	exprData    map[string]interface{}
}

func Ignite() *Goft { //这就是所谓的构造函数，ignite有 发射、燃烧， 很有激情。符合我们骚动的心情
	g := &Goft{Engine: gin.New(), beanFactory: NewBeanFactory(), exprData: map[string]interface{}{}}
	g.Use(ErrorHandler()) //强迫加载的异常处理中间件
	config := InitConfig()
	g.beanFactory.setBean(config) //整个配置加载进bean中
	if config.Server.Html != "" {

		g.LoadHTMLGlob(config.Server.Html)
	}
	return g
}
func (this *Goft) Launch() { //最终启动函数， 不用run，没有逼格
	var port int32 = 8080
	if config := this.beanFactory.GetBean(new(SysConfig)); config != nil {
		port = config.(*SysConfig).Server.Port
	}
	getCronTask().Start()
	this.Run(fmt.Sprintf(":%d", port))
}
func (this *Goft) Handle(httpMethod, relativePath string, handler interface{}) *Goft {
	if h := Convert(handler); h != nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this // 大功告成
}

//我们弄个方法名叫做Attach ，代表加入中间件。 参数 么。。。。
func (this *Goft) Attach(f Fairing) *Goft {
	this.Use(func(context *gin.Context) {
		err := f.OnRequest(context) //到这一步 看懂没
		if err != nil {
			context.AbortWithStatusJSON(400, gin.H{"error": err.Error()}) // 这个能看懂的啊
		} else {
			context.Next() //继续往下走
		}
	})
	return this
}

//加入Bean
func (this *Goft) Beans(beans ...Bean) *Goft {
	// 取出bean的名称 ，然后 加入到exprData里面
	for _, bean := range beans {
		this.exprData[bean.Name()] = bean
	}
	this.beanFactory.setBean(beans...)
	return this
}

// 加一个group 参数
func (this *Goft) Mount(group string, classes ...IClass) *Goft { // 这是挂载， 后面还需要加功能。
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this) //这一步是关键 。 这样在main里面 就不需要 调用了
		this.beanFactory.inject(class)
		this.Beans(class) //控制器 也作为bean加入到bean容器
	}
	return this
}

//0/3 * * * * *  //增加定时任务
func (this *Goft) Task(cron string, expr interface{}) *Goft {
	var err error
	if f, ok := expr.(func()); ok {
		_, err = getCronTask().AddFunc(cron, f)
	} else if exp, ok := expr.(Expr); ok { //代表 是表达式
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
