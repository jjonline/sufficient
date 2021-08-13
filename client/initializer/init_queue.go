package initializer

import (
	"github.com/jjonline/go-lib-backend/queue"
	"github.com/jjonline/golang-backend/app/console/jobs"
	"github.com/jjonline/golang-backend/client"
	"github.com/jjonline/golang-backend/conf"
)

//go:noinline
func initQueue() *queue.Queue {
	q := queue.New(queue.Redis, client.Redis, client.Logger.Zap, conf.Config.Redis.QueueConcurrent)

	// 注册所有任务类，要求必须注册成功否则阻止启动
	if err := q.Bootstrap(jobs.TaskInstance); err != nil {
		panic("queue task instance bootstrap error, please check queue task at first")
	}

	return q
}
