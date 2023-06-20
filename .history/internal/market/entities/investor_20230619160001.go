package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(ID)
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}