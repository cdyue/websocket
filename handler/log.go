package handler

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-stack/stack"
	"github.com/inconshreveable/log15"
)

//Log 日志
func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		log := log15.New("Method", context.Request.Method, "Api", context.Request.URL.Path, "IP", context.ClientIP())
		handler := log15.StreamHandler(os.Stdout, log15.TerminalFormat())
		h := CallerHandler("[ %+v ]", handler)
		log.SetHandler(h)

		context.Set("log", log)
		context.Next()

	}
}

//CallerHandler 显示caller完整信息，包括package/source:line
func CallerHandler(format string, h log15.Handler) log15.Handler {
	return log15.FuncHandler(func(r *log15.Record) error {
		s := stack.Trace().TrimBelow(r.Call).TrimRuntime()
		if len(s) > 0 {
			r.Ctx = append(r.Ctx, "caller", fmt.Sprintf(format, s[0]))
		}
		return h.Log(r)
	})
}
