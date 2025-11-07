package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/common"
)

func BindMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        for _, err := range c.Errors {
            if err.Type == gin.ErrorTypeBind {
                // 参数错误
                c.JSON(200, common.Response{
                    Code: 10002,
                    Msg:  "参数不合法",
                })
                c.Abort()
                return
            }
        }
    }
}
