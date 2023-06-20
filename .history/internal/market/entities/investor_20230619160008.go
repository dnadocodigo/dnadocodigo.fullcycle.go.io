package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor( )
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}