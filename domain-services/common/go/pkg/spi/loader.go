package spi

import (
	"ecommerce/common/pkg/identity"
	"log"
)

type ExtensionLoader struct {
	registry *ExtensionRegistry
}

func NewExtensionLoader(registry *ExtensionRegistry) *ExtensionLoader {
	return &ExtensionLoader{
		registry: registry,
	}
}

func (l *ExtensionLoader) LoadExtensions(initializers ...func(*ExtensionRegistry)) {
	for _, init := range initializers {
		init(l.registry)
	}
	log.Println("All extensions loaded successfully")
}

func (l *ExtensionLoader) GetExtensionsForIdentity(id *identity.BusinessIdentity) []ExtensionPoint {
	return l.registry.GetExtensions(id)
}

func (l *ExtensionLoader) GetProductExtension(id *identity.BusinessIdentity) (ProductExtension, bool) {
	return l.registry.GetProductExtension(id)
}

func (l *ExtensionLoader) GetTradeExtension(id *identity.BusinessIdentity) (TradeExtension, bool) {
	return l.registry.GetTradeExtension(id)
}

func (l *ExtensionLoader) GetI18nExtension(id *identity.BusinessIdentity) (I18nExtension, bool) {
	return l.registry.GetI18nExtension(id)
}

func (l *ExtensionLoader) GetPaymentExtension(id *identity.BusinessIdentity) (PaymentExtension, bool) {
	return l.registry.GetPaymentExtension(id)
}
