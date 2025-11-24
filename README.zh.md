[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/wirekratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/wirekratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/wirekratos)](https://pkg.go.dev/github.com/orzkratos/wirekratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/wirekratos/main.svg)](https://coveralls.io/github/orzkratos/wirekratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/wirekratos.svg)](https://github.com/orzkratos/wirekratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/wirekratos)](https://goreportcard.com/report/github.com/orzkratos/wirekratos)

# wirekratos

Kratos 集成的 Wire workspace 模式修复包。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[English](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 问题描述

使用 Google Wire 配合 Go workspace 模式 (go.work) 时，`go generate ./...` 会失败：

```
go generate ./...
go: -mod may only be set to readonly or vendor when in workspace mode, but it is set to "mod"
        Remove the -mod flag to use the default readonly value,
        or set GOWORK=off to disable workspace mode.
xxx/xxx/wire_gen.go:3: running "go": exit status 1
```

Wire 生成的代码包含与 workspace 模式冲突的指令：
```go
//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject
```

## 解决方案

自动移除 Wire 生成文件中的 `-mod=mod` 标志，使其兼容 workspace 模式：
```go
//go:generate go run github.com/google/wire/cmd/wire
```

## Wire 状态 (2025)

- 最新版本：v0.7.0（2025年8月22日）
- 仓库状态：已归档，只读访问（2025年8月25日）
- Workspace 问题：未修复（PR #410 仍处于 open 状态）
- 由于 Wire 不会再有更新，本包保持必要

## 安装

```bash
go install github.com/orzkratos/wirekratos/cmd/wirekratos@latest
```

## 使用方法

### 自动模式（Kratos 项目）
```bash
wirekratos -framework=kratos
```
自动检测 Kratos 项目中的 `cmd/项目名/wire_gen.go`。

### 相对路径模式
```bash
wirekratos -name=cmd/myproject/wire_gen.go
```

### 绝对路径模式
```bash
wirekratos -path=/absolute/path/to/wire_gen.go
```

### 调试模式
```bash
wirekratos -framework=kratos -debug
```

## 演示项目

展示在生产环境 Kratos 应用中集成的完整演示项目：

**[wirekratos-demos](https://github.com/orzkratos/wirekratos-demos)** - 功能完整的 Kratos 应用，集成 Wire 依赖注入

* [demo1kratos](https://github.com/orzkratos/wirekratos-demos/tree/main/demo1kratos) - 基础 Kratos 项目设置
* [demo2kratos](https://github.com/orzkratos/wirekratos-demos/tree/main/demo2kratos) - 高级 Kratos 项目，使用 workspace

## 工作原理

1. 定位 `wire_gen.go` 文件
2. 检查 Wire 生成代码的签名
3. 查找 `//go:generate` 指令（第3行）
4. 如果存在则移除 `-mod=mod` 标志
5. 保持其他代码不变

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-22 10:00:00.000000 +0000 UTC -->

## 📄 许可证

MIT License - 查看 [LICENSE](LICENSE)。

---

## 💬 联系和反馈

欢迎贡献！提交错误报告、建议功能和贡献代码：

- 🐛 **发现错误？** 在 GitHub 上提交问题并附带复现步骤
- 💡 **有新想法？** 创建问题进行讨论
- 📖 **文档不清楚？** 报告以便我们改进
- 🚀 **需要新功能？** 分享使用场景帮助我们理解需求
- ⚡ **性能问题？** 通过报告慢速操作帮助我们优化
- 🔧 **配置问题？** 询问有关复杂设置的问题
- 📢 **关注项目进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享此包如何改进工作流程
- 💬 **反馈？** 我们欢迎建议和意见

---

## 🔧 开发

新代码贡献，遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）。
2. **Clone**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）。
3. **导航**：导航到克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）。
5. **编码**：实现更改并附带全面测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：更新文档以支持面向客户的更改并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容代码
10. **推送**：推送到分支（`git push origin feature/xxx`）。
11. **PR**：在 GitHub 上打开合并请求（在 GitHub 网页上）并附带详细描述。

请确保测试通过并包含相关文档更新。

---

## 🌟 支持

欢迎通过提交合并请求和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给 GitHub 星标** 如果这个项目对你有帮助
- 🤝 **与队友分享** 和（golang）编程朋友
- 📝 **撰写技术博客** 关于开发工具和工作流程 - 我们提供内容写作支持
- 🌟 **加入生态系统** - 致力于支持开源和（golang）开发场景

**祝编码愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/wirekratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/wirekratos)
