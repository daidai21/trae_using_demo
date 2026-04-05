# 电商交易平台 - 业务增长迭代 v1.0 - 实施计划

## [ ] Task 1: 国际化基础设施搭建
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 前端i18n框架选型和搭建（react-i18next）
  - 后端多语言消息机制
  - 语言包目录结构
  - 语言切换组件
- **Acceptance Criteria Addressed**: [AC-1]
- **Test Requirements**:
  - `programmatic` TR-1.1: i18n框架初始化成功
  - `human-judgement` TR-1.2: 语言切换UI可用
- **Notes**: 支持后续至少5种语言扩展

## [ ] Task 2: 印尼语语言包 (id-ID)
- **Priority**: P0
- **Depends On**: Task 1
- **Description**: 
  - 翻译所有UI文本为印尼语
  - 准备id-ID语言包文件
  - 集成到i18n系统
- **Acceptance Criteria Addressed**: [AC-1]
- **Test Requirements**:
  - `programmatic` TR-2.1: id-ID语言包加载成功
  - `human-judgement` TR-2.2: 印尼语文本翻译质量检查
- **Notes**: 需要专业翻译人员

## [ ] Task 3: 多货币支持 (product-service)
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 商品模型支持多货币价格
  - 价格显示格式化（IDR, USD）
  - 货币汇率配置
- **Acceptance Criteria Addressed**: [AC-2, AC-3]
- **Test Requirements**:
  - `programmatic` TR-3.1: IDR格式显示正确 (Rp X.XXX.XXX)
  - `programmatic` TR-3.2: USD格式显示正确 ($X,XXX.XX)
- **Notes**: 后续可扩展更多货币

## [ ] Task 4: 订单多货币支持 (trade-service)
- **Priority**: P0
- **Depends On**: Task 3
- **Description**: 
  - 订单模型支持多货币金额
  - 税费计算适配（美国Sales Tax）
  - 支付网关调用传递正确货币
- **Acceptance Criteria Addressed**: [AC-2, AC-3]
- **Test Requirements**:
  - `programmatic` TR-4.1: 订单金额以对应货币存储和显示
- **Notes**: 对接支付网关前准备

## [ ] Task 5: auction-service 数据模型和基础设施
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 创建auction-service目录结构
  - 设计Auction和Bid数据模型
  - 配置数据库（auction.db）
  - go.mod和依赖初始化
- **Acceptance Criteria Addressed**: [AC-4]
- **Test Requirements**:
  - `programmatic` TR-5.1: auction-service编译成功
  - `programmatic` TR-5.2: 数据库表创建成功
- **Notes**: 端口建议8084

## [ ] Task 6: WebSocket Hub (auction-service)
- **Priority**: P0
- **Depends On**: Task 5
- **Description**: 
  - 实现WebSocket连接管理
  - 消息广播机制
  - 在线人数统计
  - 房间概念（每个拍卖一个房间）
- **Acceptance Criteria Addressed**: [AC-5]
- **Test Requirements**:
  - `programmatic` TR-6.1: WebSocket连接建立成功
  - `programmatic` TR-6.2: 消息广播功能正常
- **Notes**: 支持至少1000并发

## [ ] Task 7: 拍卖CRUD和出价功能 (auction-service)
- **Priority**: P0
- **Depends On**: Task 6
- **Description**: 
  - 拍卖CRUD API
  - 出价验证逻辑（高于当前最高价+加价幅度）
  - 出价历史记录
  - 拍卖状态流转
- **Acceptance Criteria Addressed**: [AC-4, AC-5]
- **Test Requirements**:
  - `programmatic` TR-7.1: 拍卖创建成功
  - `programmatic` TR-7.2: 出价验证正确执行
  - `programmatic` TR-7.3: 低于当前最高价的出价被拒绝
- **Notes**: 集成WebSocket实时更新

