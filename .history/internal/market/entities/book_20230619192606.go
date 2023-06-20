package entity

import (
	"container/heap"
	"sync"
)

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
	buyOrders := NewOrderQueue()
	sellOrder := NewOrderQueue()

	heap.Init(buyOrders)
	heap.Init(sellOrders)

	for order := range b.OrdersChan{
		if order.OrderType == "BUY"{
			buyOrders.Push(order){
				if sellOrder.Len() > 0 && sellOrder.Orders[0].Price 
			}
		}
	}
}

