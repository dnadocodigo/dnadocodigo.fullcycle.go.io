package entity

import "sync"

type Book struct {
	Order         []*Order
	Transaction   []*Transaction
	OrdersChan    chan *Order
	OrdersChanOut chan *Order
	Wg            *sync.WaitGroup
}
func NewBook(orderChan chan *Order, orderChanOut chan *Order, wg *sync.WaitGroup){
	return &Book{
		Order: []*Order[],
		Transaction: []*Transaction{},
		OrdersChan: orderChan,
		OrdersChanOut: orderChanOut,
		Wg: wg,
	}
}
func(b *Book) Trade(){
	buyOrders ;= NewOrderQueue();
	sellOrder 
}

