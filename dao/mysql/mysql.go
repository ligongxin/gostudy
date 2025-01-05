package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"web-app/settings"
)

var db *sqlx.DB

func Init(conf settings.MysqlConfig) (err error) {

	//dsn := "root:123456@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Dbname)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败")
		return err
	}
	db.SetMaxIdleConns(conf.MaxIdleConns) // 最大空闲数
	db.SetMaxOpenConns(conf.MaxOpenConns) // 最大连接数
	return
}
