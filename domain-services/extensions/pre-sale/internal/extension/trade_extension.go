package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type PreSaleTradeExtension struct{}

func (e *PreSaleTradeExtension) Name() string {
	return "trade"
}

func (e *PreSaleTradeExtension) Priority() int {
	return 100
}

func (e *PreSaleTradeExtension) BeforeCreateOrder(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error {
	return nil
}

func (e *PreSaleTradeExtension) ProcessPayment(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error {
	return nil
}
