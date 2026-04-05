# 电商交易平台架构升级 V1.0 - 验证清单

## 后端验证

- [ ] backend-legacy/和frontend-legacy/目录存在且代码完整
- [ ] domain-services/目录结构符合规范
- [ ] domain-services/common/go/pkg/包含response和utils
- [ ] domain-services/common/go.mod模块名正确
- [ ] user-service可以go build无错误
- [ ] user-service可以在8081端口启动
- [ ] user-service的/auth/login和/auth/register接口正常
- [ ] product-service可以go build无错误
- [ ] product-service可以在8082端口启动
- [ ] product-service的/products和/merchants接口正常
- [ ] trade-service可以go build无错误
- [ ] trade-service可以在8083端口启动
- [ ] trade-service的/cart和/orders接口正常
- [ ] trade-service内部按cart/buy/order子域组织
- [ ] api-gateway可以go build无错误
- [ ] api-gateway可以在8080端口启动
- [ ] 通过api-gateway可以访问所有服务接口
- [ ] 所有后端服务可以同时启动

## 前端验证

- [ ] pnpm-workspace.yaml配置正确
- [ ] 根package.json配置正确
- [ ] pnpm install可以正常执行
- [ ] packages/ui目录结构完整
- [ ] packages/ui可以正常构建
- [ ] packages/user目录结构完整
- [ ] packages/user可以正常构建
- [ ] packages/user对应用户服务API
- [ ] packages/product目录结构完整
- [ ] packages/product可以正常构建
- [ ] packages/product对应商家商品服务API
- [ ] packages/trade目录结构完整
- [ ] packages/trade可以正常构建
- [ ] packages/trade内部按cart/buy/order子域组织
- [ ] packages/trade对应交易服务API
- [ ] apps/web目录结构完整
- [ ] apps/web可以pnpm dev启动
- [ ] apps/web路由配置正确
- [ ] apps/web整合了所有packages模块

## 功能验证

- [ ] 用户可以正常注册
- [ ] 用户可以正常登录
- [ ] 可以浏览商品列表
- [ ] 可以查看商品详情
- [ ] 可以浏览商家列表
- [ ] 可以查看商家详情
- [ ] 可以添加商品到购物车
- [ ] 可以修改购物车商品数量
- [ ] 可以删除购物车商品
- [ ] 可以创建订单
- [ ] 可以查看订单列表
- [ ] 可以查看订单详情
- [ ] 可以更新订单状态

## 文档验证

- [ ] README文档已更新
- [ ] 提供了后端启动脚本
- [ ] 提供了前端启动脚本
- [ ] 架构升级指南已编写
- [ ] 目录结构与文档一致

