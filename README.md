# WebSSH

A lightweight web-based SSH client written in Go. No information is stored after each session.

![WebSSH Console](screenshot.png)

## Features

- 🔐 Password or Private Key authentication
- 🔑 Encrypted private keys with passphrase support
- 🖥️ Full xterm.js terminal emulator
- 🔒 Zero persistence — no session data stored
- ⚡ Single binary, zero dependencies
- 🌐 WebSocket-based real-time communication

## Usage

### Download

Download the latest binary from [Releases](../../releases).

### Run

```bash
# Default port 8888
./webssh

# Custom port
./webssh -port 9000
```

Then open `http://localhost:8888` in your browser.

### Build from Source

```bash
git clone https://github.com/yourusername/webssh
cd webssh
go build -o webssh ./cmd/webssh
./webssh
```

## Configuration

| Flag | Default | Description |
|------|---------|-------------|
| `-port` | `8888` | HTTP server port |

SSH port defaults to **22** but can be changed in the UI per connection.

## Security

- Host key verification uses `InsecureIgnoreHostKey` for simplicity — suitable for internal/trusted networks
- No credentials, history, or session data are ever written to disk
- All communication happens over WebSocket in-memory

## License

MIT
