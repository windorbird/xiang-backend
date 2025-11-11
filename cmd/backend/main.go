package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
    "github.com/windorbird/xiang-backend/internal/presentation/api/router"
    "github.com/windorbird/xiang-backend/internal/presentation/middleware"
    "github.com/windorbird/xiang-backend/log"
)

func init() {
    log.InitLogger()
}

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Logger.Fatal("Error loading .env file")
    }

    r := gin.New()
    r.Use(gin.Recovery()) // Gin 自带 recovery，用 zap 再兜底
    r.Use(middleware.Logger())
    r.Use(middleware.BindMiddleware())
    r.Use(middleware.ResponseMiddleware())

    redis.InitRedis()

    if err := r.SetTrustedProxies([]string{"127.0.0.1", "::1"}); err != nil {
        fmt.Printf("SetTrustedProxies error: %v", err)
    }

    // 初始化抖音小程序路由
    router.InitRouter(r)

    // 启动服务器
    err = r.Run("127.0.0.1:8080")
    if err != nil {
        panic(err)
    }
}
