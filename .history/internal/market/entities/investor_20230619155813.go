package entity

type Investor struct {
	ID   string
	Name string
	InvestorAssetPosition []*InvestorAssetPosition
}
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}