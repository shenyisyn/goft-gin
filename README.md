# Goft-Gin
* 在web框架gin的基础上做的脚手架

## 安装
go get -u github.com/shenyisyn/goft-gin@v0.4.1

## 功能说明
 控制器、简易依赖注入、中间件、表达式、任务组件等。
 
 后续功能正在发布中,因此**可能有较大改动**
## 使用视频(不定期更新)
### 第一章：控制器
* [第1讲:控制器的使用：返回String和JSON](http://www.jtthink.com/course/play/2784)
直接开门见山。先讲下控制器的使用

* [第2讲:中间件的使用(1)：判断必要参数](http://www.jtthink.com/course/play/2785)
今天演示下中间件的使用方式。在执行控制器方法前可以xxoo

* [第3讲:中间件的使用(2)：修改响应内容](http://www.jtthink.com/course/play/2786)
当执行完成控制器方法后进行响应值的修改

* [第4讲:路由级的中间件(1):基本使用](http://www.jtthink.com/course/play/2787)
原生gin的中间件无法定位到具体的URL。改造后目前支持路由级的中间件，支持绑定具体的URL进行中间件执行

* [第5讲:路由级的中间件(2):参数验证和业务分离（上）](http://www.jtthink.com/course/play/2797)
今天顺便做个例子，请求控制的业务代码常规来讲应该怎么写

* [第6讲:路由级的中间件(2):参数验证和业务分离（下）](http://www.jtthink.com/course/play/2798)
今天我们把业务代码和参数验证进行分离，原理也是使用路由级中间件来完成

----------------------------你们要的ORM来了
* [第7讲:依赖注入和ORM 使用 (Gorm)](http://www.jtthink.com/course/play/2799)
重写了IoC。实现控制器注入，并演示ORM注入的方式

* [第8讲:ORM执行简化:直接返回SQL语句(GORM)](http://www.jtthink.com/course/play/2829)
目前支持在控制器中直接返回SQL，即可自动JSON输出

* [第9讲:ORM执行简化:控制器直接返回SQL语句(XORM)](http://www.jtthink.com/course/play/2830)
上节课我们注入GORM。假设你不想使用GORM。那么今天提供XORM的适配器写法

* [第10讲:控制器返回SQL语句：支持参数](http://www.jtthink.com/course/play/2843)
我们在控制器只返回SQL即可输出JSON，今天支持参数，防止注入风险

* [第11讲:控制器返回SQL语句：支持自定义JSON字段](http://www.jtthink.com/course/play/2844)
沿着上节课。今天我们支持自定义JSON字段的输出

* [第12讲:DAO层示例：用户DAO的写法](https://www.jtthink.com/course/play/2846)
今天演示下 我们脚手架下DAO层的写法

* [第13讲:Service层示例：用户Service层的基本写法](https://www.jtthink.com/course/play/2871)
有了上节课基础，我们进而演示下Service层的写法

* [第14讲:Service层示例：用户登录示例](https://www.jtthink.com/course/play/2872)
趁热打铁，再来个用户登录示例

* [第15讲:ORM简化：自定义输出key、Query执行](https://www.jtthink.com/course/play/2873)
应网友要求做了一些功能的支持，可以在控制器中直接获取Query结果 

### 第二章：简化和修改版的领域驱动 （本章带点教学的味道。）
* [第16讲:超简领域驱动模型入门：基本分层](https://www.jtthink.com/course/play/2905)
DDD很火，我们的脚手架怎能少了它。由于Go语言的一些特征，我们做了很大简化。本章做个精简版说明，然后撸代码

* [第17讲:领域层:用户实体编写和值对象(初步)](https://www.jtthink.com/course/play/2906)
我们先从领域层开始，以用户登录注册和日志为例。展开领域层的实体编写

* [第18讲:领域层:用户实体和值对象（2）--构造函数](https://www.jtthink.com/course/play/2907)
承接上节课，我们完成领域层实体构造函数的编写

* [第19讲:领域层:实体接口、聚合的概念](https://www.jtthink.com/course/play/2908)
补充下上节课，我们把实体加入接口。由于Go没有继承，因此今天只是演示种写法。并初步认识聚合

* [第20讲:领域层:初步划分聚合（用户为例）](https://www.jtthink.com/course/play/2926)
沿着上节课，我们以用户为例，简单划分下用户聚合

* [第21讲:领域层:仓储层(Repository)、基础设施层](https://www.jtthink.com/course/play/2927)
在我们上节课的基础上，扩展出仓储层。并且初步接触下基础设施层对仓储层的作用

* [第22讲:领域层:聚合方法示例(用户为例)](https://www.jtthink.com/course/play/2928)
在上节课的基础上，我们做个代码示例。其中做法也做了一定修改和简化，使之更适合我们的项目需求

* [第23讲:领域层之:领域服务层的基本使用](https://www.jtthink.com/course/play/2929)
领域层基本构建完毕，今天补充下服务层的基本用法

* [第24讲:应用层入门(Application):DTO数据传输对象](https://www.jtthink.com/course/play/2961)
本课时进入应用层的讲解，先说下DTO的基本作用

* [第25讲:应用层入门：DTO和实体的映射](https://www.jtthink.com/course/play/2962)
上节课建立了DTO对象，今天演示下和实体之间的映射

* [第26讲:应用层：应用服务层的基本用法、超简案例演示](https://www.jtthink.com/course/play/2963)
应用服务层也是很重要的一层，它是领域层和展现出的枢纽。今天写个超简案例做下演示

* [第27讲:进入interface层:脚手架开始发挥作用](https://www.jtthink.com/course/play/2964)
今天这课时我们终于进入实际的功能开发和展现，我们的脚手架终于要开始发挥作用了

* [第28讲:interface层:异常的处理的方法](https://www.jtthink.com/course/play/2982)
我们尽可能的不要在interface层出现过多的if else判断。尤其是类似error的处理要做封装，今天演示套路

* [第29讲:引入GORM、仓储层取值](https://www.jtthink.com/course/play/2983)
今天我们初步把基础层、展现层、领域层和应用层连接在了一起
## License
© jtthink, 2020~time.Now
Released under the [Apache License 2.0](https://github.com/shenyisyn/goft-gin/blob/master/LICENSE)
