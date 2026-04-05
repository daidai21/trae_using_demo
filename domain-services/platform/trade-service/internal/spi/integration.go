package spi

import (
	"context"
	"log"
	"sync"
	"ecommerce/common/pkg/identity"
	"ecommerce/common/pkg/spi"
)

var (
	registry   *spi.ExtensionRegistry
	loader     *spi.ExtensionLoader
	initOnce   sync.Once
)

func InitializeSPI() {
	initOnce.Do(func() {
		log.Println("正在初始化 Trade Service SPI 框架...")

		registry = spi.NewExtensionRegistry()
		loader = spi.NewExtensionLoader(registry)

		loader.LoadExtensions(
			InitPlatformExtensions,
		)

		log.Println("Trade Service SPI 框架初始化完成")
	})
}

func InitPlatformExtensions(registry *spi.ExtensionRegistry) {
	log.Println("正在注册 Trade Service 扩展...")
}

func GetSPIExtensionRegistry() *spi.ExtensionRegistry {
	return registry
}

func GetSPIExtensionLoader() *spi.ExtensionLoader {
	return loader
}

func BeforeCreateOrderWithExtension(ctx context.Context, order interface{}, id *identity.BusinessIdentity) error {
	if loader == nil {
		InitializeSPI()
	}

	if tradeExt, ok := loader.GetTradeExtension(id); ok {
		log.Printf("使用扩展 BeforeCreateOrder: %s (业务身份: %s)", tradeExt.Name(), id.String())
		return tradeExt.BeforeCreateOrder(ctx, order, id)
	}

	return nil
}

func ProcessPaymentWithExtension(ctx context.Context, order interface{}, id *identity.BusinessIdentity) error {
	if loader == nil {
		InitializeSPI()
	}

	if tradeExt, ok := loader.GetTradeExtension(id); ok {
		log.Printf("使用扩展 ProcessPayment: %s (业务身份: %s)", tradeExt.Name(), id.String())
		return tradeExt.ProcessPayment(ctx, order, id)
	}

	return nil
}

func GetPaymentGatewaysForIdentity(id *identity.BusinessIdentity) []string {
	if loader == nil {
		InitializeSPI()
	}

	if paymentExt, ok := loader.GetPaymentExtension(id); ok {
		return paymentExt.GetPaymentGateways()
	}

	return []string{"default"}
}
