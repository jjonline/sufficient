package console

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jjonline/golang-backend/client"
	"github.com/jjonline/golang-backend/config"
	"github.com/jjonline/golang-backend/route"
	"github.com/jjonline/golang-backend/utils"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

// BootStrap 引导启动
func BootStrap() {
	// output base info
	client.Logger.Info(fmt.Sprintf("Golang Version  : %s", runtime.Version()))
	client.Logger.Info(fmt.Sprintf("MAX Cpu Num     : %d", runtime.GOMAXPROCS(-1)))
	client.Logger.Info(fmt.Sprintf("Command Args    : %s", os.Args))
	client.Logger.Info(fmt.Sprintf("Log Path        : %s", config.Config.Log.Path))

	_, signalChan := quitCtx()

	// 仅启动队列
	if config.Cmd.OnlyQueue {
		return
	}

	// 仅启动定时任务
	if config.Cmd.OnlyCrontab {
		return
	}

	// 跟随启动队列
	if config.Cmd.WithQueue {

	}

	// 跟随启动定时任务
	if config.Cmd.WithCrontab {

	}
	startHttpApp(signalChan)
}

// startHttpApp http api服务启动
func startHttpApp(signalChan chan os.Signal) {
	// 启动模式
	gin.SetMode(utils.RunMode())

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.Config.Server.Port),
		Handler:        route.Bootstrap(),
		ReadTimeout:    time.Duration(config.Config.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Config.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20, //1MB
	}

	// http serv handle exit signal
	go func() {
		<-signalChan

		// 超时context
		timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer timeoutCancel()
		// We received an interrupt signal, shut down.
		if err := server.Shutdown(timeoutCtx); err != nil {
			// Error from closing listeners, or context timeout:
			client.Logger.Error("Http服务暴力停止：" + err.Error())
		} else {
			// time.Sleep(1 * time.Second)
			// successful shutdown process ok
			client.Logger.Info("Http服务优雅停止")
		}
		// closer quit chan
		close(signalChan)
	}()

	// continue serv http service
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		client.Logger.Error("Http服务异常：" + err.Error())
		close(signalChan)
	}

	// wait for stop main process
	<-signalChan
	client.Logger.Info("进程已退出：服务已关闭")
}

// quitCtx 全局退出信号
func quitCtx() (context.Context, chan os.Signal) {
	ctx, cancel := context.WithCancel(context.Background())
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	go func() {
		<-quitChan // wait quit signal

		signal.Stop(quitChan)
		cancel()
		close(quitChan)
	}()

	return ctx, quitChan
}