## [ ] Task 8: 预售商品字段 (product-service)
- **Priority**: P0
- **Depends On**: None
- **Description**: 
  - 商品类型新增"预售商品"
  - 预售价格、定金金额字段
  - 预售开始/结束时间、尾款截止时间字段
- **Acceptance Criteria Addressed**: [AC-6]
- **Test Requirements**:
  - `programmatic` TR-8.1: 预售商品创建成功
  - `programmatic` TR-8.2: 预售时间和价格字段存储正确
- **Notes**: 与普通商品共用表，type字段区分

## [ ] Task 9: 预售订单和支付 (trade-service)
- **Priority**: P0
- **Depends On**: Task 8
- **Description**: 
  - 预售订单状态机设计
  - 定金支付（库存锁定）
  - 尾款支付（提醒、入口）
  - 逾期未支付自动取消
- **Acceptance Criteria Addressed**: [AC-6, AC-7]
- **Test Requirements**:
  - `programmatic` TR-9.1: 定金支付后订单状态变为deposit_paid
  - `programmatic` TR-9.2: 尾款支付后订单状态变为completed
- **Notes**: 包含库存锁定逻辑

## [ ] Task 10: 前端国际化改造 (packages/ui, apps/web)
- **Priority**: P0
- **Depends On**: Task 1
- **Description**: 
  - packages/ui新增i18n配置
  - 语言切换组件
  - apps/web根级i18n初始化
  - 所有页面文本替换为t()函数
- **Acceptance Criteria Addressed**: [AC-1]
- **Test Requirements**:
  - `programmatic` TR-10.1: i18n在前端初始化成功
  - `human-judgement` TR-10.2: 语言切换功能正常
- **Notes**: 保持现有功能完整性

## [ ] Task 11: 前端拍卖域 (packages/auction)
- **Priority**: P1
- **Depends On**: None
- **Description**: 
  - 创建packages/auction目录
  - AuctionList页面（拍卖列表）
  - AuctionDetail页面（含出价UI）
  - CreateAuction页面（发起拍卖）
  - useAuctionSocket WebSocket Hook
- **Acceptance Criteria Addressed**: [AC-4, AC-5]
- **Test Requirements**:
  - `programmatic` TR-11.1: 拍卖列表加载成功
  - `human-judgement` TR-11.2: 出价UI交互流畅
  - `programmatic` TR-11.3: WebSocket实时更新正常
- **Notes**: 端口8084对接

## [ ] Task 12: 前端预售功能 (packages/product, packages/trade)
- **Priority**: P1
- **Depends On**: Task 8
- **Description**: 
  - 商品列表/详情页显示预售标识
  - 预售信息展示（定金、尾款时间等）
  - 预售下单流程
  - 定金/尾款支付UI
- **Acceptance Criteria Addressed**: [AC-6, AC-7]
- **Test Requirements**:
  - `programmatic` TR-12.1: 预售商品标识显示正确
  - `human-judgement` TR-12.2: 预售下单流程完整
- **Notes**: 与普通订单体验一致

## [ ] Task 13: 印尼支付网关对接 (可选Phase 2)
- **Priority**: P1
- **Depends On**: Task 4
- **Description**: 
  - 对接Midtrans等印尼支付网关
  - 支付回调处理
  - 手续费计算
- **Acceptance Criteria Addressed**: [AC-2]
- **Test Requirements**:
  - `programmatic` TR-13.1: 印尼支付流程完整
- **Notes**: 可延后到Phase 2

## [ ] Task 14: 美国支付网关对接 (可选Phase 2)
- **Priority**: P1
- **Depends On**: Task 4
- **Description**: 
  - 对接Stripe/PayPal等美国支付网关
  - ACH转账支持（如需要）
  - Sales Tax计算
- **Acceptance Criteria Addressed**: [AC-3]
- **Test Requirements**:
  - `programmatic` TR-14.1: 美国支付流程完整
- **Notes**: 可延后到Phase 2
