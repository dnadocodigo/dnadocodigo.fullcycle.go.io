package entity

type Investor struct {
	ID   string
	Name string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string)*Investor{
	return &Investor{
		ID: id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}
func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition){
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func(i *Investor) GetAssetPosition(assetID string)*InvestorAssetPosition{
	for , assetPosition := range i.AssetPosistion{
		if assetPosition.assetID == assetID{
			return assetPosistion
		}
		return nil
	}
}

type InvestorAssetPosition struct{
	AssetID string
	Shares int
}