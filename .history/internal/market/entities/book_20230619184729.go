package entity

type Book struct{
	Order []*Order
	Transaction []*Transaction
	OrdersChan chan *Order
	OrdersChanOut chan *Order
	Wg *sync.Wa
}