package algo

type Order struct {
	ID               int
	DeliveryInterval []int64
	Region           int
	Weight           float64
	Price            int
	DeliveredAt      int64
	WaitTime         int64
}

type Group struct {
	ID     int
	Orders []Order
}

func (g *Group) AddOrder(ord Order, crType string) {
}

type Assign struct {
	CourierID    int
	WorkingHours []int64
	Type         string
	Regions      []int
	Groups       []Group
}
