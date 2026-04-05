# 电商交易平台 v4.0

一个完整的电商交易平台，采用微服务架构和前端 Monorepo，现已升级为 Platform v4.0，集成了 SPI 扩展框架。

## 技术栈

### 后端
- **语言**: Golang 1.21+
- **Web框架**: Hertz
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT
- **架构**: 微服务 + SPI 扩展框架

### 前端
- **框架**: React 18
- **UI组件库**: Ant Design 5
- **路由**: React Router v6
- **HTTP客户端**: Axios
- **构建工具**: Vite
- **语言**: TypeScript
- **包管理**: pnpm
- **架构**: Monorepo

## 项目结构

```
trae_using_demo/
├── domain-services/
│   ├── api-gateway/              # API 网关（端口 8080）
│   │   ├── cmd/server/
│   │   ├── internal/proxy/
│   │   └── go.mod
│   ├── common/                   # 公共包
│   │   └── go/pkg/
│   │       ├── identity/         # 业务身份抽象
│   │       ├── spi/              # SPI 扩展框架
│   │       ├── response/
│   │       └── utils/
│   ├── platform/                 # 平台核心服务（v4.0）
│   │   ├── user-service/         # 端口 9081（计划中）
│   │   ├── product-service/      # 端口 9082（计划中）
│   │   └── trade-service/        # 端口 9083（计划中）
│   └── extensions/               # 业务扩展
│       ├── id-market/            # 印尼市场
│       ├── us-market/            # 美国市场
│       ├── pre-sale/             # 预售功能
│       └── auction/              # 直播拍卖
├── frontend/                     # 前端 Monorepo
│   ├── apps/web/                 # 主 Web 应用
│   ├── packages/
│   │   ├── ui/                   # 共享 UI 组件
│   │   ├── user/                 # 用户域
│   │   ├── product/              # 商品域
│   │   ├── trade/                # 交易域
│   │   └── auction/              # 拍卖域
│   └── pnpm-workspace.yaml
├── .trae/
│   ├── documents/                # 架构计划
│   └── specs/                    # 规范文档
└── README.md
```

## 架构 v4.0 - 平台化与 SPI

### 核心概念：业务身份
```
{国家}.{玩法}

示例：
CN.normal  - 中国普通交易
ID.normal  - 印尼普通交易
US.normal  - 美国普通交易
CN.preSale - 中国预售
CN.auction - 中国直播拍卖
```

### SPI 扩展框架

#### 扩展点
- **ProductExtension**: 商品类型扩展、价格计算
- **TradeExtension**: 订单创建、支付处理
- **I18nExtension**: 多语言、货币格式化
- **PaymentExtension**: 支付网关集成

#### 目录结构
```
common/go/pkg/
├── identity/         # 业务身份
│   ├── identity.go   # BusinessIdentity 结构体
│   └── resolver.go   # 身份解析器
└── spi/
    ├── extension.go  # 扩展点接口
    ├── registry.go   # 扩展注册表
    ├── loader.go     # 扩展加载器
    └── usage_demo.go # 完整使用示例
```

## 功能特性

### 用户认证
- 用户注册与登录
- JWT Token 认证（24小时有效期）
- 密码 bcrypt 加密存储

### 商家管理
- 商家入驻与信息编辑
- 商家列表与详情查看

### 商品管理
- 商品发布与编辑
- 商品列表与详情
- 库存管理
- 多货币支持（通过 SPI）

### 购物车
- 添加/修改/删除购物车商品
- 购物车列表查看

### 订单管理
- 创建订单
- 订单列表与详情
- 订单状态管理
- 库存自动扣减

### 业务增长功能（通过 SPI 扩展实现）
- **印尼市场**: 印尼语、IDR 货币、Midtrans/Doku 支付
- **美国市场**: 美式英语、USD 货币、Stripe/PayPal 支付、销售税
- **预售功能**: 定金+尾款支付模式
- **直播拍卖**: 实时出价、WebSocket 支持

