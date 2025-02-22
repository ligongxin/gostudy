package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web-app/controller"
	"web-app/dao/mysql"
	"web-app/dao/redis"
	"web-app/logger"
	"web-app/pkg/snowflake"
	"web-app/router"
	"web-app/settings"
	"web-app/task"
)

// @title 投票
// @version 1.0.00
// @description 这里写描述信息
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath
func main() {
	// 1、加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("config init faild %s\n", err)
		return
	}
	// 2、初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("loger init faild %s\n", err)
		return
	}
	zap.L().Debug("初始化日志成功")
	defer zap.L().Sync()
	// 3、初始化数据库连接
	if err := mysql.Init(settings.Conf.MysqlConfig); err != nil {
		fmt.Printf("mysql init faild %s\n", err)
		return
	}
	defer mysql.Close()

	// 4、初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init faild %s\n", err)
		return
	}
	defer redis.Close()

	// 初始化全局翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("Trans init faild %s\n", err)
		return
	}
	// 雪花生成id初始化
	if err := snowflake.Init("2006-01-02", 1); err != nil {
		fmt.Printf("snowflake Init faild %s\n", err)
		return
	}
	// 注册定时任务
	if err := task.Init(); err != nil {
		fmt.Printf("Failed to add cron job: %v", err)
		return
	}
	// 5、初始化路由
	r := router.SetupRoute(settings.Conf.Mode)
	// 6、启动服务优化关机
	src := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := src.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	fmt.Println("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := src.Shutdown(ctx); err != nil {
		fmt.Printf("server Shutdown fail %s\n", err)
	}
	fmt.Println("Server exiting")
}
