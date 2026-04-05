package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type AuctionTradeExtension struct{}

func (e *AuctionTradeExtension) Name() string {
	return "trade"
}

func (e *AuctionTradeExtension) Priority() int {
	return 100
}

func (e *AuctionTradeExtension) BeforeCreateOrder(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error {
	return nil
}

func (e *AuctionTradeExtension) ProcessPayment(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error {
	return nil
}
