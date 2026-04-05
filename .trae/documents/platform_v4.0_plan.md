# 电商交易平台 - 平台化架构升级计划 v4.0

## 概述

本计划将当前业务增长架构升级为平台化架构，通过SPI（Service Provider Interface）扩展机制实现平台与业务的分离。

**注：本次仅进行后端架构升级，前端保持不变。**

## 设计理念

### 1. 业务身份抽象

采用 `{国家}.{玩法}` 的业务身份标识体系：

```
CN.normal  - 中国普通交易
ID.normal  - 印尼普通交易
US.normal  - 美国普通交易
CN.preSale - 中国预售
CN.auction - 中国拍卖
...
```

### 2. 平台与业务分离

**平台层**
- 保留核心普通交易模式
- 提供SPI扩展接口
- 负责扩展点管理和加载

**业务扩展层**
- ID市场扩展
- US市场扩展
- 预售功能扩展
- 直播拍卖扩展

## 架构设计

### 目录结构

```
domain-services/
├── common/                    # 公共包
│   └── go/
│       └── pkg/
│           ├── spi/          # 新增：SPI扩展框架
│           │   ├── extension.go
│           │   ├── registry.go
│           │   └── loader.go
│           ├── identity/     # 新增：业务身份抽象
│           │   ├── identity.go
│           │   └── resolver.go
│           └── ...
├── platform/                  # 新增：平台核心服务
│   ├── user-service/         # 平台用户服务
│   ├── product-service/      # 平台商品服务（仅普通模式）
│   └── trade-service/        # 平台交易服务（仅普通模式）
└── extensions/               # 新增：业务扩展目录
    ├── id-market/            # ID市场扩展
    ├── us-market/            # US市场扩展
    ├── pre-sale/             # 预售扩展
    └── auction/              # 拍卖扩展
```

### SPI扩展框架设计

#### ExtensionPoint - 扩展点定义

```go
// ExtensionPoint 扩展点接口
type ExtensionPoint interface {
    Name() string
    Priority() int
}

// ProductExtension 商品扩展点
type ProductExtension interface {
    ExtensionPoint
    // 商品类型扩展
    GetProductTypes() []string
    // 价格计算扩展
    CalculatePrice(ctx context.Context, product *model.Product, identity *identity.BusinessIdentity) (float64, error)
}

// TradeExtension 交易扩展点
type TradeExtension interface {
    ExtensionPoint
    // 订单创建扩展
    BeforeCreateOrder(ctx context.Context, order *model.Order, identity *identity.BusinessIdentity) error
    // 支付处理扩展
    ProcessPayment(ctx context.Context, order *model.Order, identity *identity.BusinessIdentity) error
}

// I18nExtension 国际化扩展点
type I18nExtension interface {
    ExtensionPoint
    // 语言包
    GetLocales() map[string]map[string]string
    // 货币格式化
    FormatCurrency(amount float64, currency string) string
}

// PaymentExtension 支付扩展点
type PaymentExtension interface {
    ExtensionPoint
    // 支付网关
    GetPaymentGateways() []string
    // 发起支付
    InitiatePayment(ctx context.Context, order *model.Order, gateway string) (string, error)
}
```

#### ExtensionRegistry - 扩展注册器

```go
// ExtensionRegistry 扩展注册表
type ExtensionRegistry struct {
    extensions map[string]map[string]ExtensionPoint
    mu         sync.RWMutex
}

// Register 注册扩展
func (r *ExtensionRegistry) Register(identity *identity.BusinessIdentity, ext ExtensionPoint)

// GetExtensions 获取指定业务身份的扩展
func (r *ExtensionRegistry) GetExtensions(identity *identity.BusinessIdentity) []ExtensionPoint

// GetExtension 获取指定类型的扩展
func (r *ExtensionRegistry) GetExtension(identity *identity.BusinessIdentity, extType string) (ExtensionPoint, bool)
```

### 业务身份设计