## 快速开始

### 后端启动

#### 启动 API 网关
```bash
cd domain-services/api-gateway
go run cmd/server/main.go
```
API 网关将在 `http://localhost:8080` 启动

#### 启动微服务（旧版，端口 8081-8084）
```bash
# 用户服务（端口 8081）
cd domain-services/user-service
go run cmd/server/main.go

# 商品服务（端口 8082）
cd domain-services/product-service
go run cmd/server/main.go

# 交易服务（端口 8083）
cd domain-services/trade-service
go run cmd/server/main.go

# 拍卖服务（端口 8084）
cd domain-services/auction-service
go run cmd/server/main.go
```

### 前端启动

```bash
cd frontend
pnpm install
pnpm dev
```
前端开发服务器将在 `http://localhost:5173` 启动

## SPI 框架使用示例

```go
import (
    "ecommerce/common/pkg/identity"
    "ecommerce/common/pkg/spi"
)

// 1. 初始化 SPI
registry := spi.NewExtensionRegistry()
loader := spi.NewExtensionLoader(registry)

// 2. 加载扩展
loader.LoadExtensions(
    id_market.InitIDMarket,
    us_market.InitUSMarket,
    pre_sale.InitPreSale,
    auction.InitAuction,
)

// 3. 为特定业务身份使用扩展
id := identity.NewBusinessIdentity(identity.CountryCN, identity.ModePreSale)
if productExt, ok := loader.GetProductExtension(id); ok {
    price, _ := productExt.CalculatePrice(ctx, product, id)
}
```

## API 接口文档

### 认证接口
- `POST /api/auth/register` - 用户注册
- `POST /api/auth/login` - 用户登录

### 商品接口（公开）
- `GET /api/products` - 获取商品列表
- `GET /api/products/:id` - 获取商品详情

### 需要认证的接口

#### 商家管理
- `POST /api/merchants` - 创建商家
- `GET /api/merchants` - 获取商家列表
- `GET /api/merchants/:id` - 获取商家详情
- `PUT /api/merchants/:id` - 更新商家

#### 商品管理
- `POST /api/products` - 创建商品
- `PUT /api/products/:id` - 更新商品
- `DELETE /api/products/:id` - 删除商品

#### 购物车
- `POST /api/cart` - 添加到购物车
- `GET /api/cart` - 获取购物车
- `PUT /api/cart/:id` - 更新购物车商品
- `DELETE /api/cart/:id` - 删除购物车商品

#### 订单管理
- `POST /api/orders` - 创建订单
- `GET /api/orders` - 获取订单列表
- `GET /api/orders/:id` - 获取订单详情
- `PUT /api/orders/:id/status` - 更新订单状态

#### 拍卖接口
- `GET /api/auctions` - 获取拍卖列表
- `GET /api/auctions/:id` - 获取拍卖详情
- `POST /api/auctions` - 创建拍卖
- `POST /api/auctions/:id/bid` - 出价
- `WebSocket /api/ws/auctions/:id` - 实时出价

## 数据库表结构

### 核心表
- **users**: 用户账户
- **merchants**: 商家档案
- **products**: 商品信息
- **product_prices**: 多货币价格
- **carts**: 购物车
- **orders**: 订单
- **order_items**: 订单项

### 扩展表（可选）
- **auctions**: 拍卖信息
- **bids**: 出价记录

## 开发说明

### Platform v4.0 迁移
- 旧版服务继续在端口 8081-8084 运行
- 新平台服务计划在端口 9081-9083 运行
- SPI 扩展实现业务功能隔离
- 平滑迁移，保持向后兼容

### SPI 扩展开发
1. 实现扩展点接口
2. 提供 `Init*()` 注册函数
3. 使用业务身份注册扩展
4. 在平台服务中加载扩展

## 许可证

MIT License
