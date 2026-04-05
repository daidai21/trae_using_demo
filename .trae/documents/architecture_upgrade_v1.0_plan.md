# 电商交易平台架构升级计划 - V1.0

## 一、架构升级概述

### 1.1 升级目标

将现有单体架构升级为：

- **后端**：微服务架构，按领域分离；从大单体服务到微服务架构， 按照业务域拆分
- **前端**：Monorepo模式，按领域分离

### 1.2 升级原则

- 保持现有功能完整性
- 按业务领域清晰划分边界
- 便于后续扩展和维护
- 保持代码可读性和可维护性

***

## 二、后端微服务架构设计

### 2.1 微服务划分

按业务领域划分为核心微服务（V1.0版本实现）和扩展服务（后续版本）：

#### 2.1.1 用户服务 (user-service)

- **端口**: 8081
- **职责**:
  - 用户注册/登录
  - JWT Token生成和验证
  - 用户信息管理
  - 密码安全（bcrypt加密）
- **数据库**: users.db (SQLite)
- **数据表**: users
- **API接口**:
  - `POST /api/auth/register` - 用户注册
  - `POST /api/auth/login` - 用户登录
  - `GET /api/users/profile` - 获取用户信息
  - `PUT /api/users/profile` - 更新用户信息

#### 2.1.2 商家商品服务 (product-service)

- **端口**: 8082
- **职责**:
  - 商家入驻和管理
  - 商品发布和管理
  - 商品列表和详情查询（公开访问）
  - 库存管理
- **数据库**: merchant.db, product.db (SQLite)
- **数据表**: merchants, products
- **API接口**:
  - `POST /api/merchants` - 商家入驻
  - `GET /api/merchants` - 商家列表
  - `GET /api/merchants/:id` - 商家详情
  - `PUT /api/merchants/:id` - 更新商家信息
  - `POST /api/products` - 发布商品
  - `GET /api/products` - 商品列表（公开）
  - `GET /api/products/:id` - 商品详情（公开）
  - `PUT /api/products/:id` - 更新商品
  - `DELETE /api/products/:id` - 删除商品

#### 2.2.1 导购服务 (guide-service)

- **职责**:
  - 商品搜索
  - 商品推荐
  - PDP（商品详情页聚合）
  - 商城首页聚合
- **数据库**: product.db (SQLite)
- **数据表**: products（只读）

#### 2.1.3 交易服务 (trade-service)

- **端口**: 8083
- **职责**:
  - **cart（购物车域）**
    - 购物车管理
    - 商品添加、删除、修改数量
  - **buy（购买域）**
    - 购买商品
    - 订单创建（事务处理）
    - 库存扣减
  - **order（订单域）**
    - 订单模型
    - 订单管理：C端查单、B端查单
    - 订单状态流转
  - **reverse（售后域，V1.0预留）**
    - 售后
    - 逆向
- **数据库**: order.db (SQLite)
- **数据表**: carts, orders, order_items
- **API接口**:
  - `POST /api/cart` - 添加商品到购物车
  - `GET /api/cart` - 获取购物车
  - `PUT /api/cart/:id` - 更新购物车商品数量
  - `DELETE /api/cart/:id` - 删除购物车商品
  - `POST /api/orders` - 创建订单
  - `GET /api/orders` - 订单列表
  - `GET /api/orders/:id` - 订单详情
  - `PUT /api/orders/:id/status` - 更新订单状态



#### 2.2.2 营销服务 (marketing-service)

- **职责**:
  - 营销C端：优惠券、满减、秒杀
  - 营销B端：营销活动配置
  - 营销模型：优惠券、活动规则
  - 营销工具：活动创建、数据统计

#### 2.2.4 支付服务 (payment-service)

- **职责**:
  - 支付C端：支付渠道集成
  - 支付B端：支付配置
  - 支付模型：支付流水、退款记录
  - 支付工具：回调处理、状态同步

#### 2.2.3 资金服务 (fund-service)

- **职责**:
  - 结算管理
  - 税费计算
  - 资金流水
  - 财务报表

#### 2.2.5 物流服务 (logistics-service)

- **职责**:
  - 物流单号管理
  - 物流状态跟踪
  - 物流服务商对接
  - 发货管理

#### 2.2.6 治理服务 (governance-service)

- **职责**:
  - 风控管理
  - 审核管理
  - 数据统计
  - 系统监控

---

### 2.3 后端目录结构

```
trae_using_demo/
├── domain-services/              # 微服务目录
│   ├── api-gateway/            # API网关
│   │   ├── cmd/gateway/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── proxy/         # 反向代理
│   │   │   └── middleware/
│   │   ├── pkg/
│   │   └── go.mod
│   ├── user-service/           # 用户服务
│   │   ├── cmd/server/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   └── model/
│   │   ├── pkg/
│   │   └── go.mod
│   ├── product-service/        # 商家商品服务
│   │   ├── cmd/server/
│   │   │   └── main.go
│   │   ├── internal/
│   │   │   ├── handler/
│   │   │   ├── service/
│   │   │   ├── repository/
│   │   │   └── model/
│   │   ├── pkg/
│   │   └── go.mod
│   └── trade-service/          # 交易服务
│       ├── cmd/server/
│       │   └── main.go
│       ├── internal/
│       │   ├── handler/
│       │   ├── service/
│       │   │   ├── cart/       # 购物车域
│       │   │   ├── buy/        # 购买域
│       │   │   └── order/      # 订单域
│       │   ├── repository/
│       │   └── model/
│       ├── pkg/
│       └── go.mod
├── shared/                     # 共享代码库
│   ├── go/
│   │   ├── pkg/
│   │   │   ├── response/      # 统一响应格式
│   │   │   ├── utils/         # 工具函数
│   │   │   └── constants/     # 常量定义
│   │   └── go.mod
│   └── proto/                 # 服务间通信定义（可选）
└── docs/                       # 文档
```

