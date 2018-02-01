package demo

import (
	"log"

	"git.kkcoding.com/kk-golang/kk-lib/db"
	"git.kkcoding.com/kk-golang/kk-lib/micro"
)

/*
	定义一个service类型，继承micro.Service
	service中必须包含task的定义，否则无法处理请求。
	service中定义以“Handle”+task类型名称命名的方法，用于处理对应的task请求。
	service必须完全实现IService接口，否则DemoService1会作为基类service进行处理，无法找到HandleTask方法。
*/
type DemoService1 struct {
	micro.Service

	QueryTask *QueryTask
}

//func (s *DemoService1) Handle(a micro.IApp, task micro.ITask) error {
//	return micro.ServiceReflectHandle(a, task, s)
//}

//func (s *DemoService1) GetTitle() string {
//	return "Demo服务1"
//}

/*
	micro中的数据库操作方式：
	micro中提供了Database类型，用于描述数据库的配置信息。Database一般在app上定义，并在配置文件中进行配置，LoadPath方法会加载配置文件内容。
	Database.Open方法负责获取真正的数据库对象sql.Db。此处虽然方法名为Open，但是实际上并不会有建立连接的操作。
	Database会保持一个sql.Db，整个服务中使用同一个sql.Db。
	Database.Open方法首次调用时会维护数据库，具体工作是尝试创建__goscheme表，创建或修改具体的IObject存储表。
	git.kkcoding.com/kk-golang/kk-lib/db包中实现了ORM封装。
*/

func (s *DemoService1) HandleQueryTask(app *DemoApp, task *QueryTask) {
	log.Println("service1 handle q", task)
	var conn, err = app.DB.Open(app)
	if err != nil {
		task.Result = TaskResult{0, err.Error(), nil}
		return
	}
	var rows, err1 = db.Query(conn, app.DemoObject, app.DB.Prefix, "where id=?", task.Id)
	if err1 != nil {
		task.Result = TaskResult{0, err1.Error(), nil}
		return
	}
	var demo DemoObject
	if rows.Next() {
		demo = DemoObject{}
		err1 = db.NewScaner(&demo).Scan(rows)
		if err1 != nil {
			task.Result = TaskResult{0, err1.Error(), nil}
			return
		}
	}
	task.Result = TaskResult{1, "", demo}
}

/*
	可以在一个app中注册多个service。
	多个service可以包含相同的task。此时多个service都响应该task，处理结果以最后一次响应结果为准。
	本例中，DemoService1中处理QueryTask，DemoService2中处理CreateTask
*/
type DemoService2 struct {
	micro.Service

	CreateTask *CreateTask
}

/*B(ReflectHandle)*/
func (s *DemoService2) Handle(a micro.IApp, task micro.ITask) error {
	return micro.ServiceReflectHandle(a, task, s)
}

/*E(ReflectHandle)*/

/*B(Title)*/
func (s *DemoService2) GetTitle() string {
	return "Demo服务2"
}

func (s *DemoService2) HandleCreateTask(app *DemoApp, task *CreateTask) {
	log.Println("service2 handle q", task)
	if task.DemoObj.Name == "" {
		task.Result = TaskResult{0, "name为空", nil}
		return
	} else {
		var conn, err = app.DB.Open(app)
		if err != nil {
			task.Result = TaskResult{0, err.Error(), nil}
			return
		}
		var result, err1 = db.Insert(conn, &task.DemoObj, app.DB.Prefix)
		if err1 != nil {
			task.Result = TaskResult{0, err1.Error(), nil}
			return
		}
		var id, err2 = result.LastInsertId()
		if err2 != nil {
			task.Result = TaskResult{0, err2.Error(), nil}
			return
		}
		task.Result = TaskResult{1, "插入成功", id}
	}
}

type TaskResult struct {
	Status  int         `json:"status"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}
