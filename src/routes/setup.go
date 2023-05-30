package routes

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/handler"
	"yandex-team.ru/bstask/middlewares"
)

func SetupRoutes(e *echo.Echo, h *handler.Handler) {
	e.GET("/ping", ping)
	e.GET("/couriers", middlewares.LimitOffsetMiddleware(middlewares.RateLimitMiddleware(h.GetCouriers)))
	e.POST("/couriers", middlewares.RateLimitMiddleware(h.CreateCouriers))
	e.GET("/couriers/:id", middlewares.RateLimitMiddleware(h.GetCourier))
	e.GET("/couriers/meta_info/:id", middlewares.RateLimitMiddleware(h.GetCourierMetaInfo))
	e.GET("/orders", middlewares.LimitOffsetMiddleware(middlewares.RateLimitMiddleware(h.GetOrders)))
	e.POST("/orders", middlewares.RateLimitMiddleware(h.CreateOrders))
	e.GET("/orders/:id", middlewares.RateLimitMiddleware(h.GetOrder))
	e.POST("/orders/complete", middlewares.RateLimitMiddleware(h.OrdersComplete))
}
