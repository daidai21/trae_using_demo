package extension

import (
	"context"
	"fmt"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

type IDMarketI18nExtension struct{}

func (e *IDMarketI18nExtension) Name() string {
	return "i18n"
}

func (e *IDMarketI18nExtension) Priority() int {
	return 100
}

func (e *IDMarketI18nExtension) GetLocales() map[string]map[string]string {
	return map[string]map[string]string{
		"id-ID": {
			"welcome": "Selamat datang",
		},
	}
}

func (e *IDMarketI18nExtension) FormatCurrency(amount float64, currency string) string {
	return "Rp " + formatIDR(amount)
}

func formatIDR(amount float64) string {
	return fmt.Sprintf("%.0f", amount)
}

func InitIDMarket(registry *spi.ExtensionRegistry) {
	id := identity.NewBusinessIdentity(identity.CountryID, identity.ModeNormal)
	registry.Register(id, &IDMarketI18nExtension{})
	registry.Register(id, &IDMarketPaymentExtension{})
}
