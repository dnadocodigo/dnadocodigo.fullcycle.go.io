package tranformers

import (
	"github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/dto"
	entity "github.com/dnadocodigo/dnadocodigo.fullcycle.go.io/internal/market/entities"
)

func tranformerInput(input dto.TradeInput) *entity.Order{
	asset:= entity.NewAsset(input.AssetID, input.AssetID, 1000)
	investor:= ene
}