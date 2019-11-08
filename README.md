[![Build Status](https://travis-ci.org/berryberrrry/ginfizz.svg?branch=master)](https://travis-ci.org/berryberrrry/ginfizz)

# GIN-FIZZ: Golang Web Framework Based On Gin

## **Contents:**

* [Installation](#Installation)
* [Quick start](#Quick-start)
* [支持的组件与使用](#支持的组件与使用)

### **Installation**

To install this package, you need to install Go(version 1.10+ is required) and set Go workspace first.

then

```
$ go get -u github.com/berryberrrry/ginfizz
```

### **Quick start**

```
    package main

    import (
        "github.com/berryberrrry/ginfizz"
        "github.com/gin-gonic/gin"
    )

    func main() {
        # Config (see more infomation in ./conf.go)
        ginfizz.FizzConfig.App.Log.LogLevel = "debug"
        ginfizz.FizzConfig.App.HttpPort = 7777
        ginfizz.FizzConfig.App.Log.LogRotator.Filename = "ginfizz.log"
        ginfizz.FizzConfig.App.DB.Enable = false

        # init
        ginfizz.InitFizz()

        engine := ginfizz.Engine()

        engine.GET("/hello-ginfizz", func(c *gin.Context) {
            c.JSON(200, map[string]interface{}{"say": "hello gin-fizz"})
        })
        ginfizz.Logger.Info("hello gin-fizz")
        ginfizz.Run()
    }
```

### **支持的组件与使用**

Tips: (你可以在 *ginfizz.InitFizz()* 之前进行项目变量的配置，控制不同组件的启用和参数)

#### **日志管理(logger)**

*ginfizz* 帮你集成了 *go.uber.org/zap* ,你只需要修改配置来控制复杂的 *logger* 逻辑.

你可以在初始化之前进行 *Logger* 的配置,通过修改 *ginfizz.FizzConfig.App.Log* 来控制logger的存储路径、存储文件名、logLevel、log文件的最大存储容量、log文件时效以及是否压缩和压缩后的文件名等等，详情可以参考 *conf.go* .
你也可以通过[go.uber.org/zap](https://github.com/uber-go/zap) 和 [gopkg.in/natefinch/lumberjack.v2](https://github.com/natefinch/lumberjack) 了解更改关于 *logger* 的信息

#### **Prometheus监控**

*ginfizz* 集成了[Prometheus](https://github.com/prometheus/client_golang),并设置了几个指标，帮助用户监控自己的服务:

```
prometheus.MustRegister(
    httpRequestCount,
    httpRequestDuration,
    httpQueuedRequestCount,
    httpProcessingRequestCount,
)
```

细心的你可能已经发现,*ginfizz* 可以帮你开启一个 *monitor server* ，默认的端口为10010,,你可以通过这个端口查看 *Prometheus* 的各项指标。

#### **限流**

*ginfizz* 默认添加了限流中间件，能够帮你限制同时处理的最大事务数。你可以通过修改 *ginfizz.FizzConfig.App.Limit.MaxAllowed=100* 来调整最大事务数，当然你也可以任性的关闭他，但我想你应该不会这么做

#### **PProf性能分析**

*ginfizz* 配置了 *PProf* ，来帮助你分析服务的性能，你可以通过 *http://127.0.0.1:8080/debug/pprof/*

```
cpu(CPU Profiling): /debug/pprof/profile，默认进行 30s 的 CPU Profiling，得到一个分析用的 profile 文件
block(Block Profiling): /debug/pprof/block，查看导致阻塞同步的堆栈跟踪
goroutine: /debug/pprof/goroutine，查看当前所有运行的 goroutines 堆栈跟踪
heap(Memory Profiling): /debug/pprof/heap，查看活动对象的内存分配情况
mutex(Mutex Profiling): /debug/pprof/mutex，查看导致互斥锁的竞争持有者的堆栈跟踪
threadcreate: /debug/pprof/threadcreate，查看创建新OS线程的堆栈跟踪
```

更多关于 *pprof* 的使用请参考[这里](https://golang.org/pkg/net/http/pprof/)

#### **优雅退出**

当 *ginfizz* 接收到退出信号时，此时，*ginfizz* 会给服务预留一个缓冲时间(默认为10s)，这段时间 *ginfizz* 放弃接受新的请求，处理还在处理的请求直到超时或者请求处理完停止server以及监控server，帮助程序优雅退出。

#### **DingTalk报警机器人**

为了方便用户及时发现服务运行的异常情况，我们封装了钉钉报警接口，配置好你的webhook，你可以更方便的调用接口来给钉钉发送消息(包括普通文本和markdown格式的文本)。

```
import (
    "github.com/berryberrrry/ginfizz/pkg/dingtalk"
)

func Test() {
    webhook := "你的webhook"
    bot := dingtalk.New(webhook)
    //发送普通文本
    err := bot.SendText("this is text")
    if err != nil {
        panic(err)
    }
}
```

#### **更容易用的mongo**
*ginfizz* 帮你封装了mongodb的一系列操作，提供了一系列类似于gorm的接口，帮助用户更好的编写mongo数据库操作