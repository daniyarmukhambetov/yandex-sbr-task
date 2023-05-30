package dto

import "yandex-team.ru/bstask/models"

type CreateOrder struct {
	Weight        float64  `json:"weight"`
	Region        uint     `json:"region"`
	DeliveryHours []string `json:"delivery_hours"`
	Cost          float64  `json:"cost"`
}

type Order struct {
	OrderID uint `json:"order_id"`
	CreateOrder
	CompletedTime string `json:"completed_time"`
}

type OrderComplete struct {
	CourierId    uint   `json:"courier_id"`
	OrderId      uint   `json:"order_id"`
	CompleteTime string `json:"complete_time"`
}

func TransformOrderModel(orderModel models.Order) Order {
	order := Order{
		OrderID: orderModel.OrderID,
		CreateOrder: CreateOrder{
			Weight:        orderModel.Weight,
			Region:        orderModel.RegionID,
			DeliveryHours: orderModel.DeliveryHours,
			Cost:          orderModel.Price,
		},
	}
	if orderModel.Status != nil {
		order.CompletedTime = orderModel.Status.CompletedTime
	}
	return order
}

func TransformOrderModels(orders []models.Order) []Order {
	response := make([]Order, 0)
	for _, order := range orders {
		response = append(response, TransformOrderModel(order))
	}
	return response
}
