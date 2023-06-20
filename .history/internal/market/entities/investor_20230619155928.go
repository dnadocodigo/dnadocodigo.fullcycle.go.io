package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

fun
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}