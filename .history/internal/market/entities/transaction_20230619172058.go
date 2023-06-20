package entity

import "time"

type Transaction struct{
	ID string
	SellingOrder *Order
	BuyingOrder *Order
	Shares int
	Price float64
	Total float64
	DateTime time.Time
}

func New