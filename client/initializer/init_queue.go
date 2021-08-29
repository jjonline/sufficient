package initializer

import (
	"github.com/jjonline/go-lib-backend/queue"
	"github.com/jjonline/sufficient/app/console/job"
	"github.com/jjonline/sufficient/client"
	"github.com/jjonline/sufficient/conf"
)

//go:noinline
func initQueue() *queue.Queue {
	q := queue.New(queue.Redis, client.Redis, client.Logger.Zap, conf.Config.Redis.QueueConcurrent)

	// 注册所有任务类，要求必须注册成功否则阻止启动
	if err := q.Bootstrap(job.TaskInstance); err != nil {
		panic("queue task instance bootstrap error, please check queue task at first")
	}

	return q
}
