package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return o.Orders[i].Price < oq[j].Orders.Price
}
// swap
// len
// push
// pop