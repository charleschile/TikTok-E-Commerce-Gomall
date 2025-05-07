

# TikTok E-Commerce：Gomall





### 00项目总体介绍

Kitex：go的rpc框架

Hertz：go的http框架

volo：rust rpc框架



#### **技术栈**：

#### **框架与库**
- **Kitex**：字节跳动开源的高性能 RPC 框架，适用于微服务通信。
- **Hertz**：用于构建高性能 Web 服务的 HTTP 框架。
- **cwgo**：Go 的代码生成工具，通常用于生成服务模板或接口实现。
- **GORM**：Go 的 ORM（对象关系映射）库，简化数据库操作。

#### **数据库与缓存**
- **MySQL**：关系型数据库，用于持久化存储结构化数据。
- **Redis**：高性能内存数据库，常用于缓存、分布式锁等场景。

#### **消息系统与服务发现**
- **NATS**：轻量级消息中间件，支持高效的消息发布/订阅。
- **Consul**：服务发现和配置中心，支持微服务架构的服务注册与健康检查。

#### **监控与追踪**
- **Prometheus**：指标监控系统，支持服务性能监控与告警。
- **OpenTelemetry**：用于分布式链路追踪和度量采集，分析系统瓶颈。

#### **容器与编排**
- **Docker**：容器化技术，使应用能够在任意环境一致运行。
- **K8s（Kubernetes）**：容器编排平台，用于自动部署、扩展和管理容器应用。

#### **项目结构**


![design](img/design.png)

### 01开发环境
hertz是http框架，适用于和用户交互的，所以先生成一个hello world

vscode插件

| 插件名               | 功能说明 |
|----------------------|----------|
| **Go**               | 提供 Go 语言的语法高亮、自动补全、调试支持。|
| **Golang Tools**     | 依赖 Go 插件，包含 lint、format、test 等开发辅助工具集。|
| **Docker**           | 支持容器开发，提供 Dockerfile、docker-compose 支持与容器管理功能。|
| **MySQL**            | 用于连接和管理 MySQL 数据库（例如 SQLTools 或类似插件）。|
| **Material Icon Theme** | 更美观的文件图标主题，提升文件识别效率。|
| **YAML**             | 支持 YAML 文件语法高亮、校验、自动补全等功能，常用于配置文件。|
| **vscode-proto3**    | 提供 `.proto` 文件（Protocol Buffers v3）语法支持，包括高亮和代码提示。|
| **Makefile Tools**   | 支持 `Makefile` 的语法高亮、构建任务与调试操作。|


终端插件
oh my zsh

plugins
- zsh-syntax-highlighting
- zsh-autosuggestions



`go mod init` 是 Go 语言模块系统的初始化命令，用于创建一个新的 Go 模块。这个命令会在当前目录下生成一个 `go.mod` 文件，该文件定义了模块的导入路径和依赖关系。

使用自己的仓库名称作为模块路径:

```bash
cd hello_world 注意必须在同一目录下才能go mod init
go mod init github.com/charleschile/TikTok-E-Commerce-Gomall/tutorial/hello_world
```

模块路径决定了其他项目如何导入你的代码。选择合适的模块路径取决于:
- 代码托管位置

使用 `go mod init` 后，你可以通过 `go get` 添加依赖，系统会自动更新 `go.mod` 文件。

```bash
go get -u github.com/cloudwego/hertz
```



go get -u github.com/cloudwego/hertz 这个命令的作用是下载并安装 Hertz 框架的最新版本。

- go get 是 Go 语言用于下载和安装依赖包的命令

- -u 参数表示更新（update）已存在的包到最新版本

- github.com/cloudwego/hertz 是要下载的包的路径，即字节跳动开源的高性能 Go HTTP 框架

执行这个命令后：

1. Go 会下载 Hertz 框架及其依赖

1. 更新你项目的 go.mod 和 go.sum 文件，添加 Hertz 作为依赖





https://www.cloudwego.io/zh/docs/hertz/getting-started/

```go
package main

import (
    "context"

    "github.com/cloudwego/hertz/pkg/app"
    "github.com/cloudwego/hertz/pkg/app/server"
    "github.com/cloudwego/hertz/pkg/common/utils"
    "github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
    h := server.Default()

    h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
            c.JSON(consts.StatusOK, utils.H{"message": "pong"})
    })

    h.Spin()
}

```




go mod tidy 命令的作用是整理和清理项目的依赖关系。具体功能包括：

添加缺失的依赖：检查代码中所有 import 语句，将使用但未在 go.mod 中声明的依赖添加进来

移除未使用的依赖：清理 go.mod 中已声明但代码中没有实际使用的依赖

更新 go.sum 文件：重新计算并确保依赖的校验和正确，保证依赖的完整性和安全性

解决依赖冲突：尝试解决项目中可能存在的依赖版本冲突问题





- context 是 Go 标准库中的一个包，用于在 API 边界和进程之间传递截止日期、取消信号和其他请求范围的值

- 在 Web 服务中，context 通常用于：

请求超时控制

取消长时间运行的操作

在请求处理流程中传递请求相关的数据



![helloworld](img/helloworld.png)





