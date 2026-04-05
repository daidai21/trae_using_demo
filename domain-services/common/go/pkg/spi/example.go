package spi

import (
	"context"
	"fmt"
	"log"
	"ecommerce/common/pkg/identity"
)

func ExampleSPIUsage() {
	registry := NewExtensionRegistry()
	loader := NewExtensionLoader(registry)

	log.Println("=== 平台化架构 v4.0 - SPI 初始化示例 ===")

	loader.LoadExtensions(
		InitAllExtensions,
	)

	log.Println("=== SPI 扩展使用示例 ===")

	examples := []struct {
		name     string
		identity *identity.BusinessIdentity
	}{
		{"印尼普通交易", identity.NewBusinessIdentity(identity.CountryID, identity.ModeNormal)},
		{"美国普通交易", identity.NewBusinessIdentity(identity.CountryUS, identity.ModeNormal)},
		{"中国预售", identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)},
		{"中国拍卖", identity.NewBusinessIdentity(identity.CountryCN, identity.ModeAuction)},
	}

	for _, ex := range examples {
		log.Printf("\n--- %s (%s) ---", ex.name, ex.identity.String())
		useExtensionsForIdentity(loader, ex.identity)
	}
}

func useExtensionsForIdentity(loader *ExtensionLoader, id *identity.BusinessIdentity) {
	if i18nExt, ok := loader.GetI18nExtension(id); ok {
		log.Printf("I18n 扩展已加载: %s (优先级: %d)", i18nExt.Name(), i18nExt.Priority())

		locales := i18nExt.GetLocales()
		for lang, messages := range locales {
			log.Printf("  语言包: %s, 欢迎消息: %s", lang, messages["welcome"])
		}

		formatted := i18nExt.FormatCurrency(100000, "LOCAL")
		log.Printf("  货币格式化: %s", formatted)
	}

	if productExt, ok := loader.GetProductExtension(id); ok {
		log.Printf("Product 扩展已加载: %s (优先级: %d)", productExt.Name(), productExt.Priority())

		types := productExt.GetProductTypes()
		log.Printf("  商品类型: %v", types)

		price, _ := productExt.CalculatePrice(context.Background(), nil, id)
		log.Printf("  价格计算: %.2f", price)
	}
}

func InitAllExtensions(registry *ExtensionRegistry) {
	log.Println("正在注册所有扩展...")

	id := identity.NewBusinessIdentity(identity.CountryID, identity.ModeNormal)
	i18nExt := &ExampleI18nExtension{
		country: identity.CountryID,
		locales: map[string]map[string]string{
			"id-ID": {"welcome": "Selamat datang"},
		},
	}
	registry.Register(id, i18nExt)
	log.Printf("  已注册: %s -> %s", id.String(), i18nExt.Name())

	us := identity.NewBusinessIdentity(identity.CountryUS, identity.ModeNormal)
	usI18nExt := &ExampleI18nExtension{
		country: identity.CountryUS,
		locales: map[string]map[string]string{
			"en-US": {"welcome": "Welcome"},
		},
	}
	registry.Register(us, usI18nExt)
	log.Printf("  已注册: %s -> %s", us.String(), usI18nExt.Name())

	preSale := identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)
	preSaleProductExt := &ExampleProductExtension{
		mode:         identity.ModePreSale,
		productTypes: []string{"pre_sale"},
	}
	registry.Register(preSale, preSaleProductExt)
	log.Printf("  已注册: %s -> %s", preSale.String(), preSaleProductExt.Name())

	auction := identity.NewBusinessIdentity(identity.CountryCN, identity.ModeAuction)
	auctionProductExt := &ExampleProductExtension{
		mode:         identity.ModeAuction,
		productTypes: []string{"auction"},
	}
	registry.Register(auction, auctionProductExt)
	log.Printf("  已注册: %s -> %s", auction.String(), auctionProductExt.Name())
}

type ExampleI18nExtension struct {
	country string
	locales map[string]map[string]string
}

func (e *ExampleI18nExtension) Name() string {
	return "i18n"
}

func (e *ExampleI18nExtension) Priority() int {
	return 100
}

func (e *ExampleI18nExtension) GetLocales() map[string]map[string]string {
	return e.locales
}

func (e *ExampleI18nExtension) FormatCurrency(amount float64, currency string) string {
	if e.country == identity.CountryID {
		return "Rp " + formatIDR(amount)
	} else if e.country == identity.CountryUS {
		return "$" + formatUSD(amount)
	}
	return "¥" + formatCNY(amount)
}

func formatIDR(amount float64) string {
	return fmt.Sprintf("%.0f", amount)
}

func formatUSD(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}

func formatCNY(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}

type ExampleProductExtension struct {
	mode         string
	productTypes []string
}

func (e *ExampleProductExtension) Name() string {
	return "product"
}

func (e *ExampleProductExtension) Priority() int {
	return 100
}

func (e *ExampleProductExtension) GetProductTypes() []string {
	return e.productTypes
}

func (e *ExampleProductExtension) CalculatePrice(ctx context.Context, product interface{}, id *identity.BusinessIdentity) (float64, error) {
	if e.mode == identity.ModePreSale {
		return 99.0, nil
	} else if e.mode == identity.ModeAuction {
		return 199.0, nil
	}
	return 100.0, nil
}
