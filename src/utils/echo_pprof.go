package utils

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http/pprof"
)

var h = middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
	Skipper: nil,
	Validator: func(u, p string, c echo.Context) (bool, error) {
		if u == "debug" && p == "N^52nsq^YGkOp@%d" {
			return true, nil
		}
		return false, nil
	},
	Realm: "someRealm_debug",
})

func PProf(e *echo.Echo) {
	e.GET("/debug/pprof", IndexHandler, h)
	e.GET("/debug/heap", HeapHandler, h)
	e.GET("/debug/goroutine", GoroutineHandler, h)
	e.GET("/debug/allocs", AllocsHandler, h)
	e.GET("/debug/block", BlockHandler, h)
	e.GET("/debug/threadcreate", ThreadCreateHandler, h)
	e.GET("/debug/cmdline", CmdlineHandler, h)
	e.GET("/debug/profile", ProfileHandler, h)
	e.POST("/debug/symbol", SymbolHandler, h)
	e.GET("/debug/trace", TraceHandler, h)
	e.GET("/debug/mutex", MutexHandler, h)
}

func IndexHandler(ctx echo.Context) error {
	pprof.Index(ctx.Response().Writer, ctx.Request())
	return nil
}

func HeapHandler(ctx echo.Context) error {
	pprof.Handler("heap").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func GoroutineHandler(ctx echo.Context) error {
	pprof.Handler("goroutine").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func AllocsHandler(ctx echo.Context) error {
	pprof.Handler("allocs").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func BlockHandler(ctx echo.Context) error {
	pprof.Handler("block").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func ThreadCreateHandler(ctx echo.Context) error {
	pprof.Handler("threadcreate").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}

func CmdlineHandler(ctx echo.Context) error {
	pprof.Cmdline(ctx.Response().Writer, ctx.Request())
	return nil
}

func ProfileHandler(ctx echo.Context) error {
	pprof.Profile(ctx.Response().Writer, ctx.Request())
	return nil
}

func SymbolHandler(ctx echo.Context) error {
	pprof.Symbol(ctx.Response().Writer, ctx.Request())
	return nil
}

func TraceHandler(ctx echo.Context) error {
	pprof.Trace(ctx.Response().Writer, ctx.Request())
	return nil
}

func MutexHandler(ctx echo.Context) error {
	pprof.Handler("mutex").ServeHTTP(ctx.Response().Writer, ctx.Request())
	return nil
}
