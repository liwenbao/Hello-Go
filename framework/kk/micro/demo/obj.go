package demo

import (
	"git.kkcoding.com/kk-golang/kk-lib/db"
)

type DemoObject struct {
	db.Object
	Name        string `json:"name" length:"128"`
	Vaild       bool   `json:"valid"`
	Description string `json:"desc" length:"-1"`
}
