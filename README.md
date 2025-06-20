# Koreader Sync Server (Golang Gin 版)

本目录为用 Golang + Gin 重写的 KOReader 同步服务。

## 主要接口
- POST   /users/create           创建用户
- GET    /users/auth             用户认证（需 x-auth-user, x-auth-key）
- PUT    /syncs/progress         更新文档进度（需认证）
- GET    /syncs/progress/:document 获取文档进度（需认证）
- GET    /healthcheck            健康检查

## 依赖
- Go 1.20+
- sqlite
- github.com/gin-gonic/gin

## 启动
1. 启动 Redis 服务（默认 6379，DB=1）
2. 进入本目录，执行：
   ```
   go mod tidy
   go run main.go
   ```

## 认证说明
- 用户注册后，客户端需用 MD5(password) 作为 key 传递。
- 所有同步接口需在 Header 中带 x-auth-user 和 x-auth-key。 
