package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id )
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}