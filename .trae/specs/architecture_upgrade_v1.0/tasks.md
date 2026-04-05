# 电商交易平台架构升级 V.0 - 实施计划

## \[ ] Task 1: 备份原有代码

- **Priority**: P0
- **Depends On**: None
- **Description**:
  - 将原有的backend/重命名为backend-legacy/
  - 将原有的frontend/重命名为frontend-legacy/
  - 确保原有代码完整保留
- **Acceptance Criteria Addressed**: \[AC-5]
- **Test Requirements**:
  - `programmatic` TR-1.1: 验证backend-legacy/和frontend-legacy/目录存在
  - `human-judgement` TR-1.2: 检查原有代码文件完整性
- **Notes**: 先备份，避免代码丢失

## \[ ] Task 2: 创建后端微服务目录结构

- **Priority**: P0
- **Depends On**: Task 1
- **Description**:
  - 创建domain-services/目录及其子目录
  - 创建所有服务的基础目录结构（cmd, internal, pkg等）
- **Acceptance Criteria Addressed**: \[AC-1, AC-5]
- **Test Requirements**:
  - `programmatic` TR-2.1: 验证所有服务目录结构存在
  - `human-judgement` TR-2.2: 检查目录结构符合规范
- **Notes**: 目录结构参照architecture\_upgrade\_v1.0\_plan.md

## \[ ] Task 3: 创建domain-services/common共享代码库

- **Priority**: P0
- **Depends On**: Task 2
- **Description**:
  - 从原backend/pkg/复制response和utils到domain-services/common/go/pkg/
  - 创建domain-services/common/go/go.mod
  - 确保共享代码可正常编译
- **Acceptance Criteria Addressed**: \[AC-1, AC-5]
- **Test Requirements**:
  - `programmatic` TR-3.1: 验证shared代码可以go build
  - `human-judgement` TR-3.2: 检查代码无编译错误
- **Notes**: 模块名使用ecommerce/common

## \[ ] Task 4: 拆分并实现user-service用户服务

- **Priority**: P0
- **Depends On**: Task 3
- **Description**:
  - 从原backend拆分用户相关代码到user-service
  - 创建用户服务的model、repository、service、handler
  - 创建main.go入口
  - 配置go.mod引用common
  - 确保服务可独立启动（端口8081）
- **Acceptance Criteria Addressed**: \[AC-1, AC-2, AC-4]
- **Test Requirements**:
  - `programmatic` TR-4.1: 验证user-service可以go build
  - `programmatic` TR-4.2: 验证服务可以在8081端口启动
  - `programmatic` TR-4.3: 验证/auth/login和/auth/register接口正常
- **Notes**: 数据库使用users.db

## \[ ] Task 5: 拆分并实现product-service商家商品服务

- **Priority**: P0
- **Depends On**: Task 3
- **Description**:
  - 从原backend拆分商家和商品相关代码到product-service
  - 创建服务的model、repository、service、handler
  - 创建main.go入口
  - 配置go.mod引用common
  - 确保服务可独立启动（端口8082）
- **Acceptance Criteria Addressed**: \[AC-1, AC-2, AC-4]
- **Test Requirements**:
  - `programmatic` TR-5.1: 验证product-service可以go build
  - `programmatic` TR-5.2: 验证服务可以在8082端口启动
  - `programmatic` TR-5.3: 验证/products和/merchants接口正常
- **Notes**: 数据库使用product.db

## \[ ] Task 6: 拆分并实现trade-service交易服务

- **Priority**: P0
- **Depends On**: Task 3
- **Description**:
  - 从原backend拆分购物车和订单相关代码到trade-service
  - 按子域组织代码：cart/, buy/, order/
  - 创建服务的model、repository、service、handler
  - 创建main.go入口
  - 配置go.mod引用common
  - 确保服务可独立启动（端口8083）
- **Acceptance Criteria Addressed**: \[AC-1, AC-2, AC-4]
- **Test Requirements**:
  - `programmatic` TR-6.1: 验证trade-service可以go build
  - `programmatic` TR-6.2: 验证服务可以在8083端口启动
  - `programmatic` TR-6.3: 验证/cart和/orders接口正常
- **Notes**: 数据库使用order.db

## \[ ] Task 7: 实现api-gateway API网关

- **Priority**: P0
- **Depends On**: Task 4, Task 5, Task 6
- **Description**:
  - 创建API网关的反向代理逻辑
  - 实现路由转发规则
  - 实现JWT认证中间件
  - 创建main.go入口
  - 配置go.mod引用common
  - 确保网关可独立启动（端口8080）
- **Acceptance Criteria Addressed**: \[AC-1, AC-2, AC-4]
- **Test Requirements**:
  - `programmatic` TR-7.1: 验证api-gateway可以go build
  - `programmatic` TR-7.2: 验证网关可以在8080端口启动
  - `programmatic` TR-7.3: 验证通过网关可以访问所有服务接口
- **Notes**: 路由规则参照architecture\_upgrade\_v1.0\_plan.md

