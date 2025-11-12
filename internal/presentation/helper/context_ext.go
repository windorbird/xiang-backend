package helper

import "github.com/gin-gonic/gin"

// Success 扩展方法
func Success(c *gin.Context, data interface{}) {
    if fn, ok := c.Get("Success"); ok {
        fn.(func(interface{}))(data)
    }
}

// Fail 扩展方法
func Fail(c *gin.Context, code int, msg string) {
    if fn, ok := c.Get("Fail"); ok {
        fn.(func(int, string))(code, msg)
    }
}
