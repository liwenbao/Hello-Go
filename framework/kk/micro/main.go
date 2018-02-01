package main

import (
	"Hello-Go/framework/kk/micro/demo"
	"log"
	"net/http"
	"time"

	"git.kkcoding.com/kk-golang/kk-lib/kk"
	"git.kkcoding.com/kk-golang/kk-lib/micro"
)

func main() {
	log.Println("Hellow Micro")
	var app = demo.DemoApp{}
	app.Name = "/demo"
	app.DB = &micro.Database{}
	app.DB.AutoIncrement = 1
	app.DB.Name = "mysql"
	app.DB.Prefix = "study_"
	app.DB.Url = "root:@tcp(127.0.0.1:5566)/study?charset=utf8&sql_mode=ANSI"
	micro.Load(&app)
	go func() {
		http.HandleFunc("/", micro.HandleFunc(&app))
		var svr = http.Server{
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 5,
			Addr:         ":8001",
		}
		log.Println(svr.ListenAndServe())
	}()
	kk.DispatchMain()
}
