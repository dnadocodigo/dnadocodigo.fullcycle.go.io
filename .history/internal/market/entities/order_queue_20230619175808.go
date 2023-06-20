package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return od.Orders.Price < o
}
// swap
// len
// push
// pop