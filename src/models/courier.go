package models

type Courier struct {
	CourierID    uint     `gorm:"primaryKey" json:"courier_id"`
	CourierType  string   `json:"courier_type"`
	Regions      []Region `gorm:"many2many:courier_region;" json:"regions"`
	WorkingHours Hours
}

type Region struct {
	RegionID   uint `gorm:"primaryKey"`
	RegionName string
	Couriers   []Courier `gorm:"many2many:courier_region;"`
}
