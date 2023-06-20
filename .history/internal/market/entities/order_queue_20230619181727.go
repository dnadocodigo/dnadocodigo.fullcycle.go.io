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
	return oq.Orders[i].Price, oq.Orders[j].Price = oq.Orders[j].Price, oq.Orders[i].p
}

// len
func (oq OrderQueue) Len(i, j int){
	return oq.Orders[]
}
// push
// pop