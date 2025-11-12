package main

import (
    "context"
    "errors"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "github.com/windorbird/xiang-backend/internal/infrastructure/persistence/redis"
    "github.com/windorbird/xiang-backend/internal/presentation/api/router"
    "github.com/windorbird/xiang-backend/internal/presentation/middleware"
    "github.com/windorbird/xiang-backend/log"
    "go.uber.org/zap"
)

func init() {

}

func main() {
    log.InitLogger()
    defer log.Logger.Sync()

    if err := godotenv.Load(".env"); err != nil {
        log.Logger.Fatal("Failed to load .env file", zap.Error(err))
    }


    r := gin.New()
    r.Use(gin.Recovery()) // Gin 自带 recovery，用 zap 再兜底
    r.Use(middleware.Logger())
    r.Use(middleware.BindMiddleware())
    r.Use(middleware.ResponseMiddleware())

    if err := redis.InitRedis(); err != nil {
        log.Logger.Fatal("Redis initialization failed", zap.Error(err))
    }

    if err := r.SetTrustedProxies([]string{"127.0.0.1", "::1"}); err != nil {
        log.Logger.Fatal("Failed to set trusted proxies", zap.Error(err))
    }

    // 初始化抖音小程序路由
    router.InitRouter(r)

    // 启动服务器
    // 创建 HTTP 服务器，包装 Gin
    server := &http.Server{
        Addr:    "127.0.0.1:8080",
        Handler: r,
    }

    // 启动服务
    go func() {
        if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
            log.Logger.Error("Server error", zap.Error(err))
        }
    }()

    // 信号监听
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Logger.Info("Shutting down server gracefully...")

    // 创建超时上下文，等待10秒
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // 优雅关闭
    if err := server.Shutdown(ctx); err != nil {
        log.Logger.Error("Server shutdown error", zap.Error(err))
    }

    log.Logger.Info("Server stopped gracefully")
}
