package extension

import (
	"context"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type USMarketTaxExtension struct{}

func (e *USMarketTaxExtension) Name() string {
	return "tax"
}

func (e *USMarketTaxExtension) Priority() int {
	return 100
}

func (e *USMarketTaxExtension) CalculateSalesTax(ctx context.Context, amount float64, state string) float64 {
	taxRates := map[string]float64{
		"CA": 0.0725,
		"NY": 0.08875,
		"TX": 0.0625,
	}
	if rate, ok := taxRates[state]; ok {
		return amount * rate
	}
	return amount * 0.06
}
