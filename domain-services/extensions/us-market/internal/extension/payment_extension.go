package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type USMarketPaymentExtension struct{}

func (e *USMarketPaymentExtension) Name() string {
	return "payment"
}

func (e *USMarketPaymentExtension) Priority() int {
	return 100
}

func (e *USMarketPaymentExtension) GetPaymentGateways() []string {
	return []string{"stripe", "paypal", "ach"}
}

func (e *USMarketPaymentExtension) InitiatePayment(ctx context.Context, order interface{}, gateway string) (string, error) {
	return "https://payment.stripe.com/us/" + gateway, nil
}
