## something about go

### 问题

- *netpool*是边沿触发还是水平触发？
  根据 `go-net/tcp/tcp_test.go` 应该是水平触发(傻逼了,应该在 Linux 环境上测试)

### 参考

- [go netpoll](https://strikefreedom.top/archives/go-netpoll-io-multiplexing-reactor)
