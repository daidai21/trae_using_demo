# 电商交易平台 - 实现计划（分解和优先级任务列表）

## [x] Task 1: 项目初始化 - 创建目录结构和配置文件
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 创建项目根目录结构（backend/、frontend/）
  - 初始化Go项目（go mod）
  - 初始化React项目（使用Vite）
  - 配置基础依赖
- **Acceptance Criteria Addressed**: 项目基础设施搭建
- **Test Requirements**:
  - `programmatic` TR-1.1: backend/ 和 frontend/ 目录成功创建
  - `programmatic` TR-1.2: backend/go.mod 文件存在且内容正确
  - `programmatic` TR-1.3: frontend/package.json 文件存在且包含基础依赖
- **Notes**: 使用Vite创建React项目以获得更好的开发体验

## [x] Task 2: 后端开发 - 数据库模型和初始化
- **Priority**: P0
- **Depends On**: Task 1
- **Description**: 
  - 定义数据模型（User、Merchant、Product、Cart、Order、OrderItem）
  - 配置GORM和SQLite数据库连接
  - 创建数据库自动迁移脚本
- **Acceptance Criteria Addressed**: AC-1, AC-2, AC-3
- **Test Requirements**:
  - `programmatic` TR-2.1: 所有数据模型定义正确
  - `programmatic` TR-2.2: 数据库连接配置正确
  - `programmatic` TR-2.3: 运行迁移后所有表成功创建
- **Notes**: 使用GORM的AutoMigrate功能

## [x] Task 3: 后端开发 - 用户认证API
- **Priority**: P0
- **Depends On**: Task 2
- **Description**: 
  - 实现用户注册接口（POST /api/auth/register）
  - 实现用户登录接口（POST /api/auth/login）
  - 实现JWT Token生成和验证中间件
  - 密码使用bcrypt加密
- **Acceptance Criteria Addressed**: AC-1
- **Test Requirements**:
  - `programmatic` TR-3.1: 注册接口返回200且创建用户记录
  - `programmatic` TR-3.2: 登录接口返回200且返回JWT Token
  - `programmatic` TR-3.3: JWT中间件正确验证Token
  - `programmatic` TR-3.4: 密码已加密存储
- **Notes**: Token有效期设置为24小时

## [x] Task 4: 后端开发 - 商家管理API
- **Priority**: P0
- **Depends On**: Task 3
- **Description**: 
  - 实现商家入驻接口（POST /api/merchants）
  - 实现商家信息查询接口（GET /api/merchants/:id）
  - 实现商家列表接口（GET /api/merchants）
  - 实现商家信息更新接口（PUT /api/merchants/:id）
- **Acceptance Criteria Addressed**: AC-2
- **Test Requirements**:
  - `programmatic` TR-4.1: 商家入驻接口成功创建商家记录
  - `programmatic` TR-4.2: 商家列表接口返回所有商家
  - `programmatic` TR-4.3: 商家详情接口返回正确信息
  - `programmatic` TR-4.4: 商家更新接口正确修改信息
- **Notes**: 需要JWT认证

## [x] Task 5: 后端开发 - 商品管理API
- **Priority**: P0
- **Depends On**: Task 4
- **Description**: 
  - 实现商品发布接口（POST /api/products）
  - 实现商品列表接口（GET /api/products）
  - 实现商品详情接口（GET /api/products/:id）
  - 实现商品更新接口（PUT /api/products/:id）
  - 实现商品删除接口（DELETE /api/products/:id）
- **Acceptance Criteria Addressed**: AC-3, AC-4
- **Test Requirements**:
  - `programmatic` TR-5.1: 商品发布接口成功创建商品记录
  - `programmatic` TR-5.2: 商品列表接口返回所有商品
  - `programmatic` TR-5.3: 商品详情接口返回正确信息
  - `programmatic` TR-5.4: 商品更新和删除接口正常工作
- **Notes**: 商品关联到商家

## [x] Task 6: 后端开发 - 购物车API
- **Priority**: P0
- **Depends On**: Task 5
- **Description**: 
  - 实现添加商品到购物车接口（POST /api/cart）
  - 实现购物车列表接口（GET /api/cart）
  - 实现修改购物车商品数量接口（PUT /api/cart/:id）
  - 实现删除购物车商品接口（DELETE /api/cart/:id）
- **Acceptance Criteria Addressed**: AC-5
- **Test Requirements**:
  - `programmatic` TR-6.1: 添加商品到购物车成功
  - `programmatic` TR-6.2: 购物车列表返回当前用户的所有商品
  - `programmatic` TR-6.3: 修改数量和删除功能正常
- **Notes**: 需要JWT认证，购物车按用户隔离

## [x] Task 7: 后端开发 - 订单管理API
- **Priority**: P0
- **Depends On**: Task 6
- **Description**: 
  - 实现创建订单接口（POST /api/orders）
  - 实现订单列表接口（GET /api/orders）
  - 实现订单详情接口（GET /api/orders/:id）
  - 实现订单状态更新接口（PUT /api/orders/:id/status）
- **Acceptance Criteria Addressed**: AC-6, AC-7
- **Test Requirements**:
  - `programmatic` TR-7.1: 创建订单成功，库存扣减，购物车清空
  - `programmatic` TR-7.2: 订单列表返回当前用户的所有订单
  - `programmatic` TR-7.3: 订单详情包含订单和订单项信息
  - `programmatic` TR-7.4: 订单状态更新正常
