package demo

import (
	"git.kkcoding.com/kk-golang/kk-lib/micro"
)

type QueryTask struct {
	micro.Task
	Id     int `json:"id"`
	Result TaskResult
}

func (T *QueryTask) GetName() string {
	return "/q"
}

func (T *QueryTask) GetResult() interface{} {
	return &T.Result
}

type CreateTask struct {
	micro.Task
	DemoObj DemoObject `json:"demo"`
	Result  TaskResult
}

func (T *CreateTask) GetName() string {
	return "/c"
}

func (T *CreateTask) GetResult() interface{} {
	return &T.Result
}
