package entity

type OrderQueue struct{
	Orders []*Order
}

// less
func (oq OrderQueue) Less(i, j int) bool
// swap
// len
// push
// pop