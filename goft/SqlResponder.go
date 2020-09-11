package goft

import "log"

type Query interface {
	Sql() string
	Args() []interface{}
	Mapping() map[string]string
	First() bool
	Key() string
	Get() interface{}
}
type SimpleQueryWithArgs struct {
	sql        string
	args       []interface{}
	mapping    map[string]string
	fetchFirst bool
	datakey    string //data:{datakey:xxxx}   add by shenyi
}

func NewSimpleQueryWithArgs(sql string, args []interface{}) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, args: args}
}
func NewSimpleQueryWithMapping(sql string, mapping map[string]string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, mapping: mapping}
}
func NewSimpleQueryWithFetchFirst(sql string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, fetchFirst: true}
}
func NewSimpleQueryWithKey(sql string, key string) *SimpleQueryWithArgs {
	return &SimpleQueryWithArgs{sql: sql, datakey: key}
}
func (this *SimpleQueryWithArgs) Sql() string {
	return this.sql
}
func (this *SimpleQueryWithArgs) Mapping() map[string]string {
	return this.mapping
}
func (this *SimpleQueryWithArgs) Args() []interface{} {
	return this.args
}
func (this *SimpleQueryWithArgs) First() bool {
	return this.fetchFirst
}
func (this *SimpleQueryWithArgs) Key() string {
	return this.datakey
}
func (this *SimpleQueryWithArgs) WithMapping(mapping map[string]string) *SimpleQueryWithArgs {
	this.mapping = mapping
	return this
}
func (this *SimpleQueryWithArgs) WithFirst() *SimpleQueryWithArgs {
	this.fetchFirst = true
	return this
}
func (this *SimpleQueryWithArgs) WithKey(key string) *SimpleQueryWithArgs {
	this.datakey = key
	return this
}
func (this *SimpleQueryWithArgs) Get() interface{} {
	ret, err := queryForMapsByInterface(this)
	if err != nil {
		log.Println("query get error:", err)
		return nil
	}
	return ret
}

type SimpleQuery string

func (this SimpleQuery) WithArgs(args ...interface{}) *SimpleQueryWithArgs {
	return NewSimpleQueryWithArgs(string(this), args)
}
func (this SimpleQuery) WithMapping(mapping map[string]string) *SimpleQueryWithArgs {
	return NewSimpleQueryWithMapping(string(this), mapping)
}
func (this SimpleQuery) WithFirst() *SimpleQueryWithArgs {
	return NewSimpleQueryWithFetchFirst(string(this))
}
func (this SimpleQuery) WithKey(key string) *SimpleQueryWithArgs {
	return NewSimpleQueryWithKey(string(this), key)
}
func (this SimpleQuery) First() bool {
	return false
}
func (this SimpleQuery) Sql() string {
	return string(this)
}
func (this SimpleQuery) Key() string {
	return ""
}
func (this SimpleQuery) Args() []interface{} {
	return []interface{}{}
}
func (this SimpleQuery) Mapping() map[string]string {
	return map[string]string{}
}
func (this SimpleQuery) Get() interface{} {
	return NewSimpleQueryWithArgs(string(this), nil).Get()
}
