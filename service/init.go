package service

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"im/model"
	"log"
)

var DbEngin *xorm.Engine

func init() {
	driverName := "sqlite3"
	dsName := "./im.db"
	err := errors.New("")
	DbEngin, err = xorm.NewEngine(driverName, dsName)
	if nil != err && "" != err.Error() {
		log.Fatal(err.Error())
	}
	//是否显示SQL语句
	DbEngin.ShowSQL(true)
	//数据库最大打开的连接数
	DbEngin.SetMaxOpenConns(5)

	//自动User
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))
	err = DbEngin.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("init data base ok")
}
