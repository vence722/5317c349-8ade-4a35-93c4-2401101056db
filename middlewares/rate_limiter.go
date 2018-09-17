package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)

func SetRateLimiterMiddleware() {
	// rate limiter
	rateLimiter := getRateLimiter()
	beego.InsertFilter("/*", beego.BeforeRouter, func(c *context.Context) {
		rateLimit(rateLimiter, c)
	}, true)

	// error handling
	beego.ErrorHandler("429", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTooManyRequests)
		w.Write([]byte("Too Many Requests"))
		return
	})
}

func getRateLimiter() *limiter.Limiter {
	// read api rate limit from config
	env := beego.AppConfig.String("runmode")
	apiRateLimit := beego.AppConfig.String(env + "::api_rate_limit")
	rate, err := limiter.NewRateFromFormatted(apiRateLimit)
	if err != nil {
		panic(err)
	}
	return limiter.New(memory.NewStore(), rate)
}

func rateLimit(l *limiter.Limiter, ctx *context.Context) {
	var limiterCtx limiter.Context
	var err error
	req := ctx.Request

	ip := limiter.GetIP(req, false)
	limiterCtx, err = l.Get(req.Context(), ip.String())
	if err != nil {
		ctx.Abort(http.StatusInternalServerError, err.Error())
		return
	}

	h := ctx.ResponseWriter.Header()
	h.Add("X-RateLimit-Limit", strconv.FormatInt(limiterCtx.Limit, 10))
	h.Add("X-RateLimit-Remaining", strconv.FormatInt(limiterCtx.Remaining, 10))
	h.Add("X-RateLimit-Reset", strconv.FormatInt(limiterCtx.Reset, 10))

	if limiterCtx.Reached {
		fmt.Printf("Too Many Requests from %s on %s\n", ip, ctx.Input.URL())
		ctx.Abort(http.StatusTooManyRequests, "429")
		return
	}
}
