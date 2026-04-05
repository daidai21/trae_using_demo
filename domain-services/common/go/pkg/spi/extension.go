package spi

import (
	"context"
	"ecommerce/common/pkg/identity"
)

type ExtensionPoint interface {
	Name() string
	Priority() int
}

type ProductExtension interface {
	ExtensionPoint
	GetProductTypes() []string
	CalculatePrice(ctx context.Context, product interface{}, identity *identity.BusinessIdentity) (float64, error)
}

type TradeExtension interface {
	ExtensionPoint
	BeforeCreateOrder(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error
	ProcessPayment(ctx context.Context, order interface{}, identity *identity.BusinessIdentity) error
}

type I18nExtension interface {
	ExtensionPoint
	GetLocales() map[string]map[string]string
	FormatCurrency(amount float64, currency string) string
}

type PaymentExtension interface {
	ExtensionPoint
	GetPaymentGateways() []string
	InitiatePayment(ctx context.Context, order interface{}, gateway string) (string, error)
}
