package main

import (
	"fmt"
	"web-app/dao/mysql"
	"web-app/logger"
	"web-app/settings"
)

func main() {
	// 1、加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("config init faild %s\n", err)
		return
	}
	// 2、初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("loger init faild %s\n", err)
		return
	}
	// 3、初始化数据库连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init faild %s\n", err)
		return
	}
	// 4、初始化redis连接
	// 5、初始化路由
	// 6、启动服务优化关机

}
