package goft

type IClass interface {
	Build(goft *Goft) //参数和方法名必须一致
	Name() string
}
