/*
 * @Author: berryberry
 * @LastAuthor: Do not edit
 * @since: 2019-05-10 10:59:09
 * @lastTime: 2019-05-27 10:34:56
 */
package ginfizz

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/berryberrrry/ginfizz/middleware/monitor"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func InitFizz() {
	if FizzConfig.App.Log.Enable {
		initLogger()
	}
	if FizzConfig.App.DB.Enable {
		initDB()
	}
}

func Engine() *gin.Engine {
	if engine == nil {
		engine = gin.Default()

		// if enable limit
		if FizzConfig.App.Limit.Enable {
			engine.Use(monitor.MaxAllowedAndMontiorQueuedProcessingRequest(FizzConfig.App.Limit.MaxAllowed))
		}
		// if enable monitor
		if FizzConfig.Monitor.Enable {
			engine.Use(monitor.Metric)
		}
		return engine
	}
	return engine
}

func Run() {

	endRunning := make(chan struct{}, 3)

	//prometheus
	Logger.Info(fmt.Sprintf("start monitor server listening %d", FizzConfig.Monitor.HttpPort))
	monitorServer := &http.Server{Addr: fmt.Sprintf(":%d", FizzConfig.Monitor.HttpPort), Handler: prometheus.Handler()}
	go func() {
		err := monitorServer.ListenAndServe()
		if err != nil {
			Logger.Errorf("start monitor server error: %s", err)
			endRunning <- struct{}{}
		}
	}()

	Logger.Info(fmt.Sprintf("start http server listening %d", FizzConfig.App.HttpPort))
	fizzServer := &http.Server{Addr: fmt.Sprintf(":%d", FizzConfig.App.HttpPort), Handler: Engine()}
	go func() {
		err := fizzServer.ListenAndServe()
		if err != nil {
			Logger.Errorf("start http server error: %s", err)
			endRunning <- struct{}{}
		}
	}()

	//优雅退出
	go func() {
		sigs := make(chan (os.Signal), 1)
		signal.Notify(sigs, syscall.SIGINT)

		quit := <-sigs

		Logger.Info("Signal received", "type", quit.String())

		ctx, canel := context.WithTimeout(context.Background(), 10*time.Second)
		defer canel()
		shutdown := make(chan struct{}, 1)
		if err := monitorServer.Shutdown(ctx); err != nil {
			Logger.Error("Monitor server shutdown error", err)
		}
		if err := fizzServer.Shutdown(ctx); err != nil {
			Logger.Error("http server shutdown error", err)
		}
		shutdown <- struct{}{}

		select {
		case <-shutdown:
		case <-ctx.Done():
		}
		endRunning <- struct{}{}
		Logger.Info("gracefully shutdown")
	}()

	<-endRunning
}
