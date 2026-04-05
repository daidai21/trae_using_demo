# 电商交易平台 - 业务增长迭代 v1.0 - 产品需求文档 (PRD)

## Overview
- **Summary**: 本计划涵盖电商交易平台的4个核心业务增长功能迭代开发，包括：开放印度尼西亚市场、开放美国市场、支持直播拍卖、支持预售功能。
- **Purpose**: 实现市场扩展和新交易模式，提升平台GMV和用户增长。
- **Target Users**: 全球买家和卖家，特别是印尼和美国市场用户。

## Goals
1. 开放印度尼西亚市场，支持印尼语和印尼盾支付
2. 开放美国市场，支持英语和美元支付
3. 支持直播拍卖功能，提供实时出价和WebSocket通信
4. 支持预售功能，实现定金/尾款支付模式

## Non-Goals (Out of Scope)
- 不包括其他国家/地区市场扩展（除ID和US外）
- 不包括VR/AR购物功能
- 不包括社交分享功能
- 不包括智能推荐系统

## Background & Context
- 现有平台已完成从单体架构到微服务+Monorepo的架构升级
- 后端：api-gateway, user-service, product-service, trade-service
- 前端：Monorepo架构（packages/ui, packages/user, packages/product, packages/trade, apps/web）
- 基于现有架构进行业务功能迭代

## Functional Requirements
- **FR-1**: 开放印度尼西亚市场 (Open ID Market)
  - 支持id-ID语言包（印尼语）
  - 支持IDR货币（印尼盾）
  - 对接印尼本地支付网关（Midtrans等）
  - 时区、地址格式等本地化适配
  
- **FR-2**: 开放美国市场 (Open US Market)
  - 完善en-US语言包（美式英语）
  - 支持USD货币（美元）
  - 对接美国支付网关（Stripe, PayPal等）
  - 美国税务计算（Sales Tax）和地址格式
  
- **FR-3**: 支持直播拍卖 (Support Live Auction)
  - 拍卖商品类型（起拍价、加价幅度、拍卖时长）
  - 实时出价功能（出价验证、历史记录）
  - WebSocket实时通信（状态推送、出价广播）
  - 可选：直播流集成（阿里云/腾讯云直播）
  - 拍卖订单自动生成
  
- **FR-4**: 支持预售 (Support Pre-Sale)
  - 预售商品类型（预售价格、定金金额、时间配置）
  - 定金支付（库存锁定）
  - 尾款支付（提醒、入口、逾期处理）
  - 预售订单状态机

## Non-Functional Requirements
- **NFR-1**: 国际化架构支持至少5种语言扩展
- **NFR-2**: WebSocket支持至少1000并发连接
- **NFR-3**: 页面加载时间 < 2秒
- **NFR-4**: 支付成功率 > 95%
- **NFR-5**: 系统可用性 > 99.9%

## Constraints
- **Technical**: 基于现有微服务和Monorepo架构，技术栈不变
- **Business**: 4个功能总周期5个月（市场扩展2个月+新模式3个月）
- **Dependencies**: 
  - 支付网关：Midtrans (ID), Stripe/PayPal (US)
  - 直播服务：阿里云/腾讯云（可选）

## Assumptions
- 现有架构（微服务+Monorepo）稳定可扩展
- 支付网关接口文档完整且可对接
- WebSocket基础设施支持水平扩展
- 翻译资源可按时交付

## Acceptance Criteria

### AC-1: 印尼市场语言切换
- **Given**: 用户访问平台
- **When**: 用户选择印尼语 (Bahasa Indonesia)
- **Then**: 所有页面文字显示为印尼语，包括UI和消息
- **Verification**: `human-judgment`

### AC-2: 印尼盾价格显示
- **Given**: 商品在印尼市场展示
- **When**: 查看商品详情页
- **Then**: 价格以IDR格式显示 (Rp X.XXX.XXX)
- **Verification**: `programmatic`

### AC-3: 美元价格显示
- **Given**: 商品在美国市场展示
- **When**: 查看商品详情页
- **Then**: 价格以USD格式显示 ($X,XXX.XX)
- **Verification**: `programmatic`

### AC-4: 拍卖商品创建
- **Given**: 商家登录并发布商品
- **When**: 选择"拍卖商品"类型，设置起拍价和时长
- **Then**: 拍卖商品创建成功，状态为"未开始"
- **Verification**: `programmatic`

### AC-5: 实时出价
- **Given**: 买家进入拍卖详情页（拍卖中）
- **When**: 买家出价，金额高于当前最高价+加价幅度
- **Then**: 出价成功，所有在线用户看到最新出价和最高出价者
- **Verification**: `programmatic`

### AC-6: 预售定金支付
- **Given**: 买家购买预售商品
- **When**: 下单并支付定金
- **Then**: 订单状态为"已付定金"，库存锁定，尾款支付时间显示
- **Verification**: `programmatic`

### AC-7: 预售尾款支付
- **Given**: 买家有已付定金的预售订单
- **When**: 在尾款支付时间内完成尾款支付
- **Then**: 订单状态变为"已完成"
- **Verification**: `programmatic`

## Open Questions
- [ ] 直播功能是必须项还是可选项？
- [ ] 印尼和美国市场是否需要独立的商家审核流程？
