package extension

import (
	"context"
	"fmt"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type USMarketI18nExtension struct{}

func (e *USMarketI18nExtension) Name() string {
	return "i18n"
}

func (e *USMarketI18nExtension) Priority() int {
	return 100
}

func (e *USMarketI18nExtension) GetLocales() map[string]map[string]string {
	return map[string]map[string]string{
		"en-US": {
			"welcome": "Welcome",
		},
	}
}

func (e *USMarketI18nExtension) FormatCurrency(amount float64, currency string) string {
	return "$" + fmt.Sprintf("%.2f", amount)
}

func InitUSMarket(registry *spi.ExtensionRegistry) {
	id := identity.NewBusinessIdentity(identity.CountryUS, identity.ModeNormal)
	registry.Register(id, &USMarketI18nExtension{})
	registry.Register(id, &USMarketPaymentExtension{})
	registry.Register(id, &USMarketTaxExtension{})
}
