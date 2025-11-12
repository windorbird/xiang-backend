package ecode

//| 状态码     | 名称                  | 场景说明               |
//| ------- | ------------------- | ------------------ |
//| **0**   | OK                  | 成功                 |
//| **400** | BadRequest          | 参数缺失/格式错误          |
//| **401** | Unauthorized        | 未登录或 Token 无效 / 过期 |
//| **403** | Forbidden           | 无权限执行              |
//| **404** | NotFound            | 资源不存在              |
//| **409** | Conflict            | 资源冲突，如重复注册、重复提交    |
//| **429** | TooManyRequests     | 请求限流               |
//| **500** | InternalServerError | 服务器错误（不要暴露内部细节）    |
//| **503** | ServiceUnavailable  | 服务不可用/限流/熔断        |

//建议全部 1xxx 用于业务错，2000+ 用于第三方、外部依赖错误

//| 状态码      | 名称                | 场景                |
//| -------- | ----------------- | ----------------- |
//| **1001** | InvalidParam      | 参数校验失败（字段非法、格式错误） |
//| **1002** | NeedLogin         | 需要重新登录（token过期）   |
//| **1003** | DuplicateAction   | 重复请求 / 幂等校验失败     |
//| **1004** | ResourceExisted   | 数据已存在（例如手机号已注册）   |
//| **1005** | ResourceNotEnough | 余额不足、库存不够         |

//| 状态码      | 名称                | 场景                  |
//| -------- | ----------------- | ------------------- |
//| **2001** | DBError           | 数据库错误               |
//| **2002** | RedisError        | 缓存异常                |
//| **2003** | MQError           | 队列异常                |
//| **2004** | ThirdPartyTimeout | 调用第三方接口超时           |
//| **2005** | ThirdPartyError   | 依赖第三方错误（短信、支付、云存储等） |

type Code int

const (
	OK Code = 0

	BadRequest          Code = 400
	Unauthorized        Code = 401
	Forbidden           Code = 403
	NotFound            Code = 404
	Conflict            Code = 409
	TooManyRequests     Code = 429
	InternalServerError Code = 500
)

const (
	InvalidParam      Code = 1001
	NeedLogin         Code = 1002
	DuplicateAction   Code = 1003
	ResourceExisted   Code = 1004
	ResourceNotEnough Code = 1005
)
