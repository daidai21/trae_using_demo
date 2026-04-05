# 电商交易平台 - 平台化架构升级 v4.0 - 实施任务清单

## Tasks

| ID | Title | Description | Priority | Status |
|----|-------|-------------|----------|--------|
| T-1 | 创建 SPI 扩展框架目录结构 | 在 common/go/pkg/ 下创建 spi/ 和 identity/ 目录 | P0 | pending |
| T-2 | 定义业务身份抽象 | 实现 BusinessIdentity 数据结构和解析方法 | P0 | pending |
| T-3 | 实现身份解析器 | 实现从请求和用户偏好解析业务身份 | P0 | pending |
| T-4 | 定义 ExtensionPoint 基础接口 | 定义扩展点基础接口 Name() 和 Priority() | P0 | pending |
| T-5 | 定义各领域扩展点接口 | 定义 ProductExtension, TradeExtension, I18nExtension, PaymentExtension | P0 | pending |
| T-6 | 实现 ExtensionRegistry | 实现扩展注册表，支持注册和查询扩展 | P0 | pending |
| T-7 | 创建 platform 目录 | 在 domain-services/ 下创建 platform/ 目录 | P0 | pending |
| T-8 | 迁移 user-service | 将 user-service 迁移到 platform/user-service | P0 | pending |
| T-9 | 重构 product-service | 迁移到 platform/product-service，移除预售和拍卖代码 | P0 | pending |
| T-10 | 重构 trade-service | 迁移到 platform/trade-service，移除预售和多市场代码 | P0 | pending |
| T-11 | 平台服务集成 SPI | 在平台服务关键路径添加 SPI 扩展点调用 | P0 | pending |
| T-12 | 创建 extensions 目录 | 在 domain-services/ 下创建 extensions/ 目录 | P0 | pending |
| T-13 | 创建 id-market 扩展 | 创建印尼市场扩展模块 | P0 | pending |
| T-14 | 创建 us-market 扩展 | 创建美国市场扩展模块 | P0 | pending |
| T-15 | 创建 pre-sale 扩展 | 创建预售功能扩展模块 | P0 | pending |
| T-16 | 创建 auction 扩展 | 将 auction-service 迁移为扩展模块 | P0 | pending |
| T-17 | 更新 api-gateway | 集成业务身份解析和路由 | P0 | pending |
| T-18 | 单元测试 | 为 SPI 框架和扩展编写单元测试 | P1 | pending |
| T-19 | 集成测试 | 验证平台和扩展的集成 | P1 | pending |
| T-20 | 文档编写 | 编写架构文档和扩展开发指南 | P2 | pending |

---

## 详细任务说明

### Phase 1: SPI 框架基础设施 (T-1 至 T-6)

**T-1: 创建 SPI 扩展框架目录结构**
- 创建 `common/go/pkg/spi/` 目录
- 创建 `common/go/pkg/identity/` 目录

**T-2: 定义业务身份抽象**
- 实现 `BusinessIdentity` 结构体
- 实现 `String()` 方法
- 实现 `Parse()` 方法

**T-3: 实现身份解析器**
- 定义 `IdentityResolver` 接口
- 实现从 HTTP 请求解析身份
- 实现从用户偏好解析身份

**T-4: 定义 ExtensionPoint 基础接口**
- 定义 `ExtensionPoint` 接口
- `Name() string` 方法
- `Priority() int` 方法

**T-5: 定义各领域扩展点接口**
- `ProductExtension` 接口
- `TradeExtension` 接口
- `I18nExtension` 接口
- `PaymentExtension` 接口

**T-6: 实现 ExtensionRegistry**
- 实现扩展注册表数据结构
- 实现 `Register()` 方法
- 实现 `GetExtensions()` 方法
- 实现 `GetExtension()` 方法

### Phase 2: 平台核心服务重构 (T-7 至 T-11)

**T-7: 创建 platform 目录**
- 创建 `domain-services/platform/` 目录

**T-8: 迁移 user-service**
- 复制 `domain-services/user-service` 到 `domain-services/platform/user-service`
- 更新 go.mod 和 import 路径

**T-9: 重构 product-service**
- 复制到 `domain-services/platform/product-service`
- 移除预售相关代码
- 移除拍卖相关代码
- 保留普通商品功能
- 集成 SPI 框架

**T-10: 重构 trade-service**
- 复制到 `domain-services/platform/trade-service`
- 移除预售相关代码
- 移除多市场相关代码
- 保留普通交易功能
- 集成 SPI 框架

**T-11: 平台服务集成 SPI**
- 在商品创建前调用扩展点
- 在订单创建前调用扩展点
- 在支付处理时调用扩展点

### Phase 3: 业务扩展模块 (T-12 至 T-16)

**T-12: 创建 extensions 目录**
- 创建 `domain-services/extensions/` 目录

**T-13: 创建 id-market 扩展**
- 创建 `extensions/id-market/` 目录
- 实现印尼语语言包
- 实现印尼支付网关集成
- 注册扩展

**T-14: 创建 us-market 扩展**
- 创建 `extensions/us-market/` 目录
- 实现美式英语语言包
- 实现美国支付网关集成
- 实现 Sales Tax 计算
- 注册扩展

**T-15: 创建 pre-sale 扩展**
- 创建 `extensions/pre-sale/` 目录
- 实现预售商品扩展
- 实现预售订单扩展
- 实现定金/尾款逻辑
- 注册扩展

**T-16: 创建 auction 扩展**
- 创建 `extensions/auction/` 目录
- 从 `auction-service` 迁移代码
- 实现拍卖商品扩展
- 实现拍卖交易扩展
- 保留 WebSocket 功能
- 注册扩展

### Phase 4: 网关适配和测试 (T-17 至 T-20)

**T-17: 更新 api-gateway**
- 集成业务身份解析
- 实现扩展路由
- 保持向后兼容

**T-18: 单元测试**
- SPI 框架单元测试
- 业务身份解析测试
- 扩展注册表测试
- 覆盖率 > 80%

**T-19: 集成测试**
- 平台核心服务集成测试
- 各扩展模块集成测试
- 向后兼容性验证

**T-20: 文档编写**
- 架构设计文档
- SPI 框架使用文档
- 扩展开发指南
