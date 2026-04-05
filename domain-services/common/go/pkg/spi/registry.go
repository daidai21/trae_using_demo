package spi

import (
	"ecommerce/common/pkg/identity"
	"sort"
	"sync"
)

type ExtensionRegistry struct {
	extensions map[string]map[string]ExtensionPoint
	mu         sync.RWMutex
}

func NewExtensionRegistry() *ExtensionRegistry {
	return &ExtensionRegistry{
		extensions: make(map[string]map[string]ExtensionPoint),
	}
}

func (r *ExtensionRegistry) Register(id *identity.BusinessIdentity, ext ExtensionPoint) {
	r.mu.Lock()
	defer r.mu.Unlock()

	key := id.String()
	if r.extensions[key] == nil {
		r.extensions[key] = make(map[string]ExtensionPoint)
	}
	r.extensions[key][ext.Name()] = ext
}

func (r *ExtensionRegistry) GetExtensions(id *identity.BusinessIdentity) []ExtensionPoint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := id.String()
	exts, ok := r.extensions[key]
	if !ok {
		return nil
	}

	result := make([]ExtensionPoint, 0, len(exts))
	for _, ext := range exts {
		result = append(result, ext)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Priority() > result[j].Priority()
	})

	return result
}

func (r *ExtensionRegistry) GetExtension(id *identity.BusinessIdentity, name string) (ExtensionPoint, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := id.String()
	exts, ok := r.extensions[key]
	if !ok {
		return nil, false
	}

	ext, ok := exts[name]
	return ext, ok
}

func (r *ExtensionRegistry) GetProductExtension(id *identity.BusinessIdentity) (ProductExtension, bool) {
	ext, ok := r.GetExtension(id, "product")
	if !ok {
		return nil, false
	}
	productExt, ok := ext.(ProductExtension)
	return productExt, ok
}

func (r *ExtensionRegistry) GetTradeExtension(id *identity.BusinessIdentity) (TradeExtension, bool) {
	ext, ok := r.GetExtension(id, "trade")
	if !ok {
		return nil, false
	}
	tradeExt, ok := ext.(TradeExtension)
	return tradeExt, ok
}

func (r *ExtensionRegistry) GetI18nExtension(id *identity.BusinessIdentity) (I18nExtension, bool) {
	ext, ok := r.GetExtension(id, "i18n")
	if !ok {
		return nil, false
	}
	i18nExt, ok := ext.(I18nExtension)
	return i18nExt, ok
}

func (r *ExtensionRegistry) GetPaymentExtension(id *identity.BusinessIdentity) (PaymentExtension, bool) {
	ext, ok := r.GetExtension(id, "payment")
	if !ok {
		return nil, false
	}
	paymentExt, ok := ext.(PaymentExtension)
	return paymentExt, ok
}
