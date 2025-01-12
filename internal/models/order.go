package models

type OrderType string

const (
	BuyOrder  OrderType = "B"
	SellOrder OrderType = "S"
)

type Order struct {
	ID           int32
	Type         OrderType
	Price        int32
	Quantity     int32
	ArrivalOrder int32
}
