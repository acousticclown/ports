# Ports 🔌

A fast, minimal, and beautiful CLI to **list**, **check**, and **kill** running ports on your system.

Built with Go. Distributed via Homebrew. Open-source.

---

## ✨ Features

* 🚀 List all running/listening ports
* 🔍 Check if a specific port is in use
* ❌ Kill a process by port number
* 🎨 Clean, colored terminal output
* ⚡ Lightweight & fast
* 🍺 Installable via Homebrew

---

## 📦 Installation

### macOS (Recommended – Homebrew)

```bash
brew tap acousticclown/ports
brew install ports
```

Verify installation:

```bash
ports
```

---

### Manual Install (Build from source)

```bash
git clone https://github.com/acousticclown/ports.git
cd ports
go build -o ports
sudo mv ports /usr/local/bin
```

---

## 🚀 Usage

### List all running ports

```bash
ports
```

Example output:

```text
🚀 Running Ports

PROCESS      PID      PROTO    PORT
-------------------------------------
node         12345    TCP      3000
postgres     23456    TCP      5432
```

---

### Check a specific port

```bash
ports check 8080
```

Output:

```text
✅ Port 8080 is running (Process: node, PID: 12345)
```

or

```text
❌ Port 8080 is NOT running
```

---

### Kill a port

```bash
ports kill 8080
```

Output:

```text
✅ Port 8080 killed successfully
```

> ⚠️ This command will terminate the process bound to the specified port.

---

## 🛠 Development

### Requirements

* Go 1.20+

### Setup

```bash
go mod tidy
go run .
```

---

## 🧠 How it works

Ports uses system networking utilities under the hood to inspect active listening ports, map them to processes and PIDs, and present the information in a clean, developer-friendly CLI.

---

## 🧩 Project Structure

```text
.
├── main.go
├── ports/
│   ├── list.go
│   ├── check.go
│   ├── kill.go
│   └── model.go
├── ui/
│   ├── colors.go
│   ├── table.go
│   └── messages.go
└── go.mod
```

---

## 🧪 Supported Platforms

* ✅ macOS
* ✅ Linux
* ❌ Windows (planned)

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome.

1. Fork the repository
2. Create your feature branch
3. Open a pull request

---

## 📄 License

MIT License © 2025

---

## ⭐ Why Ports?

Because remembering `lsof`, `grep`, `awk`, and `kill -9` shouldn’t be a daily ritual.
