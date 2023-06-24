package tranformers

import (
	"github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/dto"
	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
	"golang.org/x/text/encoding"
)

func tranformerInput(input dto.TradeInput) *entity.Order{
	asset:= entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor:= entity.NewInvestor(input.InvestorID)
	order:= entity.NewOrder(input.OrderID, investor, asset, input.Shares, input.Price, input.OrderType)
	if input.CurrentShares > 0{
		assetPosition:= entity.NewInvestorAssetPosition(input.AssetID, input.CurrentShares)
		investor.
	}
}