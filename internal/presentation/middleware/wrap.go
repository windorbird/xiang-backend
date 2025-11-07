package middleware

import (
    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/helper"
)

func Wrap(h func(*helper.Context)) gin.HandlerFunc {
    return func(c *gin.Context) {
        h(&helper.Context{c})
    }
}
