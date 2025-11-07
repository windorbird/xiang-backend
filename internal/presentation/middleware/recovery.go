package middleware

import (
    "log"
    "runtime/debug"

    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/common"
)

func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if r := recover(); r != nil {
                log.Println("[PANIC]", r)
                log.Println(string(debug.Stack()))

                c.JSON(200, common.Response{
                    Code: 50000,
                    Msg:  "服务器内部错误",
                })
                c.Abort()
            }
        }()

        c.Next()
    }
}
