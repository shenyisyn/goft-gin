package goft

import (
	"github.com/robfig/cron/v3"
	"sync"
)

//goft-task

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor //任务列表
var once sync.Once
var onceCron sync.Once
var taskCron *cron.Cron //定时任务
func init() {
	chlist := getTaskList() //得到任务列表
	go func() {
		for t := range chlist {
			doTask(t)
		}
	}()
}
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor) //初始化
	})
	return taskList
}

type TaskExecutor struct {
	f        TaskFunc
	p        []interface{} //参数
	callback func()
}

func NewTaskExecutor(f TaskFunc, p []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, p: p, callback: callback}
}
func (this *TaskExecutor) Exec() { //执行任务
	this.f(this.p...)
}
func Task(f TaskFunc, cb func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, cb) //增加任务队列
	}()
}
