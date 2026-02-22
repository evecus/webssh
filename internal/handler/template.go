package handler

const indexHTML = `<!DOCTYPE html>
<html lang="zh">
<head>
  <meta charset="UTF-8"/>
  <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  <title>WebSSH Console</title>
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link href="https://fonts.googleapis.com/css2?family=JetBrains+Mono:wght@400;600&family=Syne:wght@400;600;800&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/xterm@5.3.0/css/xterm.css"/>
  <style>
    *, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
    :root {
      --bg: #0a0e1a;
      --surface: #111827;
      --surface2: #1a2236;
      --border: #1e2d45;
      --accent: #00d4ff;
      --accent2: #7c3aed;
      --text: #e2e8f0;
      --text-muted: #64748b;
      --success: #10b981;
      --error: #ef4444;
      --font-mono: 'JetBrains Mono', monospace;
      --font-sans: 'Syne', sans-serif;
    }
    html, body { height: 100%; }
    body {
      background: var(--bg);
      color: var(--text);
      font-family: var(--font-sans);
      display: flex;
      flex-direction: column;
      min-height: 100vh;
    }
    body::before {
      content: '';
      position: fixed;
      inset: 0;
      background:
        radial-gradient(ellipse 80% 50% at 20% 20%, rgba(0,212,255,0.06) 0%, transparent 60%),
        radial-gradient(ellipse 60% 60% at 80% 80%, rgba(124,58,237,0.07) 0%, transparent 60%);
      pointer-events: none;
      z-index: 0;
    }
    body::after {
      content: '';
      position: fixed;
      inset: 0;
      background-image:
        linear-gradient(rgba(0,212,255,0.03) 1px, transparent 1px),
        linear-gradient(90deg, rgba(0,212,255,0.03) 1px, transparent 1px);
      background-size: 40px 40px;
      pointer-events: none;
      z-index: 0;
    }
    .container {
      position: relative;
      z-index: 1;
      width: 100%;
      max-width: 960px;
      margin: 0 auto;
      padding: 40px 24px;
      flex: 1;
    }
    .header {
      text-align: center;
      margin-bottom: 48px;
      animation: fadeDown 0.6s ease both;
    }
    .header h1 {
      font-size: clamp(2rem, 5vw, 3.2rem);
      font-weight: 800;
      letter-spacing: -0.02em;
      background: linear-gradient(135deg, #00d4ff 0%, #7c3aed 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
      margin-bottom: 8px;
    }
    .header .subtitle {
      font-family: var(--font-mono);
      font-size: 0.75rem;
      color: var(--text-muted);
      letter-spacing: 0.2em;
      text-transform: uppercase;
    }
    .header .accent-line {
      width: 60px;
      height: 2px;
      background: linear-gradient(90deg, var(--accent), var(--accent2));
      margin: 16px auto 0;
      border-radius: 2px;
    }
    .card {
      background: var(--surface);
      border: 1px solid var(--border);
      border-radius: 16px;
      padding: 40px;
      position: relative;
      overflow: hidden;
      animation: fadeUp 0.6s ease 0.1s both;
    }
    .card::before {
      content: '';
      position: absolute;
      top: 0; left: 0; right: 0;
      height: 1px;
      background: linear-gradient(90deg, transparent, var(--accent), transparent);
      opacity: 0.5;
    }
    .form-grid {
      display: grid;
      grid-template-columns: 1fr 1fr;
      gap: 24px;
    }
    @media (max-width: 600px) {
      .form-grid { grid-template-columns: 1fr; }
      .card { padding: 24px; }
    }
    .field { display: flex; flex-direction: column; gap: 8px; }
    .field.full { grid-column: 1 / -1; }
    label {
      font-family: var(--font-mono);
      font-size: 0.7rem;
      letter-spacing: 0.15em;
      text-transform: uppercase;
      color: var(--text-muted);
    }
    label span.required { color: var(--accent); margin-left: 2px; }
    input[type="text"],
    input[type="password"],
    input[type="number"] {
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 12px 16px;
      font-family: var(--font-mono);
      font-size: 0.9rem;
      color: var(--text);
      transition: border-color 0.2s, box-shadow 0.2s;
      outline: none;
      width: 100%;
    }
    input:focus {
      border-color: var(--accent);
      box-shadow: 0 0 0 3px rgba(0,212,255,0.12);
    }
    input::placeholder { color: var(--text-muted); }
    .file-input-wrapper {
      display: flex;
      align-items: center;
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 8px;
      overflow: hidden;
      transition: border-color 0.2s;
    }
    .file-input-wrapper:focus-within {
      border-color: var(--accent);
      box-shadow: 0 0 0 3px rgba(0,212,255,0.12);
    }
    .file-btn {
      background: var(--surface);
      border: none;
      border-right: 1px solid var(--border);
      padding: 12px 18px;
      font-family: var(--font-mono);
      font-size: 0.78rem;
      color: var(--accent);
      cursor: pointer;
      white-space: nowrap;
      transition: background 0.2s;
      display: flex;
      align-items: center;
      gap: 6px;
    }
    .file-btn:hover { background: rgba(0,212,255,0.08); }
    .file-name {
      flex: 1;
      padding: 12px 16px;
      font-family: var(--font-mono);
      font-size: 0.8rem;
      color: var(--text-muted);
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
    }
    #private-key-file { display: none; }
    .auth-tabs {
      display: flex;
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 8px;
      padding: 4px;
      width: fit-content;
    }
    .auth-tab {
      padding: 8px 20px;
      border: none;
      background: transparent;
      color: var(--text-muted);
      font-family: var(--font-mono);
      font-size: 0.75rem;
      letter-spacing: 0.1em;
      text-transform: uppercase;
      cursor: pointer;
      border-radius: 6px;
      transition: all 0.2s;
    }
    .auth-tab.active {
      background: var(--surface);
      color: var(--accent);
      box-shadow: 0 1px 4px rgba(0,0,0,0.3);
    }
    .auth-pane { display: none; grid-column: 1 / -1; }
    .auth-pane.active { display: contents; }
    .btn-connect {
      grid-column: 1 / -1;
      margin-top: 8px;
      padding: 14px 32px;
      background: linear-gradient(135deg, var(--accent) 0%, var(--accent2) 100%);
      border: none;
      border-radius: 8px;
      color: #fff;
      font-family: var(--font-sans);
      font-size: 1rem;
      font-weight: 600;
      letter-spacing: 0.05em;
      cursor: pointer;
      transition: opacity 0.2s, transform 0.15s;
    }
    .btn-connect:hover { opacity: 0.9; transform: translateY(-1px); }
    .btn-connect:active { transform: translateY(0); }
    .btn-connect:disabled { opacity: 0.5; cursor: not-allowed; transform: none; }
    .status-bar {
      display: flex;
      align-items: center;
      gap: 10px;
      margin-top: 20px;
      padding: 10px 16px;
      background: var(--surface2);
      border: 1px solid var(--border);
      border-radius: 8px;
      font-family: var(--font-mono);
      font-size: 0.78rem;
      min-height: 44px;
      animation: fadeUp 0.4s ease both;
    }
    .status-dot {
      width: 8px; height: 8px;
      border-radius: 50%;
      background: var(--text-muted);
      flex-shrink: 0;
      transition: background 0.3s;
    }
    .status-dot.connecting { background: #f59e0b; animation: pulse 1s infinite; }
    .status-dot.connected { background: var(--success); }
    .status-dot.error { background: var(--error); }
    .status-text { color: var(--text-muted); flex: 1; }
    .btn-disconnect {
      padding: 4px 12px;
      background: transparent;
      border: 1px solid var(--error);
      border-radius: 6px;
      color: var(--error);
      font-family: var(--font-mono);
      font-size: 0.7rem;
      cursor: pointer;
      transition: all 0.2s;
      display: none;
    }
    .btn-disconnect:hover { background: var(--error); color: #fff; }
    .btn-disconnect.visible { display: block; }
    #terminal-card {
      margin-top: 24px;
      background: #0d1117;
      border: 1px solid var(--border);
      border-radius: 16px;
      overflow: hidden;
      display: none;
      animation: fadeUp 0.5s ease both;
    }
    #terminal-card.visible { display: block; }
    .terminal-header {
      display: flex;
      align-items: center;
      padding: 12px 16px;
      background: var(--surface);
      border-bottom: 1px solid var(--border);
      gap: 8px;
    }
    .term-dot { width: 12px; height: 12px; border-radius: 50%; }
    .term-dot-red { background: #ff5f57; }
    .term-dot-yellow { background: #ffbd2e; }
    .term-dot-green { background: #28c840; }
    .terminal-title {
      font-family: var(--font-mono);
      font-size: 0.75rem;
      color: var(--text-muted);
      margin-left: 8px;
      flex: 1;
    }
    #terminal { padding: 8px; }
    @keyframes fadeDown {
      from { opacity: 0; transform: translateY(-20px); }
      to { opacity: 1; transform: translateY(0); }
    }
    @keyframes fadeUp {
      from { opacity: 0; transform: translateY(20px); }
      to { opacity: 1; transform: translateY(0); }
    }
    @keyframes pulse {
      0%, 100% { opacity: 1; }
      50% { opacity: 0.4; }
    }
  </style>
</head>
<body>
<div class="container">
  <div class="header">
    <h1>WebSSH Console</h1>
    <p class="subtitle">Secure Shell in Your Browser</p>
    <div class="accent-line"></div>
  </div>

  <div class="card">
    <div class="form-grid">
      <div class="field">
        <label>主机地址 (Hostname)<span class="required">*</span></label>
        <input type="text" id="host" placeholder="192.168.1.1 or example.com"/>
      </div>
      <div class="field">
        <label>端口 (Port)</label>
        <input type="number" id="port" value="22" min="1" max="65535"/>
      </div>
      <div class="field">
        <label>用户名 (Username)<span class="required">*</span></label>
        <input type="text" id="username" placeholder="root"/>
      </div>
      <div class="field"></div>

      <div class="field full">
        <div class="auth-tabs">
          <button class="auth-tab active" data-tab="password">密码登录</button>
          <button class="auth-tab" data-tab="key">私钥登录</button>
        </div>
      </div>

      <div class="auth-pane active" id="pane-password">
        <div class="field full">
          <label>密码 (Password)</label>
          <input type="password" id="password" placeholder="请输入密码"/>
        </div>
      </div>

      <div class="auth-pane" id="pane-key">
        <div class="field">
          <label>私钥 (Private Key)</label>
          <div class="file-input-wrapper">
            <button class="file-btn" onclick="document.getElementById('private-key-file').click()">
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/></svg>
              选择文件
            </button>
            <span class="file-name" id="key-file-name">未选择私钥文件</span>
          </div>
          <input type="file" id="private-key-file"/>
        </div>
        <div class="field">
          <label>密钥口令 (Passphrase)</label>
          <input type="password" id="passphrase" placeholder="如果需要请输入密钥口令"/>
        </div>
      </div>

      <button class="btn-connect" id="btn-connect" onclick="connect()">连接 →</button>
    </div>
  </div>

  <div class="status-bar" id="status-bar" style="display:none;">
    <div class="status-dot" id="status-dot"></div>
    <span class="status-text" id="status-text">准备就绪</span>
    <button class="btn-disconnect" id="btn-disconnect" onclick="disconnect()">断开连接</button>
  </div>

  <div id="terminal-card">
    <div class="terminal-header">
      <div class="term-dot term-dot-red"></div>
      <div class="term-dot term-dot-yellow"></div>
      <div class="term-dot term-dot-green"></div>
      <span class="terminal-title" id="terminal-title">terminal</span>
    </div>
    <div id="terminal"></div>
  </div>
</div>

<script src="https://cdn.jsdelivr.net/npm/xterm@5.3.0/lib/xterm.js"></script>
<script src="https://cdn.jsdelivr.net/npm/xterm-addon-fit@0.8.0/lib/xterm-addon-fit.js"></script>
<script>
let ws = null, term = null, fitAddon = null, currentTab = 'password', privateKeyData = '';

document.querySelectorAll('.auth-tab').forEach(btn => {
  btn.addEventListener('click', () => {
    currentTab = btn.dataset.tab;
    document.querySelectorAll('.auth-tab').forEach(b => b.classList.remove('active'));
    document.querySelectorAll('.auth-pane').forEach(p => p.classList.remove('active'));
    btn.classList.add('active');
    document.getElementById('pane-' + currentTab).classList.add('active');
  });
});

document.getElementById('private-key-file').addEventListener('change', (e) => {
  const file = e.target.files[0];
  if (!file) return;
  document.getElementById('key-file-name').textContent = file.name;
  const reader = new FileReader();
  reader.onload = (ev) => { privateKeyData = ev.target.result; };
  reader.readAsText(file);
});

function setStatus(state, msg) {
  const bar = document.getElementById('status-bar');
  bar.style.display = 'flex';
  document.getElementById('status-dot').className = 'status-dot ' + state;
  document.getElementById('status-text').textContent = msg;
  const discBtn = document.getElementById('btn-disconnect');
  state === 'connected' ? discBtn.classList.add('visible') : discBtn.classList.remove('visible');
}

function initTerm() {
  if (term) term.dispose();
  term = new Terminal({
    theme: {
      background: '#0d1117', foreground: '#e2e8f0', cursor: '#00d4ff',
      selectionBackground: 'rgba(0,212,255,0.2)',
      black: '#1a2236', red: '#ef4444', green: '#10b981', yellow: '#f59e0b',
      blue: '#3b82f6', magenta: '#7c3aed', cyan: '#00d4ff', white: '#e2e8f0',
      brightBlack: '#334155', brightRed: '#f87171', brightGreen: '#34d399',
      brightYellow: '#fbbf24', brightBlue: '#60a5fa', brightMagenta: '#a78bfa',
      brightCyan: '#67e8f9', brightWhite: '#f8fafc',
    },
    fontFamily: "'JetBrains Mono', monospace",
    fontSize: 14, lineHeight: 1.5,
    cursorBlink: true, cursorStyle: 'bar',
    scrollback: 5000,
  });
  fitAddon = new FitAddon.FitAddon();
  term.loadAddon(fitAddon);
  term.open(document.getElementById('terminal'));
  setTimeout(() => fitAddon.fit(), 50);

  term.onData((data) => {
    if (ws && ws.readyState === WebSocket.OPEN)
      ws.send(JSON.stringify({ type: 'input', data }));
  });

  window.addEventListener('resize', () => {
    if (fitAddon) fitAddon.fit();
    if (term && ws && ws.readyState === WebSocket.OPEN)
      ws.send(JSON.stringify({ type: 'resize', rows: term.rows, cols: term.cols }));
  });
}

function connect() {
  const host = document.getElementById('host').value.trim();
  const port = parseInt(document.getElementById('port').value) || 22;
  const username = document.getElementById('username').value.trim();
  if (!host || !username) { setStatus('error', '请填写主机地址和用户名'); return; }

  const password = currentTab === 'password' ? document.getElementById('password').value : '';
  const private_key = currentTab === 'key' ? privateKeyData : '';
  const passphrase = currentTab === 'key' ? document.getElementById('passphrase').value : '';
  if (!password && !private_key) { setStatus('error', '请提供密码或私钥'); return; }

  document.getElementById('btn-connect').disabled = true;
  setStatus('connecting', '正在连接 ' + host + ':' + port + '...');

  const proto = location.protocol === 'https:' ? 'wss' : 'ws';
  ws = new WebSocket(proto + '://' + location.host + '/ws');

  ws.onopen = () => ws.send(JSON.stringify({ type: 'connect', host, port, username, password, private_key, passphrase }));

  ws.onmessage = (e) => {
    const msg = JSON.parse(e.data);
    if (msg.type === 'connected') {
      setStatus('connected', '已连接: ' + username + '@' + host + ':' + port);
      document.getElementById('terminal-title').textContent = username + '@' + host;
      document.getElementById('terminal-card').classList.add('visible');
      initTerm();
      setTimeout(() => {
        if (ws && ws.readyState === WebSocket.OPEN && term)
          ws.send(JSON.stringify({ type: 'resize', rows: term.rows, cols: term.cols }));
        if (term) term.focus();
      }, 100);
    } else if (msg.type === 'output') {
      if (term) term.write(msg.data);
    } else if (msg.type === 'error') {
      setStatus('error', '错误: ' + msg.data);
      document.getElementById('btn-connect').disabled = false;
      if (ws) { ws.close(); ws = null; }
    } else if (msg.type === 'closed') {
      setStatus('error', '连接已关闭');
      document.getElementById('btn-connect').disabled = false;
      ws = null;
    }
  };

  ws.onerror = () => {
    setStatus('error', '连接失败');
    document.getElementById('btn-connect').disabled = false;
    ws = null;
  };

  ws.onclose = () => {
    if (document.getElementById('status-dot').classList.contains('connected'))
      setStatus('error', '连接已断开');
    document.getElementById('btn-connect').disabled = false;
    ws = null;
  };
}

function disconnect() {
  if (ws) { ws.close(); ws = null; }
  setStatus('error', '已主动断开连接');
  document.getElementById('btn-connect').disabled = false;
  document.getElementById('terminal-card').classList.remove('visible');
  if (term) { term.dispose(); term = null; }
}
</script>
</body>
</html>`
