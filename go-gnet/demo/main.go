package main

import (
	"fmt"
	"log"

	"github.com/panjf2000/gnet/v2"
)

type EchoHandler struct {
	gnet.EventHandler
}

func (e *EchoHandler) OnBoot(engine gnet.Engine) (action gnet.Action) {
	log.Printf("Engine Boot Cur-Connection nums: %d \n", engine.CountConnections())
	return gnet.None
}

func (e *EchoHandler) OnShutdown(engine gnet.Engine) {
	log.Printf("Engine ShutDown Cur-Connection nums: %d \n", engine.CountConnections())
}

func (e *EchoHandler) OnOpen(conn gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("Accept New Connection %s \n", conn.RemoteAddr().String())
	//you can save conn... to do sth
	return
}
func (e *EchoHandler) OnClose(conn gnet.Conn, err error) (action gnet.Action) {
	log.Printf("DisConnect Connection %s \n", conn.RemoteAddr().String())
	//you can rem conn.. to do sth
	return gnet.None
}

func (e *EchoHandler) OnTraffic(conn gnet.Conn) gnet.Action {
	log.Printf("Traffic Connection %s \n", conn.RemoteAddr().String())
	data := make([]byte, 1024)
	n, err := conn.Read(data)
	if err != nil {
		return gnet.Close
	}
	if _, err := conn.Write(data[:n]); err != nil {
		return gnet.Close
	}
	fmt.Printf("Msg: %s", data[:n])
	return gnet.None
}

func main() {
	fmt.Println(gnet.Run(&EchoHandler{
		EventHandler: &gnet.BuiltinEventEngine{},
	},
		"tcp://127.0.0.1:8080",
	))
}
