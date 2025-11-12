package douyin

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "os"
)
type DetectResp struct {
    IsDanger bool `json:"is_danger"`
}


func DetectContent(accessToken string, content string) (*DetectResp, error) {
    // 4.0 检查参数
    if len(accessToken) == 0 {
        return nil, errors.New("accessToken is empty") // 如果accessToken为空，返回错误
    }
    if len(content) == 0 {
        return nil, errors.New("content is empty") // 如果content为空，返回错误
    }

    // 4.1 确定请求地址
    var reqURL = os.Getenv("DouYinCheckTextURL") // 从环境变量中获取抖音文本检测接口地址
    if len(reqURL) == 0 {
      return nil, errors.New("DouYinCheckTextURL is empty") // 如果环境变量为空，返回错误
    }

    // 4.2 构建请求体

    // 创建请求体结构，包含要检测的内容任务
    reqBody := RequestBody{
        Tasks: []Task{
            {Content: content}, // 将需要检测的内容放入任务列表
        },
    }
    // 将请求体序列化为JSON格式
    reqBodyJSON, err := json.Marshal(reqBody)
    if err != nil {
        return nil, errors.New("request body marshal failed, body:" + string(reqBodyJSON)) // 序列化失败返回错误
    }

    // 4.3 创建HTTP请求
    req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(reqBodyJSON))
    if err != nil {
        return nil, errors.New("create request failed, err:" + err.Error())
    }

    // 4.4 设置请求头（必填）
    // 创建POST请求，指定请求URL和请求体
    req.Header.Set("X-Token", accessToken)
    req.Header.Set("Content-Type", "application/json") // 接口强制要求的Content-Type
 // 请求创建失败返回错误
    // 4.5 发送请求并获取响应
    client := &http.Client{}
    resp, err := client.Do(req)
    // 设置访问令牌请求头
    if err != nil {
        return nil, errors.New("send request failed, err:" + err.Error())
    }
    defer resp.Body.Close() // 确保响应体被关闭，避免资源泄漏

    // 4.6 处理响应（区分成功/失败）
    // 读取响应体内容
    var respBody []byte
    if respBody, err = io.ReadAll(resp.Body); err != nil {
        return nil, fmt.Errorf("read response body failed, err:%s, body:%s", err.Error(), string(respBody))
    }

    // 先尝试解析为成功响应
    var successResp Response
    err = json.Unmarshal(respBody, &successResp)
    if err != nil{
        return nil, fmt.Errorf("response body unmarshal failed, err:%s, body:%s", err.Error(), string(respBody))
    }
    if len(successResp.Data) == 0 {
        return nil, fmt.Errorf("response data is empty, body:%s", string(respBody))
    }
    if successResp.Data[0].Code != 0 {
        return nil, fmt.Errorf("check text failed, taskID:%s, code:%d, msg:%s", successResp.Data[0].TaskID, successResp.Data[0].Code, successResp.Data[0].Msg)
    }
    if len(successResp.Data[0].Predicts) == 0 {
        return nil, fmt.Errorf("response predicts is empty, body:%s", string(respBody))
    }
    return &DetectResp{IsDanger: successResp.Data[0].Predicts[0].Hit}, nil
}