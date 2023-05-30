package middlewares

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

func RateLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	limiter := rate.NewLimiter(10, 20) //
	return func(c echo.Context) error {
		if limiter.Allow() {
			return next(c)
		}
		return echo.ErrTooManyRequests
	}
}
