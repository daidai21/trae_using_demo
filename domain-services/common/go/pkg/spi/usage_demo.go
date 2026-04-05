package spi

import (
	"context"
	"fmt"
	"log"
	"ecommerce/common/pkg/identity"
)

func CompleteUsageDemo() {
	log.Println("========================================")
	log.Println("  平台化架构 v4.0 - SPI 完整使用示例")
	log.Println("========================================")

	ctx := context.Background()

	registry := NewExtensionRegistry()
	loader := NewExtensionLoader(registry)

	log.Println("\n--- 步骤 1: 初始化并注册所有扩展 ---")
	loader.LoadExtensions(
		RegisterDemoExtensions,
	)

	log.Println("\n--- 步骤 2: 为不同业务身份调用扩展 ---")

	demoCases := []struct {
		name         string
		identity     *identity.BusinessIdentity
		productPrice float64
	}{
		{
			name:         "印尼普通交易 (ID.normal)",
			identity:     identity.NewBusinessIdentity(identity.CountryID, identity.ModeNormal),
			productPrice: 150000,
		},
		{
			name:         "美国普通交易 (US.normal)",
			identity:     identity.NewBusinessIdentity(identity.CountryUS, identity.ModeNormal),
			productPrice: 99.99,
		},
		{
			name:         "中国预售 (CN.preSale)",
			identity:     identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale),
			productPrice: 199.00,
		},
		{
			name:         "中国拍卖 (CN.auction)",
			identity:     identity.NewBusinessIdentity(identity.CountryCN, identity.ModeAuction),
			productPrice: 299.00,
		},
		{
			name:         "中国普通交易 (CN.normal - 无扩展)",
			identity:     identity.NewBusinessIdentity(identity.CountryCN, identity.ModeNormal),
			productPrice: 100.00,
		},
	}

	for _, demo := range demoCases {
		log.Printf("\n--- %s ---", demo.name)
		demonstrateExtensionUsage(ctx, loader, demo.identity, demo.productPrice)
	}

	log.Println("\n========================================")
	log.Println("  示例运行完成")
	log.Println("========================================")
}

func demonstrateExtensionUsage(ctx context.Context, loader *ExtensionLoader, id *identity.BusinessIdentity, productPrice float64) {
	log.Printf("业务身份: %s", id.String())

	if i18nExt, ok := loader.GetI18nExtension(id); ok {
		log.Printf("✓ I18n 扩展已加载: %s", i18nExt.Name())

		formatted := i18nExt.FormatCurrency(productPrice, "LOCAL")
		log.Printf("  货币格式化: %s", formatted)

		for lang, messages := range i18nExt.GetLocales() {
			log.Printf("  语言 %s: %s", lang, messages["welcome"])
		}
	} else {
		log.Printf("✗ 未找到 I18n 扩展")
	}

	if productExt, ok := loader.GetProductExtension(id); ok {
		log.Printf("✓ Product 扩展已加载: %s", productExt.Name())

		types := productExt.GetProductTypes()
		log.Printf("  支持商品类型: %v", types)

		calculatedPrice, _ := productExt.CalculatePrice(ctx, nil, id)
		log.Printf("  CalculatePrice 调用结果: %.2f", calculatedPrice)
	} else {
		log.Printf("✗ 未找到 Product 扩展，使用默认实现")
		log.Printf("  支持商品类型: [normal]")
		log.Printf("  默认价格: %.2f", 100.0)
	}
}

func RegisterDemoExtensions(registry *ExtensionRegistry) {
	log.Println("正在注册演示扩展...")

	idNormal := identity.NewBusinessIdentity(identity.CountryID, identity.ModeNormal)
	idI18nExt := &DemoI18nExtension{
		country: identity.CountryID,
		locales: map[string]map[string]string{
			"id-ID": {
				"welcome":        "Selamat datang di platform kami",
				"product":        "Produk",
				"cart":           "Keranjang",
				"checkout":       "Checkout",
			},
		},
	}
	registry.Register(idNormal, idI18nExt)
	log.Printf("  ✓ %s -> %s", idNormal.String(), idI18nExt.Name())

	usNormal := identity.NewBusinessIdentity(identity.CountryUS, identity.ModeNormal)
	usI18nExt := &DemoI18nExtension{
		country: identity.CountryUS,
		locales: map[string]map[string]string{
			"en-US": {
				"welcome":        "Welcome to our platform",
				"product":        "Product",
				"cart":           "Cart",
				"checkout":       "Checkout",
			},
		},
	}
	registry.Register(usNormal, usI18nExt)
	log.Printf("  ✓ %s -> %s", usNormal.String(), usI18nExt.Name())

	cnPreSale := identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)
	preSaleExt := &DemoProductExtension{
		mode:         identity.ModePreSale,
		productTypes: []string{"pre_sale"},
		basePrice:    99.0,
	}
	registry.Register(cnPreSale, preSaleExt)
	log.Printf("  ✓ %s -> %s", cnPreSale.String(), preSaleExt.Name())

	cnAuction := identity.NewBusinessIdentity(identity.CountryCN, identity.ModeAuction)
	auctionExt := &DemoProductExtension{
		mode:         identity.ModeAuction,
		productTypes: []string{"auction"},
		basePrice:    199.0,
	}
	registry.Register(cnAuction, auctionExt)
	log.Printf("  ✓ %s -> %s", cnAuction.String(), auctionExt.Name())
}

type DemoI18nExtension struct {
	country string
	locales map[string]map[string]string
}

func (e *DemoI18nExtension) Name() string {
	return "i18n"
}

func (e *DemoI18nExtension) Priority() int {
	return 100
}

func (e *DemoI18nExtension) GetLocales() map[string]map[string]string {
	return e.locales
}

func (e *DemoI18nExtension) FormatCurrency(amount float64, currency string) string {
	if e.country == identity.CountryID {
		return "Rp " + fmt.Sprintf("%.0f", amount)
	} else if e.country == identity.CountryUS {
		return "$" + fmt.Sprintf("%.2f", amount)
	}
	return "¥" + fmt.Sprintf("%.2f", amount)
}

type DemoProductExtension struct {
	mode         string
	productTypes []string
	basePrice    float64
}

func (e *DemoProductExtension) Name() string {
	return "product"
}

func (e *DemoProductExtension) Priority() int {
	return 100
}

func (e *DemoProductExtension) GetProductTypes() []string {
	return e.productTypes
}

func (e *DemoProductExtension) CalculatePrice(ctx context.Context, product interface{}, id *identity.BusinessIdentity) (float64, error) {
	log.Printf("  [CalculatePrice] 被调用，业务身份: %s, 模式: %s", id.String(), e.mode)

	if e.mode == identity.ModePreSale {
		deposit := e.basePrice * 0.3
		log.Printf("  [CalculatePrice] 预售模式: 定金 30%% = %.2f", deposit)
		return deposit, nil
	} else if e.mode == identity.ModeAuction {
		startPrice := e.basePrice
		log.Printf("  [CalculatePrice] 拍卖模式: 起拍价 = %.2f", startPrice)
		return startPrice, nil
	}

	return e.basePrice, nil
}
