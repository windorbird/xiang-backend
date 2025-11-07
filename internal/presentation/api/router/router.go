package router

import (
    "github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
    api := router.Group("/api")

    // 初始化抖音小程序路由
    InitDouyinRouter(api)
}
