package task

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

// 全局定时任务调度器

var cronScheduler *cron.Cron

// 初始化

func Init() (err error) {
	// 支持秒级的定时器
	cronScheduler = cron.New(cron.WithSeconds())

	// 注册任务
	_, err = cronScheduler.AddFunc("0 */30 * * * *", RefreshAndSettle)
	if err != nil {
		return err
	}
	// 启动调度器
	cronScheduler.Start()
	zap.L().Info("Cron scheduler started")
	return nil
}

// TriggerManualRefreshAndSettle 暴露方法：手动触发任务（可供外部调用）
func TriggerManualRefreshAndSettle() {
	go RefreshAndSettle()
}
