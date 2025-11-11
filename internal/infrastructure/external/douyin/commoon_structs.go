package douyin

type RequestBody struct {
	Tasks []Task `json:"tasks"` // 检测任务列表
}

// Task 单个检测任务结构
type Task struct {
	Content string `json:"content"` // 待检测文本内容
}

// Response 2. 定义响应相关结构体
// Response 接口完整响应结构
type Response struct {
	LogID string         `json:"log_id"` // 请求唯一ID
	Data  []ResponseData `json:"data"`   // 检测结果列表
}

// ResponseData 单个检测任务的结果结构
type ResponseData struct {
	Msg      string    `json:"msg"`      // 结果消息
	Code     int       `json:"code"`     // 状态码
	TaskID   string    `json:"task_id"`  // 任务ID
	Predicts []Predict `json:"predicts"` // 置信度列表
	DataID   *string   `json:"data_id"`  // 数据ID（可能为null）
}

// Predict 置信度详情结构
type Predict struct {
	Prob      float64 `json:"prob"`       // 概率（仅供参考）
	Hit       bool    `json:"hit"`        // 是否包含违规内容（true=违规）
	Target    *string `json:"target"`     // 服务/目标（可能为null）
	ModelName string  `json:"model_name"` // 模型/标签
}

// ErrorResponse 3. 错误响应结构体（请求失败时返回）
type ErrorResponse struct {
	ErrorID   string `json:"error_id"`  // 错误ID
	Code      int    `json:"code"`      // 错误码
	Message   string `json:"message"`   // 错误信息
	Exception string `json:"exception"` // 异常详情
}
