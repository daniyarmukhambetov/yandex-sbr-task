package handler

import (
	"github.com/labstack/echo/v4"
	"yandex-team.ru/bstask/dto"
	"yandex-team.ru/bstask/models"
)

func (h *Handler) GetOrders(ctx echo.Context) error {
	var orders []models.Order
	if err := h.db.Limit(ctx.Get("limit").(int)).Offset(ctx.Get("offset").(int)).Preload("Status").Find(&orders).Error; err != nil {
		return err
	}
	return ctx.JSON(200, dto.TransformOrderModels(orders))
}

func (h *Handler) CreateOrders(ctx echo.Context) error {
	data := struct {
		Orders []dto.CreateOrder `json:"orders"`
	}{}
	if err := ctx.Bind(&data); err != nil {
		return ctx.JSON(400, err.Error())
	}
	newOrders := make([]models.Order, 0)
	tx := h.db.Begin()
	var txErr error
	for _, order := range data.Orders {
		newOrder := models.Order{
			Weight:        order.Weight,
			RegionID:      order.Region,
			DeliveryHours: order.DeliveryHours,
			Price:         order.Cost,
		}
		if err := tx.Create(&newOrder).Error; err != nil {
			txErr = err
			break
		}
		newOrders = append(newOrders, newOrder)
	}
	if txErr != nil {
		tx.Rollback()
		return txErr
	}
	tx.Commit()
	return ctx.JSON(200, dto.TransformOrderModels(newOrders))
}

func (h *Handler) GetOrder(ctx echo.Context) error {
	var order models.Order
	if err := h.db.Preload("Status").Find(&order, ctx.Param("id")).Error; err != nil {
		return err
	}
	return ctx.JSON(200, dto.TransformOrderModel(order))
}

func (h *Handler) OrdersComplete(ctx echo.Context) error {
	data := struct {
		CompleteInfo []dto.OrderComplete `json:"complete_info"`
	}{}
	if err := ctx.Bind(&data); err != nil {
		return err
	}
	tx := h.db.Begin()
	orderIds := make([]uint, 0)
	var txError error
	for _, info := range data.CompleteInfo {
		orderInfo := models.OrderStatus{
			OrderID:       info.OrderId,
			CourierID:     info.CourierId,
			CompletedTime: info.CompleteTime,
		}
		if err := tx.Create(&orderInfo).Error; err != nil {
			txError = err
			break
		}
		orderIds = append(orderIds, info.OrderId)
	}
	if txError != nil {
		tx.Rollback()
		return txError
	}
	tx.Commit()
	var orders []models.Order
	h.db.Preload("Status").Find(&orders)
	return ctx.JSON(201, dto.TransformOrderModels(orders))
}
