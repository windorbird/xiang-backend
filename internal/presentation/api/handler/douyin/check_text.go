package douyin

import (
    externalDouyin "github.com/windorbird/xiang-backend/internal/infrastructure/external/douyin"
    "github.com/windorbird/xiang-backend/internal/presentation/helper"
)


type CheckTextReq struct {
    Text string `json:"text" binding:"required"`
}

func CheckText(ctx *helper.Context)  {

    var req CheckTextReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.FailErr(err)
        return
    }
    accessToken, err := externalDouyin.GetAccessToken(ctx.Request.Context())
    if err != nil {
        ctx.FailErr(err)
        return
    }
    // 调用抖音API检查文本
    resp, err := externalDouyin.DetectContent(accessToken,  req.Text)
    if err != nil {
        ctx.FailErr(err)
        return
    }
    ctx.Success(resp)
}
