package helper

import (
    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/common"
)

type Context struct {
    *gin.Context
}

func (c *Context) Success(data interface{}) {
    c.Set("response", common.Response{
        Code: 0,
        Msg:  "success",
        Data: data,
    })
}

func (c *Context) Fail(code int, msg string) {
    c.Set("response", common.Response{
        Code: code,
        Msg:  msg,
    })
}

func (c *Context) FailErr(err error) {
    // 如果是业务错误 AppError
    if e, ok := err.(*common.AppError); ok {
        c.Fail(e.Code, e.Msg)
        return
    }

    // 普通 error
    c.Fail(10001, err.Error())
}
