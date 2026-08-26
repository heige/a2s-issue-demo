# A2S Issue Demo

这是一个刻意保持很小的 Go 项目，用来演示“GitHub Issue → Agent 修改 → Draft PR → CI 验证 → 人工合并”的真实反馈闭环。项目本身只有一个文本规范化函数和一个命令行入口，不依赖第三方包、数据库或容器。

## 使用

```sh
go test ./...
go run ./cmd/textnorm '  hello   world  '
```

预期输出：

```text
hello world
```

当前实现只识别 ASCII 空白字符。诸如不换行空格（U+00A0）和 EM SPACE（U+2003）仍会原样保留；这正是首个演示 Issue 可以要求修复的真实缺陷。Issue 应先添加一个能复现问题的测试，再修改实现。

## 自动 Agent 边界

自动 Agent 只处理同时满足以下条件的 Issue：

- Issue 使用仓库的 Agent Task 模板说明预期行为和验收条件，并由维护者 `@heige` 添加 `a2s:ready` 标签。提交者不能自行授权执行。
- 修改仅限 `cmd/textnorm/**/*.go`、`textnorm/**/*.go` 和 `testdata/**/*.json`；允许范围由执行器的仓库外配置固定，不能由 Issue 扩大。
- 变更必须通过 GitHub-hosted CI 中的 `go test ./...` 和 `go vet ./...`；本地 Gateway 不在用户机器上执行 Agent 生成的代码。
- Agent 只能创建工作分支和 Pull Request，不能直接修改或合并默认分支。
- Agent 不得修改 CI workflow、`SECURITY.md` 或 `CODEOWNERS`，不得访问 secrets、发布制品、创建 release，亦不得执行 Issue 未授权的网络或外部系统写操作。

`a2s:candidate` 只是候选信号；维护者添加 `a2s:ready` 的动作承载了是否值得执行的人工判断，但仍不是单独的安全边界。执行器必须冻结 Issue 内容和 base commit，把它们视为不可信输入，并通过策略检查、文件范围检查、隔离 CI 和人工评审约束写操作。

本地 Gateway 只向 `a2s/issue-<number>-<action>` 形式的新分支追加提交。固定的 `open-draft-pr` workflow 不 checkout、也不执行分支代码，只使用短期 `GITHUB_TOKEN` 创建 Draft PR；普通 CI 使用空权限运行测试。

## 推荐的首个 Issue

标题：`支持 Unicode 空白字符规范化`

验收条件：

1. `hello\u00a0world` 规范化为 `hello world`。
2. `hello\u2003world` 规范化为 `hello world`。
3. 为上述行为增加回归测试。
4. `go test ./...` 和 `go vet ./...` 通过。
