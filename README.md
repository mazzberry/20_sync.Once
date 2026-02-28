<div align="center">

# 🔄 Concurrency in Go — `sync.Once`

<img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
<img src="https://img.shields.io/badge/Pattern-Singleton-blueviolet?style=for-the-badge"/>
<img src="https://img.shields.io/badge/Concurrency-sync.Once-green?style=for-the-badge"/>

<br/>
<p>Exploring the <strong>Singleton Pattern</strong> in Go using <code>sync.Once</code> and comparing it with an unsafe approach</p>

</div>

---

## 📌 Overview

This project demonstrates how to safely initialize a single instance **exactly once** in a **concurrent** Go program.  
Two approaches are compared:

| Approach | Thread-Safe | Description |
|----------|-------------|-------------|
| `GetConfig()` | ❌ Unsafe | Broken Double-Checked Locking with Mutex |
| `GetConfigWithOnce()` | ✅ Safe | Correct usage of `sync.Once` |

---

## 🧩 Code Breakdown

### ❌ The Bad Way — Broken Double-Checked Locking

```go
func GetConfig() *Config {
    if config != nil {        // ← First check happens WITHOUT a lock (Race Condition!)
        mx.Lock()
        defer mx.Unlock()
        if config == nil {
            config = &Config{}
        }
    }
    return config
}
```

> **⚠️ Problem:** The first check (`config != nil`) runs without holding the lock.  
> In a concurrent environment, two goroutines can pass this check simultaneously,  
> causing `config` to be initialized more than once.

---

### ✅ The Right Way — `sync.Once`

```go
var once sync.Once

func GetConfigWithOnce() *Config {
    once.Do(func() {
        config = &Config{}
        fmt.Println("creating config") // printed only once
    })
    return config
}
```

> **✅ Why is this safe?**  
> `sync.Once` guarantees that the function inside `Do` is executed **exactly once**,  
> even if thousands of goroutines call it simultaneously.

---

## 🚀 Run

```bash
git clone <repo-url>
cd 20_concurrency_syncOnce
go run main.go
```

### 🖥️ Expected Output

```
get config from old instance
0 : 0xc0000b4020
get config from old instance
1 : 0xc0000b4020
...
```

> All 100 iterations return the **same address** — confirming `Config` was created only once.

---

## 💡 Key Takeaways

<table>
<tr>
<td>🔒</td>
<td><strong>sync.Once</strong> uses internal atomic operations — no manual Mutex needed</td>
</tr>
<tr>
<td>⚡</td>
<td>After the first execution, <code>once.Do</code> is essentially a no-op with near-zero overhead</td>
</tr>
<tr>
<td>🎯</td>
<td>Perfect for: Database connections, Config loading, Logger initialization</td>
</tr>
</table>

---

## 📚 References

- [Go Documentation — sync.Once](https://pkg.go.dev/sync#Once)
- [The Go Memory Model](https://go.dev/ref/mem)

---

<div align="center">
<sub>Part of the Go Concurrency learning series</sub>
</div>