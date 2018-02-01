package main

import (
	"Hello-Go/framework/kk/db/demo"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"git.kkcoding.com/kk-golang/kk-lib/db"
)

const (
	tablePrefix string = "dbdemo_"
)

func main() {
	log.Println("Hello Db")
	dataBase, err := sql.Open("mysql", "root:@tcp(127.0.0.1:5566)/study?charset=utf8&sql_mode=ANSI")
	if err != nil {
		log.Panic(err)
	}

	/*
		db提供了ORM封装、数据库表维护
		Init初始化数据库方案
	*/
	err = db.Init(dataBase)
	if err != nil {
		log.Panic(err)
	}

	objProduct := &demo.Product{Name: "Pro1", Price: 1.50}
	/*
		Add负责维护数据对象的存储表，有则修改，无则新增
	*/
	err = db.Add(dataBase, objProduct, tablePrefix, 1)
	if err != nil {
		log.Panic(err)
	}

	/*
		Insert插入对象
	*/
	if ret, err := db.Insert(dataBase, objProduct, tablePrefix); err != nil {
		log.Panic(err)
	} else {
		objProduct.Id, _ = ret.LastInsertId()
		log.Println("insert", *objProduct)
	}

	objTransaction := demo.NewTransaction(objProduct, 3)

	err = db.Add(dataBase, objTransaction, tablePrefix, 1)
	if err != nil {
		log.Panic(err)
	}

	if ret, err := db.Insert(dataBase, objTransaction, tablePrefix); err != nil {
		log.Panic(err)
	} else {
		objTransaction.Id, _ = ret.LastInsertId()
	}

	/*
		Update修改对象，以对象id为依据修改其它属性
	*/
	objProduct.Price += 1.0
	if _, err := db.Update(dataBase, objProduct, tablePrefix); err != nil {
		log.Panic(err)
	} else {
		log.Println("update to", *objProduct)
	}

	/*
		UpdateWithKeys修改对象，以对象id为依据，修改map中指定为true的属性
		map的键为要修改的字段名称，必须是小写，否则无法修改
		所有字段都不修改时，会造成sql异常
		这里内存中修改了Name，但是并没有提交数据库修改。
	*/
	objProduct.Price += 1.0
	objProduct.Name = "Pro2"
	if _, err := db.UpdateWithKeys(dataBase, objProduct, tablePrefix, map[string]bool{"price": true}); err != nil {
		log.Panic(err)
	} else {
		log.Println("update to", *objProduct)
	}

	productUpdated := &demo.Product{}
	/*
		Query查询，内部不能主动使用id定位记录，需要外部传入条件sql和参数
		查询结果没有自动封装对数据对象，需要通过构造Scaner，将数据行解析为数据对象。
	*/
	if rows, err := db.Query(dataBase, productUpdated, tablePrefix, "where id = ?", objProduct.GetId()); err != nil {
		log.Panic(err)
	} else if rows.Next() {
		if e := db.NewScaner(productUpdated).Scan(rows); e != nil {
			log.Panic(e)
		} else {
			log.Println("query", *productUpdated)
		}
	}

	/*
		QueryWithKeys查询，key指定要查询的字段。
	*/

}