## \[ ] Task 8: 配置前端pnpm workspaces

- **Priority**: P0
- **Depends On**: Task 1
- **Description**:
  - 创建pnpm-workspace.yaml
  - 创建根package.json
  - 配置workspaces
- **Acceptance Criteria Addressed**: \[AC-3, AC-5]
- **Test Requirements**:
  - `programmatic` TR-8.1: 验证pnpm install可以正常执行
  - `human-judgement` TR-8.2: 检查workspace配置正确

## \[ ] Task 9: 创建packages/ui共享组件库

- **Priority**: P0
- **Depends On**: Task 8
- **Description**:
  - 从原frontend拆分共享组件到packages/ui
  - 创建components/, hooks/, utils/目录
  - 创建package.json
  - 导出共享组件和工具
- **Acceptance Criteria Addressed**: \[AC-3, AC-5]
- **Test Requirements**:
  - `human-judgement` TR-9.1: 检查共享组件完整
  - `programmatic` TR-9.2: 验证package可以正常构建

## \[ ] Task 10: 拆分packages/user用户模块

- **Priority**: P0
- **Depends On**: Task 9
- **Description**:
  - 从原frontend拆分用户相关代码到packages/user
  - 创建components/, hooks/, services/, types/目录
  - 创建package.json，依赖ui包
  - 导出用户相关组件和服务
- **Acceptance Criteria Addressed**: \[AC-3, AC-4, AC-5]
- **Test Requirements**:
  - `human-judgement` TR-10.1: 检查用户模块完整
  - `programmatic` TR-10.2: 验证package可以正常构建
- **Notes**: 对应用户服务

## \[ ] Task 11: 拆分packages/product商品模块

- **Priority**: P0
- **Depends On**: Task 9
- **Description**:
  - 从原frontend拆分商品和商家相关代码到packages/product
  - 创建components/, hooks/, services/, types/目录
  - 创建package.json，依赖ui包
  - 导出商品和商家相关组件和服务
- **Acceptance Criteria Addressed**: \[AC-3, AC-4, AC-5]
- **Test Requirements**:
  - `human-judgement` TR-11.1: 检查商品模块完整
  - `programmatic` TR-11.2: 验证package可以正常构建
- **Notes**: 对应商家商品服务

## \[ ] Task 12: 拆分packages/trade交易模块

- **Priority**: P0
- **Depends On**: Task 9
- **Description**:
  - 从原frontend拆分购物车和订单相关代码到packages/trade
  - 按子域组织代码：cart/, buy/, order/
  - 创建components/, hooks/, services/, types/目录
  - 创建package.json，依赖ui包
  - 导出交易相关组件和服务
- **Acceptance Criteria Addressed**: \[AC-3, AC-4, AC-5]
- **Test Requirements**:
  - `human-judgement` TR-12.1: 检查交易模块完整
  - `programmatic` TR-12.2: 验证package可以正常构建
- **Notes**: 对应交易服务

## \[ ] Task 13: 整合apps/web主应用

- **Priority**: P0
- **Depends On**: Task 10, Task 11, Task 12
- **Description**:
  - 创建apps/web应用
  - 配置路由
  - 整合所有packages模块
  - 创建App.tsx和main.tsx
  - 配置vite.config.ts
  - 创建package.json，依赖所有packages
- **Acceptance Criteria Addressed**: \[AC-3, AC-4, AC-5]
- **Test Requirements**:
  - `programmatic` TR-13.1: 验证可以pnpm dev启动
  - `human-judgement` TR-13.2: 检查所有页面可以正常访问

## \[ ] Task 14: 后端联调测试

- **Priority**: P1
- **Depends On**: Task 7
- **Description**:
  - 启动所有后端微服务
  - 通过API网关测试所有接口
  - 验证功能完整性
- **Acceptance Criteria Addressed**: \[AC-2, AC-4]
- **Test Requirements**:
  - `programmatic` TR-14.1: 所有API接口正常响应
  - `human-judgement` TR-14.2: 功能完整可用

## \[ ] Task 15: 前端联调测试

- **Priority**: P1
- **Depends On**: Task 13
- **Description**:
  - 启动前端应用
  - 测试所有页面和功能
  - 验证与后端API集成正常
- **Acceptance Criteria Addressed**: \[AC-4]
- **Test Requirements**:
  - `human-judgement` TR-15.1: 所有页面正常显示
  - `human-judgement` TR-15.2: 所有功能正常工作

## \[ ] Task 16: 整体联调与文档更新

- **Priority**: P1
- **Depends On**: Task 14, Task 15
- **Description**:
  - 前后端整体联调
  - 更新README文档
  - 编写启动脚本
  - 编写架构升级指南
- **Acceptance Criteria Addressed**: \[AC-4, AC-5]
- **Test Requirements**:
  - `human-judgement` TR-16.1: 文档完整清晰
  - `human-judgement` TR-16.2: 可以按照文档启动项目

