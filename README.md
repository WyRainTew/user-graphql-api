# GraphQL 用户信息 API

一个简单的GraphQL API，用于查询用户信息，基于Go和gqlgen构建。

## 项目介绍

这个项目是使用Go语言和gqlgen库实现的GraphQL API服务，支持通过userId查询用户的基本信息。项目遵循现代GraphQL开发实践，采用了模式优先的设计方法。

## 功能特点

- 使用GraphQL查询用户信息
- 内存中存储示例用户数据
- 支持通过userId查询单个用户
- 返回用户姓名、年龄和性别信息
- 提供GraphQL Playground界面进行交互式查询测试

## 技术栈

- **Go**: 主要编程语言
- **gqlgen**: GraphQL的Go实现库
- **GraphQL**: API查询语言

## 项目结构

```
user-graphql-api/
├── go.mod              // Go模块定义
├── go.sum              // 依赖版本锁定
├── gqlgen.yml          // gqlgen配置文件
├── graph               // GraphQL相关代码
│   ├── generated.go    // 自动生成的GraphQL代码
│   ├── model           // 数据模型
│   │   └── models.go   // 用户数据模型定义
│   ├── resolver.go     // 解析器和数据源
│   ├── schema.graphqls // GraphQL模式定义
│   └── schema.resolvers.go // 解析器实现
├── server.go           // 主服务器文件
└── tools.go            // 工具依赖
```

## 快速开始

1. 克隆项目
```bash
git clone https://github.com/WyRainTew/user-graphql-api.git
cd user-graphql-api
```

2. 安装依赖
```bash
go mod tidy
```

3. 运行服务器
```bash
go run server.go
```

4. 打开GraphQL Playground
在浏览器中访问 http://localhost:8080/

## 查询示例

```graphql
query {
  userInfo(userId: "aaa") {
    name
    age
    sex
  }
}
```

上述查询将返回ID为"aaa"的用户信息，例如：

```json
{
  "data": {
    "userInfo": {
      "name": "张三",
      "age": 28,
      "sex": "男"
    }
  }
}
```

## 可用用户ID

项目中预设了三个用户数据：
- aaa：张三，28岁，男


##查询例子
![image](https://github.com/user-attachments/assets/4f507bcc-7b12-466b-871d-cdc48395dbe2)


## 参考资料

本项目基于[gqlgen](https://github.com/99designs/gqlgen)开发，参考了其中的示例，特别是[todo示例](https://github.com/99designs/gqlgen/tree/master/_examples/todo)。 
