package entity

type Investor struct {
	ID   string
	Name string
	InvestorAssetPosition []*In
}
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}