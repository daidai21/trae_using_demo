package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type AuctionProductExtension struct{}

func (e *AuctionProductExtension) Name() string {
	return "product"
}

func (e *AuctionProductExtension) Priority() int {
	return 100
}

func (e *AuctionProductExtension) GetProductTypes() []string {
	return []string{"auction"}
}

func (e *AuctionProductExtension) CalculatePrice(ctx context.Context, product interface{}, identity *identity.BusinessIdentity) (float64, error) {
	return 0, nil
}

func InitAuction(registry *spi.ExtensionRegistry) {
	id := identity.NewBusinessIdentity(identity.CountryCN, identity.ModeAuction)
	registry.Register(id, &AuctionProductExtension{})
	registry.Register(id, &AuctionTradeExtension{})
}
