package entity

type Order struct {
	ID            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	PendingShares int
	Price         float64
	OrderType     string
	Status        string
	Transactions []*Transaction
}

func NewOrder(id string, i){
	return &Order{
		ID: id,
	}
}
