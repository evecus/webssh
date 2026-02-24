# WebSSH Console

一个轻量、美观的 Web SSH 终端，支持认证、持久化存储和移动端优化。

## 功能特性

- 🎨 **紫粉默认主题**（可切换蓝白/黑蓝/森绿）
- 🔐 **可选登录认证**（`--auth true`）
- 💾 **可选数据持久化**（`--store true`）
- 📱 **移动端全面优化**（虚拟键盘、全屏终端、触摸复制）
- 🔑 支持密码和私钥两种 SSH 认证方式

---

## 运行参数

| 参数 | 默认值 | 说明 |
|------|--------|------|
| `--port <端口>` | `8888` | HTTP 监听端口 |
| `--auth true` | `false` | 启用网页登录认证 |
| `--store true` | `false` | 启用数据持久化存储 |

### 示例

```bash
# 最简运行（无认证、无存储）
./webssh

# 自定义端口
./webssh --port 9000

# 启用认证（首次访问自动跳转设置页）
./webssh --auth true

# 启用存储（保存主题/字体/SSH配置）
./webssh --store true

# 全功能（认证 + 存储 + 自定义端口）
./webssh --auth true --store true --port 8080
```

---

## 四种运行模式说明

### 情况一：仅认证（`--auth true`）
- 首次访问引导设置账户和密码
- **只**持久化账户密码，其他设置不保存

### 情况二：仅存储（`--store true`）
- **无需**账户密码即可访问
- 保存：主题设置、字体设置、SSH 连接信息
- 主界面新增：**主机名**输入框、**保存**按钮、**SSH列表**按钮

### 情况三：无参数
- 无认证、无存储，直接使用

### 情况四：认证 + 存储（`--auth true --store true`）
- 需要账户密码登录
- 保存所有数据：账户、主题、字体、SSH 连接

---

## Docker 部署

### 环境变量

| 变量 | 默认值 | 说明 |
|------|--------|------|
| `PORT` | `8888` | 监听端口 |
| `AUTH` | `false` | 启用认证 |
| `STORE` | `false` | 启用存储 |

### Docker 命令

```bash
# 基础运行
docker run -d -p 8888:8888 webssh

# 启用认证 + 存储
docker run -d -p 8888:8888 \
  -e AUTH=true \
  -e STORE=true \
  -v $(pwd)/data:/app/data \
  webssh

# 使用 docker-compose
docker-compose up -d
```

### 构建镜像

```bash
docker build -t webssh .
```

---

## data 目录结构

启用 `--store` 或 `--auth` 后，会在二进制文件同目录自动创建 `data/` 目录：

```
data/
├── auth.json       # 账户密码（仅 --auth 时创建）
├── settings.json   # 主题/字体设置（仅 --store 时）
└── ssh.json        # SSH 连接配置（仅 --store 时）
```

---

## 编译

```bash
git clone ...
cd webssh-main
go build -o webssh ./cmd/webssh
./webssh --help
```
