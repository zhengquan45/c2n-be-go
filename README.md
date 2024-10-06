# C2E-BE-GO

用 Golang 语言主流的 Web 框架 Gin 重写 c2e-be, 故而工程以 c2e-be-go 命名。

📘 [Gin 文档传送门 ](https://gin-gonic.com/zh-cn/)

📘 [Gorm 文档传送门](https://gorm.io/docs/)

📘 [Geth 官方文档]([https://geth.ethereum.org/docs/]())

.
├── api
│   ├── handler
│   │   └── user.go        # 用户相关的业务逻辑处理函数
│   ├── request
│   │   └── user.go        # 请求参数的定义和验证
│   ├── response
│       └── user.go        # 响应数据的定义
├── config
│   └── config.go          # 配置文件的加载、管理
├── database
│   └── database.go        # 数据库的初始化和连接管理
├── docs                   # 项目文档或接口文档
├── middleware
│   └── auth.go            # 中间件，如认证、日志等
├── model
│   └── user.go            # GORM 模型定义
├── repository
│   └── user_repository.go # 数据库访问逻辑，封装数据库操作
├── routes
│   └── router.go          # 路由定义
├── service
│   └── user_service.go    # 业务逻辑层，调用 repository 处理业务
├── main.go                # 主入口文件
├── go.mod                 # Go modules 管理文件
└── go.sum                 # Go modules 依赖文件
