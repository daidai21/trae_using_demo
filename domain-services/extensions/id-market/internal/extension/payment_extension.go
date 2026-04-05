package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type IDMarketPaymentExtension struct{}

func (e *IDMarketPaymentExtension) Name() string {
	return "payment"
}

func (e *IDMarketPaymentExtension) Priority() int {
	return 100
}

func (e *IDMarketPaymentExtension) GetPaymentGateways() []string {
	return []string{"midtrans", "doku", "gopay"}
}

func (e *IDMarketPaymentExtension) InitiatePayment(ctx context.Context, order interface{}, gateway string) (string, error) {
	return "https://payment.midtrans.com/id/" + gateway, nil
}
