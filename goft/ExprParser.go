package goft

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

const (
	VarPattern       = `[0-9a-zA-Z_\.\(\)]+`
	CompareSign      = ">|>=|<=|<|==|!="
	CompareSignToken = "gt|ge|le|lt|eq|ne"
	ComparePattern   = `^(` + VarPattern + `)\s*(` + CompareSign + `)\s*(` + VarPattern + `)\s*$`
)

type Expr string //表达式类型
//可比较表达式 解析类， 譬如a>3   b!=4 a!=n    a>3  [gt .a  3]
type ComparableExpr string

func (this ComparableExpr) filter() string {
	reg, err := regexp.Compile(ComparePattern)
	if err != nil {
		return ""
	}
	ret := reg.FindStringSubmatch(string(this))
	if ret != nil && len(ret) == 4 {
		token := getCompareToken(ret[2])
		if token == "" {
			return ""
		}
		return fmt.Sprintf("%s %s %s", token, parseToken(ret[1]), parseToken(ret[3]))
	}
	return ""
}

//普通表达式，如 .user.Age  .user.Info(101)
type SimpleExpr string

func (this SimpleExpr) filter() string {
	// 处理括号里面的参数
	return ""
}

//根据比较符 ，获取token
func getCompareToken(sign string) string {
	for index, item := range strings.Split(CompareSign, "|") {
		if item == sign {
			return strings.Split(CompareSignToken, "|")[index]
		}
	}
	return ""
}

//对于数字不加.(点)
func parseToken(token string) string {
	if IsNumeric(token) {
		return token
	} else {
		return "." + token
	}
}

//是否是"比较表达式"
func IsComparableExpr(expr string) bool {
	reg, err := regexp.Compile(ComparePattern)
	if err != nil {
		return false
	}
	return reg.MatchString(expr)
}

//执行表达式，临时方法后期需要修改
func ExecExpr(expr Expr, data map[string]interface{}) (string, error) {
	tpl := template.New("expr").Funcs(map[string]interface{}{
		"echo": func(params ...interface{}) interface{} {
			return fmt.Sprintf("echo:%v", params[0])
		},
	})
	t, err := tpl.Parse(fmt.Sprintf("{{%s}}", expr))
	if err != nil {
		return "", err
	}
	var buf = &bytes.Buffer{}
	err = t.Execute(buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
