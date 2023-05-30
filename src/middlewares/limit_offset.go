package middlewares

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

func LimitOffsetMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var limit int
		var offset int
		if i, err := strconv.Atoi(c.QueryParam("limit")); err != nil {
			limit = 1
		} else {
			limit = i
		}
		if i, err := strconv.Atoi(c.QueryParam("offset")); err != nil {
			offset = 0
		} else {
			offset = i
		}
		c.Set("limit", limit)
		c.Set("offset", offset)
		return next(c)
	}
}
