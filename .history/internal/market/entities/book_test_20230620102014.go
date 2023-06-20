package entity

import (
	"sync"
	"testing"
)

func TestBuyAsset(t *testing.T){
	asset1 := NewAsset("asset1", "Asset 1", 100)

	investor := NewInvestor("1")
	investor2 := NewInvestor("2")

	investorAssetPosition := NewInvestorAssetPosition("asset1", 10)
	investor.AddAssetPosition(investorAssetPosition)

	wg := sync.WaitGroup{}
	orderChan := make(chan *Order)
	orderChanOut := make(chan *Order)

	book := NewBook(orderChan, orderChanOut, &wg)
	go book.Trade()

	// add buy order
	wg.Add(1)
	order := NewOrder("1", investor, asset1, 5, 5, "SELL")
	orderChan <- order

	// add sell order
	order2 := NewOr

}