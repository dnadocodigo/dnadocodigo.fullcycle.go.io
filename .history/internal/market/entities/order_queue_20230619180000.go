package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return od.Orders[i].Price < oq[j].Order
}
// swap
// len
// push
// pop