module ecommerce/product-service

go 1.21

require (
	ecommerce/common v0.0.0
	github.com/cloudwego/hertz v0.9.0
	github.com/golang-jwt/jwt/v5 v5.2.1
	golang.org/x/crypto v0.28.0
	gorm.io/driver/sqlite v1.5.7
	gorm.io/gorm v1.25.12
)

replace ecommerce/common => ../common/go
