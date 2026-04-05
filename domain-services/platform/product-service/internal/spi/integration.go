package spi

import (
	"context"
	"fmt"
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
		log.Println("正在初始化 Product Service SPI 框架...")

		registry = spi.NewExtensionRegistry()
		loader = spi.NewExtensionLoader(registry)

		loader.LoadExtensions(
			InitPlatformExtensions,
		)

		log.Println("Product Service SPI 框架初始化完成")
	})
}

func InitPlatformExtensions(registry *spi.ExtensionRegistry) {
	log.Println("正在注册 Product Service 扩展...")
}

func GetSPIExtensionRegistry() *spi.ExtensionRegistry {
	return registry
}

func GetSPIExtensionLoader() *spi.ExtensionLoader {
	return loader
}

func CalculateProductPriceWithExtension(ctx context.Context, product interface{}, id *identity.BusinessIdentity) (float64, error) {
	if loader == nil {
		InitializeSPI()
	}

	if productExt, ok := loader.GetProductExtension(id); ok {
		log.Printf("使用扩展 CalculatePrice: %s (业务身份: %s)", productExt.Name(), id.String())
		return productExt.CalculatePrice(ctx, product, id)
	}

	log.Printf("未找到扩展，使用默认价格计算 (业务身份: %s)", id.String())
	return calculateDefaultPrice(product), nil
}

func calculateDefaultPrice(product interface{}) float64 {
	return 100.0
}

func GetProductTypesForIdentity(id *identity.BusinessIdentity) []string {
	if loader == nil {
		InitializeSPI()
	}

	if productExt, ok := loader.GetProductExtension(id); ok {
		return productExt.GetProductTypes()
	}

	return []string{"normal"}
}

func FormatCurrencyForIdentity(amount float64, currency string, id *identity.BusinessIdentity) string {
	if loader == nil {
		InitializeSPI()
	}

	if i18nExt, ok := loader.GetI18nExtension(id); ok {
		return i18nExt.FormatCurrency(amount, currency)
	}

	return formatDefaultCurrency(amount, currency)
}

func formatDefaultCurrency(amount float64, currency string) string {
	return "¥" + fmt.Sprintf("%.2f", amount)
}
