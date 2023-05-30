package dto

import "yandex-team.ru/bstask/models"

type CreateCourier struct {
	CourierType  string   `json:"courier_type"`
	Regions      []uint   `json:"regions"`
	WorkingHours []string `json:"working_hours"`
}

type Courier struct {
	CourierID uint `json:"courier_id"`
	CreateCourier
	Rating   int64 `json:"rating,omitempty"`
	Earnings int64 `json:"earnings,omitempty"`
}

func TransformCourierModels(couriers []models.Courier) []Courier {
	response := make([]Courier, 0)
	for _, courier := range couriers {
		response = append(
			response,
			TransformCourierModel(courier),
		)
	}
	return response
}

func TransformCourierModel(courier models.Courier) Courier {
	regions := make([]uint, 0)
	for _, reg := range courier.Regions {
		regions = append(regions, reg.RegionID)
	}
	return Courier{
		CourierID: courier.CourierID,
		CreateCourier: CreateCourier{
			CourierType:  courier.CourierType,
			Regions:      regions,
			WorkingHours: courier.WorkingHours,
		},
	}
}