---

### 2.4 服务间通信

- **同步通信**: HTTP REST API（V1.0简化版本）
- **数据一致性**: 最终一致性，每个服务独立管理自己的数据
- **服务发现**: 静态配置（V1.0），后续可引入服务注册中心
- **容错机制**: 超时控制、重试机制（后续版本）

---

### 2.5 数据库设计

每个微服务拥有独立的数据库，按业务域隔离：

#### 用户服务数据库 (users.db)
- `users` - 用户表

#### 商家商品服务数据库 (product.db)
- `merchants` - 商家表
- `products` - 商品表

#### 交易服务数据库 (order.db)
- `carts` - 购物车表
- `orders` - 订单表
- `order_items` - 订单详情表

---

### 2.6 部署架构

```
                    ┌─────────────┐
                    │   客户端     │
                    └──────┬──────┘
                           │
                    ┌──────▼──────┐
                    │  API网关     │ (8080)
                    └──────┬──────┘
           ┌───────────────┼───────────────┐
    ┌──────▼──────┐  ┌───▼────┐  ┌─────▼──────┐
    │ 用户服务     │  │商家商品 │  │  交易服务   │
    │  (8081)     │  │ (8082)  │  │   (8083)   │
    └──────┬──────┘  └───┬────┘  └─────┬──────┘
           │               │               │
    ┌──────▼──────┐  ┌───▼────┐  ┌─────▼──────┐
    │ users.db    │  │product. │  │  order.db  │
    └─────────────┘  │  db     │  └────────────┘
                     └─────────┘
```

***

## 三、前端Monorepo架构设计

### 3.1 Monorepo结构

使用pnpm workspaces实现Monorepo：

```
trae_using_demo/
├── apps/                      # 应用目录
│   └── web/                  # 主Web应用
│       ├── src/
│       ├── package.json
│       └── vite.config.ts
├── packages/                  # 包目录
│   ├── ui/                   # 共享UI组件库
│   │   ├── src/
│   │   └── package.json
│   ├── auth/                 # 用户认证模块
│   │   ├── src/
│   │   └── package.json
│   ├── product/              # 商品商家模块
│   │   ├── src/
│   │   └── package.json
│   └── order/                # 订单购物车模块
│       ├── src/
│       └── package.json
├── package.json
├── pnpm-workspace.yaml
└── pnpm-lock.yaml
```

### 3.2 模块划分

#### 3.2.1 packages/auth (用户认证模块)

- **职责**:
  - 登录/注册页面
  - 认证状态管理
  - 用户信息管理
- **导出**:
  - 组件: LoginForm, RegisterForm, UserProfile
  - Hooks: useAuth
  - Services: authService

#### 3.2.2 packages/product (商品商家模块)

- **职责**:
  - 商品列表
  - 商品详情
  - 商家管理
  - 商品发布
- **导出**:
  - 组件: ProductList, ProductDetail, MerchantList, ProductForm
  - Hooks: useProducts, useMerchants
  - Services: productService, merchantService

#### 3.2.3 packages/order (订单购物车模块)

- **职责**:
  - 购物车
  - 订单列表
  - 订单详情
- **导出**:
  - 组件: Cart, OrderList, OrderDetail
  - Hooks: useCart, useOrders
  - Services: cartService, orderService

#### 3.2.4 packages/ui (共享UI组件库)

- **职责**:
  - 基础布局组件
  - 通用业务组件
  - 工具函数
- **导出**:
  - 组件: Layout, Navbar, ProductCard, Loading, ErrorBoundary
  - Hooks: useApi, useNotification
  - Utils: apiClient, constants

#### 3.2.5 apps/web (主应用)

- **职责**:
  - 路由配置
  - 模块整合
  - 全局状态管理
  - 应用入口

***

## 四、升级步骤

### 阶段一：架构规划和准备 ✅

1. 制定架构升级计划
2. 准备微服务和Monorepo基础结构

### 阶段二：后端微服务改造

1. 创建shared共享代码库
2. 拆分user-service
3. 拆分product-service
4. 拆分order-service
5. 实现api-gateway
6. 后端联调测试

### 阶段三：前端Monorepo改造

1. 配置pnpm workspaces
2. 创建packages/ui共享组件库
3. 拆分packages/auth模块
4. 拆分packages/product模块
5. 拆分packages/order模块
6. 整合apps/web主应用
7. 前端联调测试

### 阶段四：整体联调与文档

1. 前后端整体联调
2. 更新文档
3. 编写升级指南

***

## 五、技术选型

### 后端微服务

- **Web框架**: Hertz（保持不变）
- **服务间通信**: HTTP REST
- **API网关**: 自定义实现（反向代理）
- **共享代码**: Go Module

### 前端Monorepo

- **包管理**: pnpm workspaces
- **构建工具**: Vite（保持不变）
- **类型系统**: TypeScript（保持不变）
- **UI组件**: Ant Design（保持不变）

***

## 六、风险与应对

| 风险            | 影响 | 应对措施                   |
| ------------- | -- | ---------------------- |
| 服务拆分后数据一致性问题  | 高  | 采用最终一致性，明确服务边界         |
| Monorepo配置复杂度 | 中  | 使用pnpm workspaces，逐步迁移 |
| 开发环境复杂度增加     | 中  | 提供docker-compose一键启动   |
| 部署复杂度增加       | 中  | 提供统一的启动脚本              |

***

## 七、交付物

1. ✅ 架构升级计划文档
2. 后端微服务代码（4个服务）
3. 前端Monorepo代码（4个packages + 1个app）
4. shared共享代码库
5. 更新的README文档
6. 架构升级指南

