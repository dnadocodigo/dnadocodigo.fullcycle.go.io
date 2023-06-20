package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string){
	return $Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition()
	}
}
type InvestorAssetPosition struct{
	AssetID string
	Shares int
}