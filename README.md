# K8s Dashboard Backend

一个用 Go + Gin + client-go 编写的 Kubernetes 集群管理 REST API 服务。提供标准 HTTP 接口，可供前端页面、Postman 或其他客户端调用，实时获取集群中 Pod、Node、Namespace 的信息。

## 功能

- `GET /pods` —— 返回所有命名空间下的 Pod 列表（名称、命名空间、状态）
- `GET /nodes` —— 返回集群中所有节点列表
- `GET /namespaces` —— 返回集群中所有命名空间列表

所有接口均以 JSON 格式返回数据。

## 技术栈

- Go
- [Gin](https://github.com/gin-gonic/gin)：Web 框架，负责路由和 HTTP 请求处理
- [client-go](https://github.com/kubernetes/client-go)：Kubernetes 官方 Go SDK

## 运行方式

确保本地已配置好可用的 kubeconfig（默认读取 `~/.kube/config`），然后：

```bash
go mod tidy
go run main.go
```

服务启动后监听 `8080` 端口。

## 接口示例

```bash
curl http://localhost:8080/pods
curl http://localhost:8080/nodes
curl http://localhost:8080/namespaces
```

返回示例（`/namespaces`）：

```json
{
  "count": 4,
  "namespaces": [
    {"name": "default"},
    {"name": "kube-node-lease"},
    {"name": "kube-public"},
    {"name": "kube-system"}
  ]
}
```

## 开发与测试环境

本项目在 WSL2 + K3s 单节点集群上开发和验证。
