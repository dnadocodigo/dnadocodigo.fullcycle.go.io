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

	wg : *&sync.WaitGroup{}
	

}