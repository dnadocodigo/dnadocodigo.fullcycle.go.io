package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool{
	return oq.Orders[i].Price < oq.Orders[j].Price
}
// swap
func (oq OrderQueue) Swap(i, j int){
	return oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// len
func (oq O )
// push
// pop