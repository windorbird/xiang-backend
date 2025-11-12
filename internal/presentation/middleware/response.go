// package middleware
//
// import (
//
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/windorbird/xiang-backend/common"
//
// )
//
//	func ResponseMiddleware() gin.HandlerFunc {
//		return func(c *gin.Context) {
//
//			c.Next()
//
//			// 已写 JSON 则跳过
//			if c.Writer.Written() {
//				return
//			}
//
//			// 处理 gin 的 c.Error(err)
//			if len(c.Errors) > 0 {
//				c.JSON(http.StatusOK, common.Response{
//					Code: 50000,
//					Msg:  c.Errors[0].Error(),
//				})
//				return
//			}
//
//			// 中间件或成功/失败设置的统一 response
//			if resp, ok := c.Get("response"); ok {
//				c.JSON(http.StatusOK, resp)
//				return
//			}
//
//			// 默认成功
//			c.JSON(http.StatusOK, common.Response{
//				Code: 0,
//				Msg:  "success",
//			})
//		}
//	}
//
// -----------------------------------------------------
package middleware

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/common"
)

func ResponseMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        // 为 Context 挂载 Success / Fail 方法
        c.Set("Success", func(data interface{}) {
            c.Set("response", common.Response{
                Code: 0,
                Msg:  "success",
                Data: data,
            })
        })

        c.Set("Fail", func(code int, msg string) {
            c.Set("response", common.Response{
                Code: code,
                Msg:  msg,
            })
        })

        // 执行 Handler
        c.Next()

        // Handler 已写过输出就不处理
        if c.Writer.Written() {
            return
        }

        // 如果捕获到了 panic / c.Error(err)
        if len(c.Errors) > 0 {
            c.JSON(http.StatusOK, common.Response{
                Code: 50000,
                Msg:  c.Errors[0].Error(),
            })
            return
        }

        // 取响应内容
        if resp, exists := c.Get("response"); exists {
            c.JSON(http.StatusOK, resp)
            return
        }

        // 什么都没写：默认成功
        c.JSON(http.StatusOK, common.Response{
            Code: 0,
            Msg:  "success",
        })
    }
}
