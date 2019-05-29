/*
 * @Author: berryberry
 * @since: 2019-05-29 20:01:44
 * @lastTime: 2019-05-29 20:25:17
 * @LastAuthor: Do not edit
 */
package pprof

import (
	"net/http"
	"net/http/pprof"

	"github.com/gin-gonic/gin"
)

// https://github.com/gin-contrib/pprof/blob/master/pprof.go
func AllowPprof(engine *gin.Engine) {
	p := engine.Group("/debug/pprof")
	{
		p.GET("/", pprofHandler(pprof.Index))
		p.GET("/cmdline", pprofHandler(pprof.Cmdline))
		p.GET("/profile", pprofHandler(pprof.Profile))
		p.POST("/symbol", pprofHandler(pprof.Symbol))
		p.GET("/symbol", pprofHandler(pprof.Symbol))
		p.GET("/trace", pprofHandler(pprof.Trace))
		p.GET("/allocs", pprofHandler(pprof.Handler("allocs").ServeHTTP))
		p.GET("/block", pprofHandler(pprof.Handler("block").ServeHTTP))
		p.GET("/goroutine", pprofHandler(pprof.Handler("goroutine").ServeHTTP))
		p.GET("/heap", pprofHandler(pprof.Handler("heap").ServeHTTP))
		p.GET("/mutex", pprofHandler(pprof.Handler("mutex").ServeHTTP))
		p.GET("/threadcreate", pprofHandler(pprof.Handler("threadcreate").ServeHTTP))
	}
}

func pprofHandler(h http.HandlerFunc) gin.HandlerFunc {
	handler := http.HandlerFunc(h)
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}
