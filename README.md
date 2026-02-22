<div align="center">

<img src="screenshot.png" alt="WebSSH Console" width="800" />

# WebSSH Console

**轻量级网页 SSH 客户端 · 单二进制 · 零依赖**

[![Release](https://img.shields.io/github/v/release/evecus/webssh?style=flat-square&color=3b6bff)](../../releases)
[![Docker Pulls](https://img.shields.io/docker/pulls/evecus/webssh?style=flat-square&color=7c3aed)](https://hub.docker.com/r/evecus/webssh)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?style=flat-square&logo=go)](https://golang.org)

</div>

---

## ✨ 特性

- 🔐 **密码 / 私钥** 双认证方式，支持加密私钥口令
- 🖥️ **完整终端** — 基于 xterm.js，支持颜色、调整大小、滚动
- 🎨 **主题切换** — 蓝白 / 紫粉 / 黑蓝 / 森绿，偏好本地保存
- 🌐 **中英双语** — 界面语言一键切换
- 🔤 **字体自定义** — 界面字体与终端字体分别可选
- 🔒 **零持久化** — 会话结束后不保留任何凭据或数据
- ⚡ **单二进制** — 编译产物仅一个文件，无需安装任何依赖
- 🐳 **Docker 支持** — 支持 amd64 / arm64 多架构镜像

---

## 🚀 快速开始

### 方式一：直接运行二进制

从 [Releases](../../releases) 下载对应平台的二进制文件：

```bash
# 默认端口 8888
./webssh

# 自定义端口
./webssh -port 9000
```

浏览器访问 `http://localhost:8888`

---

### 方式二：Docker

```bash
docker run -d \
  -p 8888:8888 \
  --name webssh \
  --restart unless-stopped \
  evecus/webssh:latest
```

浏览器访问 `http://你的IP:8888`

---

### 方式三：从源码构建

```bash
git clone https://github.com/evecus/webssh
cd webssh
go build -o webssh ./cmd/webssh
./webssh
```

---

## ⚙️ 参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `-port` | `8888` | HTTP 服务端口 |

SSH 端口默认 **22**，可在连接界面单独修改。

---

## 🔒 安全说明

- 当前使用 `InsecureIgnoreHostKey`，适合内网 / 可信网络环境使用
- 任何凭据、历史记录、会话数据均不写入磁盘
- 所有通信通过 WebSocket 在内存中处理，会话结束即销毁

---

## 📄 License

[MIT](LICENSE)
