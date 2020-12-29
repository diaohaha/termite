package db

import (
	"github.com/diaohaha/termite/dal"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

var RODB *gorm.DB
var WTDB *gorm.DB

func InitDB() {
	// 初始化数据库资源
	connect_read_db_addr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dal.Env.DB_read.Username,
		dal.Env.DB_read.Password,
		dal.Env.DB_read.Host,
		dal.Env.DB_read.Port,
		dal.Env.DB_read.Name,
	)
	rdb, err := gorm.Open("mysql", connect_read_db_addr)
	if err != nil {
		panic(err.Error())
	}
	RODB = rdb

	connect_write_db_addr := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dal.Env.DB_default.Username,
		dal.Env.DB_default.Password,
		dal.Env.DB_default.Host,
		dal.Env.DB_default.Port,
		dal.Env.DB_default.Name,
	)
	wdb, err := gorm.Open("mysql", connect_write_db_addr)
	if err != nil {
		panic(err.Error())
	}
	WTDB = wdb

	// debug
	//RODB.LogMode(true)
	//WTDB.LogMode(true)
	//RODB.LogMode(false)

	WTDB.DB().SetMaxIdleConns(0)
	WTDB.DB().SetMaxOpenConns(200)
	WTDB.DB().SetConnMaxLifetime(30 * time.Second)
	RODB.DB().SetMaxIdleConns(0)
	RODB.DB().SetMaxOpenConns(200)
	RODB.DB().SetConnMaxLifetime(30 * time.Second)

	//go func() {
	//	for {
	//		err = WTDB.DB().Ping()
	//		if err != nil {
	//			stable.Logger.Error("DB PING ERROR!", zap.Error(err))
	//		}
	//		time.Sleep(time.Second * 50)
	//	}
	//}()

}