```go
// BusinessIdentity 业务身份
type BusinessIdentity struct {
    Country string // 国家代码 (CN, ID, US)
    Mode    string // 玩法模式 (normal, pre_sale, auction)
}

// String 身份字符串表示
func (id *BusinessIdentity) String() string {
    return fmt.Sprintf("%s.%s", id.Country, id.Mode)
}

// Parse 解析业务身份
func Parse(identity string) (*BusinessIdentity, error)

// Resolver 身份解析器
type IdentityResolver interface {
    // 从请求中解析身份
    Resolve(ctx context.Context, req *http.Request) (*BusinessIdentity, error)
    // 从用户偏好解析身份
    ResolveFromUser(userID uint) (*BusinessIdentity, error)
}
```

## 实施步骤

### Step 1: SPI框架基础设施

**文件清单**
- `common/go/pkg/spi/extension.go` - 扩展点接口定义
- `common/go/pkg/spi/registry.go` - 扩展注册表
- `common/go/pkg/spi/loader.go` - 扩展加载器
- `common/go/pkg/identity/identity.go` - 业务身份定义
- `common/go/pkg/identity/resolver.go` - 身份解析器

**任务**
1. 定义 ExtensionPoint 基础接口
2. 定义各领域扩展点接口（Product, Trade, I18n, Payment）
3. 实现 ExtensionRegistry 扩展注册表
4. 实现 BusinessIdentity 业务身份抽象
5. 实现 IdentityResolver 身份解析器

### Step 2: 平台核心服务重构

**文件清单**
- 创建 `domain-services/platform/` 目录
- 重构 `user-service` → `platform/user-service`
- 重构 `product-service` → `platform/product-service`（移除预售、拍卖相关代码）
- 重构 `trade-service` → `platform/trade-service`（移除预售、多市场代码）

**任务**
1. 创建 platform 目录结构
2. 迁移核心服务到 platform 目录
3. 剥离业务扩展代码，仅保留普通交易模式
4. 在平台服务中集成 SPI 框架
5. 在关键路径添加扩展点调用

### Step 3: 业务扩展模块创建

**目录结构**
```
extensions/
├── id-market/
│   ├── go.mod
│   └── internal/
│       ├── extension/
│       │   ├── i18n_extension.go
│       │   └── payment_extension.go
│       └── init.go
├── us-market/
│   ├── go.mod
│   └── internal/
│       ├── extension/
│       │   ├── i18n_extension.go
│       │   ├── payment_extension.go
│       │   └── tax_extension.go
│       └── init.go
├── pre-sale/
│   ├── go.mod
│   └── internal/
│       ├── extension/
│       │   ├── product_extension.go
│       │   └── trade_extension.go
│       └── init.go
└── auction/
    ├── go.mod
    └── internal/
        ├── extension/
        │   ├── product_extension.go
        │   └── trade_extension.go
        └── init.go
```

**任务**
1. 创建 id-market 扩展模块（印尼语语言包、印尼支付网关）
2. 创建 us-market 扩展模块（美式英语、美国支付网关、Sales Tax计算）
3. 创建 pre-sale 扩展模块（预售商品、预售订单、定金/尾款）
4. 创建 auction 扩展模块（拍卖服务迁移至此）

### Step 4: API网关适配

**任务**
1. 更新 api-gateway，集成业务身份解析
2. 根据业务身份路由到相应的扩展服务
3. 保持向后兼容

## 关键设计决策

### 1. 扩展加载策略

- **编译时加载**：通过 Go plugin 或导入初始化
- **运行时发现**：通过配置文件或注册中心发现扩展
- **优先级机制**：支持扩展优先级，解决冲突

### 2. 身份解析策略

- 优先级：URL参数 > Header > Cookie > 用户偏好 > 默认
- 支持动态切换业务身份
- 身份变更时重新加载扩展

### 3. 扩展点命名规范

```
{domain}.{action}.{phase}

例：
product.create.before
trade.order.after
payment.process
```

## 风险与缓解

| 风险 | 影响 | 缓解措施 |
|------|------|----------|
| 扩展加载性能 | 中 | 实现扩展缓存，延迟加载 |
| 扩展冲突 | 高 | 优先级机制，明确扩展契约 |
| 向后兼容性 | 高 | 提供默认实现，兼容现有API |
| 调试复杂度 | 中 | 完善的扩展日志和监控 |

## 成功指标

- [ ] SPI框架代码覆盖率 > 80%
- [ ] 所有4个业务扩展正常运行
- [ ] 平台核心代码无业务逻辑
- [ ] 扩展点文档完整
- [ ] 性能损耗 < 5%
