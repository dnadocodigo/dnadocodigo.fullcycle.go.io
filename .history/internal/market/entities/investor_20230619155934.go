package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func New
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}