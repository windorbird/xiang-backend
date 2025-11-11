package router

import (
    "github.com/gin-gonic/gin"
    "github.com/windorbird/xiang-backend/internal/presentation/api/handler/douyin"
    "github.com/windorbird/xiang-backend/internal/presentation/middleware"
    //"github.com/windorbird/xiang-backend/services/douyin"
)

//r.GET("/user", wrap(func(c *helper.Context) {
//    c.Success(map[string]interface{}{
//        "id":   1,
//        "name": "Tom",
//    })
//}))
//
//r.GET("/error", wrap(func(c *helper.Context) {
//    c.Fail(10001, "账号密码错误")
//}))
//
//r.GET("/panic", wrap(func(c *helper.Context) {
//    panic("爆炸")
//}))

// InitDouyinRouter 初始化抖音小程序路由组
func InitDouyinRouter(api *gin.RouterGroup) {
    douyinGroup := api.Group("/douyin")
    {
        douyinGroup.GET("/login", middleware.Wrap(douyin.LoginHandler))
        douyinGroup.POST("/check-text", middleware.Wrap(douyin.CheckText))
    }
    // 可以在这里添加更多抖音小程序的路由
    //douyinGroup.GET("/user/info", getUserInfo)
    //douyinGroup.POST("/upload", uploadHandler)
}
