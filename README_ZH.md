# govnr

使用 `govnr` 来启动受监控的 goroutine。

本包提供以下功能：
* `Once()` - 启动一个 goroutine，并记录未捕获的 panic。
* `Forever()` - 启动一个 goroutine，如果发生 panic，会记录错误并重新启动（除非上下文已取消）。
* `Recover()` - 在当前 goroutine 中运行函数，自动捕获 panic 并记录，避免程序崩溃。

[文档](https://godoc.org/github.com/orbs-network/govnr) 已提供，但仍有改进空间，欢迎提交 PR！

在 [Orbs Network 的 Go 实现](https://github.com/orbs-network/orbs-network-go) 中广泛使用，确保所有后台进程稳定运行。

### 示例用法
```go
type stdoutErrorer struct {}

func (s *stdoutErrorer) Error(err error) {
    fmt.Println(err.Error())
}

errorHandler := &stdoutErrorer{}
ctx, cancel := context.WithCancel(context.Background())

data := make(chan int)
handle := govnr.Forever(ctx, "示例进程", errorHandler, func() {
    for {
        select {
        case i := <-data:
            fmt.Printf("goroutine 收到数据: %d\n", i)
        case <-ctx.Done():
            return
        }
    }
})

supervisor := &govnr.TreeSupervisor{}
supervisor.Supervise(handle)

data <- 3
data <- 2
data <- 1
cancel()

shutdownCtx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
supervisor.WaitUntilShutdown(shutdownCtx)
```  

---

### 功能说明
| 方法 | 作用 |  
|------|------|  
| `Once()` | 启动一个受监控的 goroutine，仅运行一次，崩溃时记录错误 |  
| `Forever()` | 启动持久化 goroutine，崩溃后自动重启（除非上下文取消） |  
| `Recover()` | 在当前 goroutine 中安全执行函数，自动恢复 panic |  

适用于需要 **长期运行** 或 **高稳定性** 的后台任务，如：
- 网络服务守护
- 定时任务调度
- 数据处理流水线

🚀 **简单易用，稳定可靠！**