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
			buyOrders.Push(order)
				if sellOrder.Len() > 0 && sellOrder.Orders[0].Price <= order.Price{
					sellOrder := sellOrder.Pop().(*Order)
					if sellOrder.PendingShares > 0{
						transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
						b.AddTransaction(transaction, b.Wg)
						sellOrder.Transactions = append(sellOrder.Transactions, transaction)
						order.Transactions = append(order.Transactions, transaction)
						b.OrdersChanOut <- sellOrder
						b.OrdersChanOut <- order
						if sellOrder.PendingShares > 0 {
							sellOrders.Push(sellOrder)
						}
					}
				}
			} else if order.OrderType == "SELL"{
				sellOrders.Push(order)
				if buyOrders.len() >0 && buyOrders.Orders[0].Price >= order.Price{
					buyOrder := buyOrders.Pop().(*Order)
					if buyOrder.PendingShares >0{
						transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
						b.AddTransaction(transaction, b.Wg)
						buyOrder.Transactions = append(buyOrder.Transactions, transaction) 
						order.Transactions = append(order.Transactions, transaction)
						b.OrdersChanOut <- buyOrder
						b.OrdersChanOut <- order
						if buyOrder.PendingShares > 0{
							buyOrders.Push(buyOrder)
						}
					}
				}
			}
		}
	}
func (b *Book) 
