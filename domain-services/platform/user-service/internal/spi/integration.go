package spi

import (
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
		log.Println("正在初始化 User Service SPI 框架...")

		registry = spi.NewExtensionRegistry()
		loader = spi.NewExtensionLoader(registry)

		loader.LoadExtensions(
			InitPlatformExtensions,
		)

		log.Println("User Service SPI 框架初始化完成")
	})
}

func InitPlatformExtensions(registry *spi.ExtensionRegistry) {
	log.Println("正在注册 User Service 扩展...")
}

func GetSPIExtensionRegistry() *spi.ExtensionRegistry {
	return registry
}

func GetSPIExtensionLoader() *spi.ExtensionLoader {
	return loader
}

func GetUserPreferredIdentity(userID uint) *identity.BusinessIdentity {
	return identity.NewBusinessIdentity(identity.CountryCN, identity.ModeNormal)
}
