package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string)
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}