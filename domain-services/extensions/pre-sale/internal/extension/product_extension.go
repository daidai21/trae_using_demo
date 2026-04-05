package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type PreSaleProductExtension struct{}

func (e *PreSaleProductExtension) Name() string {
	return "product"
}

func (e *PreSaleProductExtension) Priority() int {
	return 100
}

func (e *PreSaleProductExtension) GetProductTypes() []string {
	return []string{"pre_sale"}
}

func (e *PreSaleProductExtension) CalculatePrice(ctx context.Context, product interface{}, identity *identity.BusinessIdentity) (float64, error) {
	return 0, nil
}

func InitPreSale(registry *spi.ExtensionRegistry) {
	id := identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)
	registry.Register(id, &PreSaleProductExtension{})
	registry.Register(id, &PreSaleTradeExtension{})
}
