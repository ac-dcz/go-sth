## gnet

### 问题

好像在 Windows 环境下，还是一个 Connection 对应一个 goroutine 的模式

具体可以看下面代码：

```go
func (eng *engine) listenStream(ln net.Listener) (err error) {
	if eng.opts.LockOSThread {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()
	}

	defer func() { eng.shutdown(err) }()

	for {
		// Accept TCP socket.
		tc, e := ln.Accept()
		if e != nil {
			err = e
			if atomic.LoadInt32(&eng.beingShutdown) == 0 {
				eng.opts.Logger.Errorf("Accept() fails due to error: %v", err)
			} else if errors.Is(err, net.ErrClosed) {
				err = errors.Join(err, errorx.ErrEngineShutdown)
			}
			return
		}
		el := eng.eventLoops.next(tc.RemoteAddr())
		c := newTCPConn(tc, el)
		el.ch <- &openConn{c: c}
		go func(c *conn, tc net.Conn, el *eventloop) {
			var buffer [0x10000]byte
			for {
				n, err := tc.Read(buffer[:])
				if err != nil {
					el.ch <- &netErr{c, err}
					return
				}
				el.ch <- packTCPConn(c, buffer[:n])
			}
		}(c, tc, el)
	}
}
```

并且 `gnet.Conn` 提供的读写方法均不是并发读写安全的。

个人感觉虽然可以结合 Codec 来做消息的编解码，但是不够优雅。我的意思是，无法根据业务的需要处理消息单元进行个性化开发

### 参考

- [gnet-pangjf2000](https://strikefreedom.top/archives/go-event-loop-networking-library-gnet)

```

```
