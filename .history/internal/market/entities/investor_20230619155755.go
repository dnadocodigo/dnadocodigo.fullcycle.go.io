package entity

type Investor struct {
	ID   string
	Name string
	InvestorAssetPosition
}
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}