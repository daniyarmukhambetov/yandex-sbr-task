package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"yandex-team.ru/bstask/dto"
	"yandex-team.ru/bstask/models"
)

func (h *Handler) GetCouriers(ctx echo.Context) error {
	var couriers []models.Courier
	if err := h.db.Limit(ctx.Get("limit").(int)).Offset(ctx.Get("offset").(int)).Preload("Regions").Find(&couriers).Error; err != nil {
		return err
	}
	response := dto.TransformCourierModels(couriers)
	return ctx.JSON(200, map[string][]dto.Courier{"couriers": response})
}

func (h *Handler) CreateCouriers(ctx echo.Context) error {
	data := struct {
		Couriers []dto.CreateCourier `json:"couriers"`
	}{}
	var newCouriers []dto.CreateCourier
	if err := ctx.Bind(&data); err != nil {
		return err
	}
	newCouriers = data.Couriers
	couriers := make([]models.Courier, 0)
	tx := h.db.Begin()
	var txErr error
	for _, courier := range newCouriers {
		regions := make([]models.Region, 0)
		for _, reg := range courier.Regions {
			regions = append(
				regions,
				models.Region{
					RegionID: reg,
				},
			)
		}
		newCourier := models.Courier{
			CourierType:  courier.CourierType,
			WorkingHours: courier.WorkingHours,
		}
		if err := tx.Create(&newCourier).Error; err != nil {
			txErr = err
			break
		}
		if err := tx.Model(&newCourier).Association("Regions").Append(regions); err != nil {
			txErr = err
			break
		}
		couriers = append(couriers, newCourier)
	}
	if txErr != nil {
		tx.Rollback()
		return txErr
	}
	tx.Commit()
	return ctx.JSON(200, map[string][]dto.Courier{"couriers": dto.TransformCourierModels(couriers)})
}

func (h *Handler) GetCourier(ctx echo.Context) error {
	var courier models.Courier
	id, bindErr := strconv.Atoi(ctx.Param("id"))
	if bindErr != nil {
		return bindErr
	}
	if err := h.db.Preload("Regions").Find(&courier, id).Error; err != nil {
		return err
	}
	return ctx.JSON(200, dto.TransformCourierModel(courier))
}

func (h *Handler) GetCourierMetaInfo(ctx echo.Context) error {
	var courier models.Courier
	id, bindErr := strconv.Atoi(ctx.Param("id"))
	if bindErr != nil {
		return bindErr
	}
	if err := h.db.Preload("Regions").Find(&courier, id).Error; err != nil {
		return err
	}
	startDate := ctx.QueryParam("start_date")
	endDate := ctx.QueryParam("end_date")
	stats := struct {
		Cnt int64
		Sm  int64
	}{}

	sql := fmt.Sprintf(
		"SELECT COUNT(orders.order_id) AS cnt, SUM(orders.price) AS sm FROM orders "+
			"RIGHT OUTER JOIN order_statuses "+
			"ON orders.order_id = order_statuses.order_id "+
			"WHERE order_statuses.courier_id = %d AND "+
			"order_statuses.completed_date >= '%s' AND order_statuses.completed_date < '%s';",
		id, startDate, endDate,
	)
	if err := h.db.Raw(sql).Scan(&stats).Error; err != nil {
		return err
	}
	resp := dto.TransformCourierModel(courier)
	resp.Rating = stats.Cnt
	resp.Earnings = stats.Sm
	return ctx.JSON(200, resp)
}
