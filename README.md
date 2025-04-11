# API-Import

API-Import 是一个用于将接口文档同步到 API 文档平台的命令行工具。

## 功能特性

- 支持将 Swagger 文档同步到 ApiFox
- 支持将 Swagger 文档同步到 YApi
- 简单易用的命令行界面
- 灵活的配置选项

## 安装

```bash
go install github.com/fzf-labs/api-import@latest
```

## 使用方法

### 同步到 ApiFox

```bash
api-import swagger apifox -p <项目ID> -t <令牌> -i <输入路径>
```

参数说明：
- `-p, --projectId`：ApiFox 项目 ID
- `-t, --token`：ApiFox 访问令牌
- `-i, --inPutPath`：Swagger 文档输入路径（默认为 "./api"）

### 同步到 YApi

```bash
api-import swagger yapi -u <YApi URL> -t <令牌> -i <输入路径>
```

参数说明：
- `-u, --url`：YApi 服务器 URL
- `-t, --token`：YApi 访问令牌
- `-i, --inPutPath`：Swagger 文档输入路径（默认为 "./api"）

## 示例

同步到 ApiFox：

```bash
api-import swagger apifox -p 123456 -t your_token_here -i ./docs/swagger
```

同步到 YApi：

```bash
api-import swagger yapi -u https://your-yapi-server.com -t your_token_here -i ./docs/swagger
```

## 项目结构

```
api-import/
├── cmd.go                  # 主命令入口
├── swagger/                # Swagger 相关功能
│   ├── cmd.go              # Swagger 命令
│   ├── apifox/             # ApiFox 集成
│   │   ├── apifox.go       # ApiFox 实现
│   │   └── cmd.go          # ApiFox 命令
│   └── yapi/               # YApi 集成
│       ├── yapi.go         # YApi 实现
│       └── cmd.go          # YApi 命令
└── utils/                  # 工具函数
```

## 依赖

- github.com/go-resty/resty/v2：HTTP 客户端
- github.com/spf13/cobra：命令行工具
- github.com/tidwall/gjson：JSON 处理

## 贡献

欢迎提交 Issue 或 Pull Request。

## 许可证

本项目使用 [MIT 许可证](LICENSE)。 