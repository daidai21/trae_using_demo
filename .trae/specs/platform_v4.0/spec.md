# 电商交易平台 - 平台化架构升级 v4.0 - 产品需求文档 (PRD)

## Overview
- **Summary**: 将当前业务增长架构升级为平台化架构，通过SPI扩展机制实现平台与业务的分离，平台仅保留普通交易模式，ID市场、US市场、预售、直播拍卖等功能通过SPI扩展实现。
- **Purpose**: 实现平台核心与业务扩展的解耦，提高系统可扩展性和可维护性。
- **Target Users**: 后端开发团队、架构师。

**注：本次仅进行后端架构升级，前端保持不变。**

## Goals
1. 抽象业务身份：`{国家}.{玩法}` 标识体系
2. 在 common 包中新增 SPI 扩展框架
3. 平台核心服务仅保留普通交易模式
4. ID市场、US市场、预售、直播拍卖通过 SPI 扩展实现
5. API 网关支持业务身份解析和路由

## Non-Goals (Out of Scope)
- 不涉及前端架构改造
- 不涉及数据库 schema 重构
- 不涉及现有 API 接口变更（保持向后兼容）
- 不涉及性能优化（后续迭代）

## Background & Context
- 现有架构：微服务（api-gateway, user-service, product-service, trade-service, auction-service）+ Monorepo
- 业务功能已全部实现（ID/US市场、预售、直播拍卖）
- 现有代码耦合度较高，平台与业务逻辑混杂
- 需要通过 SPI 机制实现平台与业务的分离

## Functional Requirements

### FR-1: 业务身份抽象 (Business Identity)
- 定义 `{国家}.{玩法}` 的业务身份标识
- 支持的国家代码：CN, ID, US
- 支持的玩法模式：normal, pre_sale, auction
- 身份示例：CN.normal, ID.normal, US.normal, CN.preSale, CN.auction
- 提供 BusinessIdentity 数据结构
- 提供身份解析器（从请求、用户偏好解析）

### FR-2: SPI 扩展框架 (SPI Framework)
- 定义 ExtensionPoint 基础接口
- 定义各领域扩展点接口：
  - ProductExtension（商品扩展）
  - TradeExtension（交易扩展）
  - I18nExtension（国际化扩展）
  - PaymentExtension（支付扩展）
- 实现 ExtensionRegistry 扩展注册表
- 支持扩展优先级机制
- 支持扩展按业务身份分组

### FR-3: 平台核心服务重构 (Platform Core Services)
- 创建 platform/ 目录
- 迁移 user-service → platform/user-service
- 迁移 product-service → platform/product-service（移除预售、拍卖代码）
- 迁移 trade-service → platform/trade-service（移除预售、多市场代码）
- 平台服务仅保留普通交易模式
- 在关键路径添加 SPI 扩展点调用

### FR-4: 业务扩展模块 (Business Extensions)
- 创建 extensions/ 目录
- id-market 扩展：印尼语语言包、印尼支付网关
- us-market 扩展：美式英语、美国支付网关、Sales Tax
- pre-sale 扩展：预售商品、预售订单、定金/尾款
- auction 扩展：拍卖服务（从 auction-service 迁移）

### FR-5: API 网关适配 (API Gateway Adaptation)
- 集成业务身份解析
- 根据业务身份路由到相应扩展服务
- 保持向后兼容

## Non-Functional Requirements
- **NFR-1**: SPI 框架代码覆盖率 > 80%
- **NFR-2**: 扩展加载性能损耗 < 5%
- **NFR-3**: 平台核心代码无业务逻辑
- **NFR-4**: 所有现有 API 保持兼容

## Constraints
- **Technical**: 技术栈不变（Go, Hertz, GORM, React, Ant Design）
- **Business**: 仅后端改造，前端不变
- **Dependencies**: 依赖现有业务功能代码

## Assumptions
- 现有业务功能代码完整可用
- 无需变更数据库 schema
- 无需变更现有 API 接口

## Acceptance Criteria

### AC-1: 业务身份解析
- **Given**: 请求包含业务身份标识
- **When**: 系统解析身份
- **Then**: 正确解析为 BusinessIdentity 对象
- **Verification**: `programmatic`

### AC-2: SPI 扩展注册
- **Given**: 扩展模块实现 ExtensionPoint
- **When**: 系统启动时注册扩展
- **Then**: 扩展成功注册到 ExtensionRegistry
- **Verification**: `programmatic`

### AC-3: 平台核心服务迁移
- **Given**: platform/ 目录已创建
- **When**: 查看平台服务代码
- **Then**: 仅包含普通交易模式，无业务扩展代码
- **Verification**: `human-judgment`

### AC-4: 业务扩展模块创建
- **Given**: extensions/ 目录已创建
- **When**: 查看扩展模块
- **Then**: 包含 id-market, us-market, pre-sale, auction 四个扩展
- **Verification**: `human-judgment`

### AC-5: 向后兼容
- **Given**: 现有 API 调用
- **When**: 使用原有方式调用 API
- **Then**: 功能正常，响应不变
- **Verification**: `programmatic`

## Open Questions
- [ ] 扩展加载策略：编译时加载还是运行时发现？
- [ ] 是否需要扩展热加载机制？
