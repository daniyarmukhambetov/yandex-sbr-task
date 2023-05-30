package models

type Order struct {
	OrderID       uint `gorm:"primaryKey"`
	Weight        float64
	RegionID      uint
	Region        Region
	DeliveryHours Hours
	Price         float64
	Status        *OrderStatus
	Date          string
}

type OrderAssign struct {
	OrderID   uint `gorm:"uniqueIndex"`
	Order     Order
	CourierID uint
	Courier   Courier
	GroupID   uint
}

type OrderStatus struct {
	OrderID       uint `gorm:"uniqueIndex"`
	CourierID     uint
	Courier       Courier
	CompletedTime string
	CompletedDate string
}
