# E-Commerce Platform v4.0

A complete e-commerce trading platform with microservices architecture and frontend Monorepo, now upgraded to Platform v4.0 with SPI extension framework.

## Tech Stack

### Backend
- **Language**: Golang 1.21+
- **Web Framework**: Hertz
- **ORM**: GORM
- **Database**: SQLite
- **Authentication**: JWT
- **Architecture**: Microservices + SPI Extension Framework

### Frontend
- **Framework**: React 18
- **UI Component Library**: Ant Design 5
- **Routing**: React Router v6
- **HTTP Client**: Axios
- **Build Tool**: Vite
- **Language**: TypeScript
- **Package Manager**: pnpm
- **Architecture**: Monorepo

## Project Structure

```
trae_using_demo/
├── domain-services/
│   ├── api-gateway/              # API Gateway (Port 8080)
│   │   ├── cmd/server/
│   │   ├── internal/proxy/
│   │   └── go.mod
│   ├── common/                   # Common packages
│   │   └── go/pkg/
│   │       ├── identity/         # Business Identity abstraction
│   │       ├── spi/              # SPI Extension Framework
│   │       ├── response/
│   │       └── utils/
│   ├── platform/                 # Platform Core Services (v4.0)
│   │   ├── user-service/         # Port 9081 (planned)
│   │   ├── product-service/      # Port 9082 (planned)
│   │   └── trade-service/        # Port 9083 (planned)
│   └── extensions/               # Business Extensions
│       ├── id-market/            # Indonesia Market
│       ├── us-market/            # US Market
│       ├── pre-sale/             # Pre-Sale Feature
│       └── auction/              # Live Auction Feature
├── frontend/                     # Frontend Monorepo
│   ├── apps/web/                 # Main Web App
│   ├── packages/
│   │   ├── ui/                   # Shared UI Components
│   │   ├── user/                 # User Domain
│   │   ├── product/              # Product Domain
│   │   ├── trade/                # Trade Domain
│   │   └── auction/              # Auction Domain
│   └── pnpm-workspace.yaml
├── .trae/
│   ├── documents/                # Architecture plans
│   └── specs/                    # Specifications
└── README.md
```

## Architecture v4.0 - Platform with SPI

### Core Concept: Business Identity
```
{Country}.{Mode}

Examples:
CN.normal  - China Normal Trading
ID.normal  - Indonesia Normal Trading
US.normal  - US Normal Trading
CN.preSale - China Pre-Sale
CN.auction - China Live Auction
```

### SPI Extension Framework

#### Extension Points
- **ProductExtension**: Product type extension, price calculation
- **TradeExtension**: Order creation, payment processing
- **I18nExtension**: Multi-language, currency formatting
- **PaymentExtension**: Payment gateway integration

#### Directory Structure
```
common/go/pkg/
├── identity/         # Business Identity
│   ├── identity.go   # BusinessIdentity struct
│   └── resolver.go   # Identity resolver
└── spi/
    ├── extension.go  # Extension point interfaces
    ├── registry.go   # Extension registry
    ├── loader.go     # Extension loader
    └── usage_demo.go # Complete usage example
```

## Features

### User Authentication
- User registration & login
- JWT Token authentication (24-hour validity)
- Password bcrypt encryption

### Merchant Management
- Merchant registration & editing
- Merchant list & detail view

### Product Management
- Product publishing & editing
- Product list & details
- Inventory management
- Multi-currency support (via SPI)

### Shopping Cart
- Add/modify/delete cart items
- Cart list view

### Order Management
- Create orders
- Order list & details
- Order status management
- Automatic inventory deduction

### Business Growth Features (via SPI Extensions)
- **ID Market**: Bahasa Indonesia, IDR currency, Midtrans/Doku payment
- **US Market**: English (US), USD currency, Stripe/PayPal payment, Sales Tax
- **Pre-Sale**: Deposit + balance payment mode
- **Live Auction**: Real-time bidding, WebSocket support

## Quick Start

### Backend Startup

#### Start API Gateway
```bash
cd domain-services/api-gateway
go run cmd/server/main.go
```
API Gateway starts at `http://localhost:8080`

#### Start Microservices (Legacy, Ports 8081-8084)
```bash
# User Service (Port 8081)
cd domain-services/user-service
go run cmd/server/main.go

# Product Service (Port 8082)
cd domain-services/product-service
go run cmd/server/main.go

# Trade Service (Port 8083)
cd domain-services/trade-service
go run cmd/server/main.go

# Auction Service (Port 8084)
cd domain-services/auction-service
go run cmd/server/main.go
```

### Frontend Startup

```bash
cd frontend
pnpm install
pnpm dev
```
Frontend dev server starts at `http://localhost:5173`

## SPI Framework Usage Example

```go
import (
    "ecommerce/common/pkg/identity"
    "ecommerce/common/pkg/spi"
)

// 1. Initialize SPI
registry := spi.NewExtensionRegistry()
loader := spi.NewExtensionLoader(registry)

// 2. Load extensions
loader.LoadExtensions(
    id_market.InitIDMarket,
    us_market.InitUSMarket,
    pre_sale.InitPreSale,
    auction.InitAuction,
)

// 3. Use extension for specific business identity
id := identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)
if productExt, ok := loader.GetProductExtension(id); ok {
    price, _ := productExt.CalculatePrice(ctx, product, id)
}
```

## API Documentation

### Authentication Endpoints
- `POST /api/auth/register` - User registration
- `POST /api/auth/login` - User login

### Product Endpoints (Public)
- `GET /api/products` - Get product list
- `GET /api/products/:id` - Get product details

### Authenticated Endpoints

#### Merchant Management
- `POST /api/merchants` - Create merchant
- `GET /api/merchants` - Get merchant list
- `GET /api/merchants/:id` - Get merchant details
- `PUT /api/merchants/:id` - Update merchant

#### Product Management
- `POST /api/products` - Create product
- `PUT /api/products/:id` - Update product
- `DELETE /api/products/:id` - Delete product

#### Shopping Cart
- `POST /api/cart` - Add to cart
- `GET /api/cart` - Get cart
- `PUT /api/cart/:id` - Update cart item
- `DELETE /api/cart/:id` - Delete cart item

#### Order Management
- `POST /api/orders` - Create order
- `GET /api/orders` - Get order list
- `GET /api/orders/:id` - Get order details
- `PUT /api/orders/:id/status` - Update order status

#### Auction Endpoints
- `GET /api/auctions` - Get auction list
- `GET /api/auctions/:id` - Get auction details
- `POST /api/auctions` - Create auction
- `POST /api/auctions/:id/bid` - Place bid
- `WebSocket /api/ws/auctions/:id` - Real-time bidding

## Database Schema

### Core Tables
- **users**: User accounts
- **merchants**: Merchant profiles
- **products**: Product information
- **product_prices**: Multi-currency prices
- **carts**: Shopping cart
- **orders**: Orders
- **order_items**: Order items

### Extension Tables (Optional)
- **auctions**: Auction information
- **bids**: Bid records

## Development Notes

### Platform v4.0 Migration
- Legacy services continue on ports 8081-8084
- New platform services planned for ports 9081-9083
- SPI extensions enable business feature isolation
- Smooth migration with backward compatibility

### SPI Extension Development
1. Implement extension point interface
2. Provide `Init*()` registration function
3. Register extension with business identity
4. Load extension in platform service

## License

MIT License
