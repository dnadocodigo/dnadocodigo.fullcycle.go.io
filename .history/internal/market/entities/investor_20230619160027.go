package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string){
	re
}
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}