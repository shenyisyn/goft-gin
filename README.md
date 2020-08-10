# Goft-Gin
* 在web框架gin的基础上做的脚手架

## 安装
go get -u github.com/shenyisyn/goft-gin@v0.3.2

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
## License
© jtthink, 2020~time.Now
Released under the [Apache License 2.0](https://github.com/shenyisyn/goft-gin/blob/master/LICENSE)
