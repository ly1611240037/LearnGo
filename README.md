# LearnGo

这是我的 Go 语言学习项目，记录 Go 基础、标准库和 Web 开发练习。

## 项目结构

```text
LearnGo/
├─ 001/                 # Go 基础与标准库练习
└─ 002/                 # HTTP、Gin 等 Web 开发练习
   └─ gin_learn_01/    # 第一个 Gin 应用
```

每个练习目录通常是一个独立的 Go Module，需要进入对应目录后执行命令。

## 第一个 Gin 应用

目录：`002/gin_learn_01`

运行：

```bash
cd 002/gin_learn_01
go run main.go
```

启动后访问 `http://localhost:8080`，浏览器会返回 `Hello, Gin!`。

## 常用 Go 命令

```bash
go version       # 查看 Go 版本
go mod tidy      # 下载并整理依赖
go run main.go   # 运行当前练习
gofmt -w main.go # 格式化代码
go test ./...    # 运行测试
```

## Git 提交方式

在项目根目录 `D:\LearnGo` 下执行 Git 命令。

### 查看修改

```bash
git status
```

### 提交代码

```bash
git add .
git commit -m "说明本次修改内容"
```

例如：

```bash
git add 002/gin_learn_01
git commit -m "add first Gin application"
```

提交信息建议简短说明本次做了什么，例如：

- `add HTTP server example`
- `learn Gin routing`
- `fix client request error`

### 同步到 GitHub 和 Gitee

当前远程仓库：

```text
github -> https://github.com/ly1611240037/LearnGo.git
origin -> https://gitee.com/ly1611240037/learn-go.git
```

提交后，分别推送到两个远程仓库：

```bash
git push github master
git push origin master
```

## 一次完整的提交流程

```bash
cd D:\LearnGo
git status
git add .
git commit -m "describe your changes"
git push github master
git push origin master
```

## 注意事项

- 修改代码后，先运行程序或测试，确认没有明显问题再提交。
- `git add .` 会添加所有修改；不确定时先执行 `git status`。
- 一次提交尽量只完成一类事情，方便以后查看和回退。
- 建议以本地代码为主，再同时推送到 GitHub 和 Gitee。
