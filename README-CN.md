# 电商交易平台

一个从0到1开发的完整电商交易平台，采用前后端分离架构。

## 技术栈

### 后端
- **语言**: Golang
- **Web框架**: Hertz
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT
- **密码加密**: bcrypt

### 前端
- **框架**: React 18
- **UI组件库**: Ant Design 5
- **路由**: React Router v6
- **HTTP客户端**: Axios
- **构建工具**: Vite
- **语言**: TypeScript

## 项目结构

```
trae_using_demo/
├── backend/              # 后端服务
│   ├── cmd/             # 应用入口
│   │   └── server/
│   ├── internal/        # 内部代码
│   │   ├── handler/     # HTTP处理器
│   │   ├── service/     # 业务逻辑层
│   │   ├── repository/  # 数据访问层
│   │   ├── model/       # 数据模型
│   │   └── middleware/  # 中间件
│   ├── pkg/             # 公共包
│   │   ├── response/    # 统一响应格式
│   │   └── utils/       # 工具函数
│   ├── bin/             # 编译输出
│   └── go.mod
├── frontend/            # 前端应用
│   ├── src/
│   │   ├── components/  # 公共组件
│   │   ├── pages/       # 页面组件
│   │   ├── services/    # API服务
│   │   ├── store/       # 状态管理
│   │   ├── types/       # TypeScript类型
│   │   ├── App.tsx
│   │   └── main.tsx
│   └── package.json
└── README.md
```

## 功能特性

### 用户认证
- 用户注册
- 用户登录
- JWT Token认证（有效期24小时）
- 密码bcrypt加密存储

### 商家管理
- 商家入驻
- 商家信息编辑
- 商家列表查询
- 商家详情查看

### 商品管理
- 商品发布
- 商品列表（公开访问）
- 商品详情
- 商品编辑
- 商品删除
- 库存管理

### 购物车
- 添加商品到购物车
- 购物车列表
- 修改商品数量
- 删除购物车商品

### 订单管理
- 创建订单（事务处理）
- 订单列表
- 订单详情
- 订单状态更新
- 库存自动扣减
- 购物车自动清空

## 快速开始

### 后端启动

```bash
cd backend
go build -o bin/server ./cmd/server
./bin/server
```

后端服务将在 `http://localhost:8080` 启动

### 前端启动

```bash
cd frontend
npm install
npm run dev
```

前端开发服务器将在 `http://localhost:5173` 启动

## API接口文档

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
- `PUT /api/merchants/:id` - 更新商家信息

#### 商品管理
- `POST /api/products` - 创建商品
- `PUT /api/products/:id` - 更新商品
- `DELETE /api/products/:id` - 删除商品

#### 购物车
- `POST /api/cart` - 添加商品到购物车
- `GET /api/cart` - 获取购物车
- `PUT /api/cart/:id` - 更新购物车商品
- `DELETE /api/cart/:id` - 删除购物车商品

#### 订单管理
- `POST /api/orders` - 创建订单
- `GET /api/orders` - 获取订单列表
- `GET /api/orders/:id` - 获取订单详情
- `PUT /api/orders/:id/status` - 更新订单状态

## 数据库表结构

### users（用户表）
- id: 主键
- username: 用户名（唯一）
- password: 密码（加密）
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

### merchants（商家表）
- id: 主键
- user_id: 用户ID（外键）
- name: 商家名称
- description: 商家描述
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

### products（商品表）
- id: 主键
- merchant_id: 商家ID（外键）
- name: 商品名称
- description: 商品描述
- price: 商品价格
- stock: 库存数量
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

### carts（购物车表）
- id: 主键
- user_id: 用户ID（外键）
- product_id: 商品ID（外键）
- quantity: 数量
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

### orders（订单表）
- id: 主键
- user_id: 用户ID（外键）
- total_amount: 订单总金额
- status: 订单状态
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

### order_items（订单项表）
- id: 主键
- order_id: 订单ID（外键）
- product_id: 商品ID（外键）
- quantity: 数量
- price: 商品单价
- created_at: 创建时间
- updated_at: 更新时间
- deleted_at: 删除时间（软删除）

## 开发说明

- 后端使用Hertz框架，GORM进行数据持久化
- 前端使用React + Ant Design，React Router进行路由管理
- 使用JWT进行身份认证，Token有效期24小时
- 密码使用bcrypt加密存储
- 创建订单时使用事务处理，确保数据一致性
- 支持CORS跨域请求

## 许可证

MIT License
