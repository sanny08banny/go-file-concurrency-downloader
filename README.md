# Simple Go HTTP File Server

This is a basic HTTP file server written in Go. It serves static files from a local directory and supports multiple concurrent downloads (e.g., 2 to 20 clients at once) thanks to Go’s built-in concurrency.

---

## 🚀 Features

- Serves files from a specified folder via HTTP.
- Handles multiple clients concurrently.
- Easy to set up and run.
- No external dependencies.

---

## 📁 Directory Structure

```
project/
├── main.go
├── shared-files/
│   └── _.jpeg  (or any file you want to serve)
```

---

## 🛠️ Setup & Run

1. Create a directory with files to share:

```bash
mkdir shared-files
cp yourfile.jpeg shared-files/_.jpeg
```

2. Create `main.go`:

```go
package main

import (
    "log"
    "net/http"
)

const (
    port    = ":8080"
    fileDir = "./shared-files"
)

func main() {
    fs := http.FileServer(http.Dir(fileDir))
    http.Handle("/", fs)

    log.Printf("Serving %s on HTTP port %s
", fileDir, port)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

3. Run the server:

```bash
go run main.go
```

4. Visit [http://localhost:8080](http://localhost:8080) or download directly:

```bash
curl -O http://localhost:8080/_.jpeg
```

---

## 🧪 How to Test Concurrent Downloads

### 🔁 Test with Curl

This will simulate 10 concurrent downloads:

```bash
for i in {1..10}; do
  curl -o "download_$i.jpeg" http://localhost:8080/_.jpeg &
done
wait
```

Check the results:

```bash
ls -lh download_*.jpeg
md5sum download_*.jpeg
```

All files should be the same size and have the same checksum.

---

### ⚡ Benchmark with `hey` (Optional)

Install `hey`:

```bash
go install github.com/rakyll/hey@latest
```

Run a benchmark with 20 concurrent clients:

```bash
hey -n 100 -c 20 http://localhost:8080/_.jpeg
```

---

## 🛡️ Security Note

This server does **not** include:

- Authentication
- Rate limiting
- HTTPS

Only use it in trusted environments or local networks. For production, consider enhancements.

---

## 📄 License

MIT License
