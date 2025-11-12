# xiang-backend
香后端，go实现


项目组织如下
```
my-app/
├── cmd/
│   └── my-app/
│       └── main.go                 # 程序入口，组装所有组件
├── internal/
│   ├── presentation/               # 表现层（原 handler）
│   │   ├── http/
│   │   │   ├── server.go           # Gin 路由注册
│   │   │   └── handler/
│   │   │       └── user_handler.go
│   │   └── middleware/             # HTTP 中间件
│   │
│   ├── application/                # 应用层（Service）
│   │   ├── service/
│   │   │   └── user_service.go
│   │   └── dto/                    # 应用层 DTO（可选）
│   │
│   ├── domain/                     # 领域层（核心）
│   │   ├── model/
│   │   │   └── user.go             # Entity
│   │   ├── repository/             # 接口定义
│   │   │   └── user_repository.go
│   │   └── errors.go               # 自定义错误
│   │
│   └── infrastructure/             # 基础设施层
│       ├── persistence/
│       │   └── gorm/
│       │       ├── gorm_user_repo.go
│       │       └── db.go           # DB 初始化
│       ├── external/
│       │   └── sendgrid_email.go   # 第三方服务
│       └── config/                 # 配置加载
│
├── pkg/                            # 可公开复用的库（谨慎使用）
├── migrations/                     # 数据库迁移
├── scripts/
└── go.mod
```