- **Notes**: 创建订单时需要事务处理

## [x] Task 8: 后端开发 - CORS和错误处理
- **Priority**: P1
- **Depends On**: Task 3
- **Description**: 
  - 配置CORS中间件允许前端跨域访问
  - 统一错误处理和响应格式
- **Acceptance Criteria Addressed**: NFR-3
- **Test Requirements**:
  - `programmatic` TR-8.1: 前端可以正常调用后端API
  - `programmatic` TR-8.2: 错误响应格式统一
- **Notes**: CORS允许本地开发域名

## [ ] Task 9: 前端开发 - 项目基础配置
- **Priority**: P0
- **Depends On**: Task 1
- **Description**: 
  - 配置React Router v6路由
  - 配置Ant Design组件库
  - 配置Axios HTTP客户端
  - 实现基础布局组件
  - 实现认证状态管理
- **Acceptance Criteria Addressed**: AC-8
- **Test Requirements**:
  - `human-judgement` TR-9.1: 路由配置正确，可以在页面间导航
  - `human-judgement` TR-9.2: Ant Design组件正确渲染
  - `programmatic` TR-9.3: Axios配置正确，包含JWT Token拦截器
- **Notes**: 使用Context API或Zustand进行状态管理

## [ ] Task 10: 前端开发 - 用户认证页面
- **Priority**: P0
- **Depends On**: Task 9, Task 3
- **Description**: 
  - 实现登录页面
  - 实现注册页面
  - 实现登出功能
- **Acceptance Criteria Addressed**: AC-1, AC-8
- **Test Requirements**:
  - `human-judgement` TR-10.1: 登录页面UI美观，表单验证正常
  - `human-judgement` TR-10.2: 注册页面UI美观，表单验证正常
  - `programmatic` TR-10.3: 登录/注册成功后Token正确存储
- **Notes**: 使用Ant Design的Form组件

## [ ] Task 11: 前端开发 - 商家管理页面
- **Priority**: P1
- **Depends On**: Task 10, Task 4
- **Description**: 
  - 实现商家入驻页面
  - 实现商家信息编辑页面
  - 实现商家列表页面
- **Acceptance Criteria Addressed**: AC-2, AC-8
- **Test Requirements**:
  - `human-judgement` TR-11.1: 商家入驻表单功能正常
  - `human-judgement` TR-11.2: 商家列表展示正确
  - `human-judgement` TR-11.3: 商家信息编辑功能正常
- **Notes**: 需要登录后访问

## [ ] Task 12: 前端开发 - 商品管理页面
- **Priority**: P0
- **Depends On**: Task 11, Task 5
- **Description**: 
  - 实现商品列表页面（首页）
  - 实现商品详情页面
  - 实现商品发布/编辑页面（商家）
- **Acceptance Criteria Addressed**: AC-3, AC-4, AC-8
- **Test Requirements**:
  - `human-judgement` TR-12.1: 商品列表页面展示美观，商品卡片布局合理
  - `human-judgement` TR-12.2: 商品详情页面信息展示完整
  - `human-judgement` TR-12.3: 商品发布/编辑表单功能正常
- **Notes**: 商品列表作为首页

## [ ] Task 13: 前端开发 - 购物车页面
- **Priority**: P0
- **Depends On**: Task 12, Task 6
- **Description**: 
  - 实现购物车页面
  - 实现添加商品到购物车功能（从商品详情页）
  - 实现修改商品数量功能
  - 实现删除商品功能
- **Acceptance Criteria Addressed**: AC-5, AC-8
- **Test Requirements**:
  - `human-judgement` TR-13.1: 购物车页面展示正确
  - `human-judgement` TR-13.2: 添加商品到购物车功能正常
  - `human-judgement` TR-13.3: 修改数量和删除功能正常
- **Notes**: 需要登录后访问

## [ ] Task 14: 前端开发 - 订单管理页面
- **Priority**: P0
- **Depends On**: Task 13, Task 7
- **Description**: 
  - 实现创建订单功能（从购物车页面）
  - 实现订单列表页面
  - 实现订单详情页面
- **Acceptance Criteria Addressed**: AC-6, AC-7, AC-8
- **Test Requirements**:
  - `human-judgement` TR-14.1: 创建订单流程顺畅
  - `human-judgement` TR-14.2: 订单列表页面展示正确
  - `human-judgement` TR-14.3: 订单详情页面信息完整
- **Notes**: 需要登录后访问

## [ ] Task 15: 联调与测试
- **Priority**: P0
- **Depends On**: Task 8, Task 14
- **Description**: 
  - 前后端接口联调
  - 完整流程测试（注册→登录→浏览商品→加购物车→下单→查看订单）
  - 修复发现的Bug
- **Acceptance Criteria Addressed**: AC-1, AC-2, AC-3, AC-4, AC-5, AC-6, AC-7, AC-8
- **Test Requirements**:
  - `programmatic` TR-15.1: 所有API接口调用正常
  - `human-judgement` TR-15.2: 完整购物流程测试通过
  - `programmatic` TR-15.3: 无严重Bug
- **Notes**: 手动测试完整流程
