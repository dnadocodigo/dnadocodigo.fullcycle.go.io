package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return oq.Orders[i].Price < oq.Orders[j].Price
}
// swap
func (op OrderQueue) Swap(i)
// len
// push
// pop