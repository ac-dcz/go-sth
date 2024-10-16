package tcp_test

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestMethodTrigger(t *testing.T) {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	defer lis.Close()
	wgR := sync.WaitGroup{}
	wgR.Add(1)
	go func() {
		defer wgR.Done()
		conn, err := lis.Accept()
		if err != nil {
			t.Errorf("accept error: %v", err)
			return
		}
		defer conn.Close()
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 1; i++ {
				t.Logf("Server W %d\n", i)
				conn.Write([]byte("helloworld"))
				time.Sleep(time.Second * 5)
			}
		}()
		go func() {
			defer wg.Done()
			for {
				data := make([]byte, 5)
				if n, err := conn.Read(data); err != nil {
					t.Errorf("read error: %v", err)
					return
				} else {
					t.Logf("%s \n", data[:n])
				}
			}
		}()
		wg.Wait()
	}()

	conn, err := net.DialTimeout("tcp", ":8080", time.Second*5)
	if err != nil {
		t.Error(err)
		return
	}
	wgR.Add(1)
	go func() {
		defer wgR.Done()
		for {
			data := make([]byte, 5)
			if n, err := conn.Read(data); err != nil {
				t.Errorf("read error: %v", err)
				return
			} else {
				t.Logf("%s \n", data[:n])
			}
		}
	}()
	for i := 0; i < 1; i++ {
		t.Logf("Client W %d\n", i)
		conn.Write([]byte("helloworld"))
		time.Sleep(time.Second * 5)
	}
	conn.Close()
	wgR.Wait()
	time.Sleep(time.Second * 1)
}
