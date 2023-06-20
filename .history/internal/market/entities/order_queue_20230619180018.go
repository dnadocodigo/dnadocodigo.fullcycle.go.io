package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return od.Orders[i].Price < oq[j].Orders.Pr
}
// swap
// len
// push
// pop