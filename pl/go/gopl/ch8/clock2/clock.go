package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Listen函数创建一个net.Listener的对象，这个对象会监听一个网络端口上到来的连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Accept方法会直接阻塞，直到一个新的连接被创建，然后会返回一个net.Conn对象来表示这个连接。
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 让每一次handleConn的调用都进入一个独立的goroutine。同时处理连接,支持并发。
	}
}

// handleConn 函数会处理一个完整的客户端连接。
// 在一个for死循环中，用time.Now()获取当前时刻，然后写到客户端。
func handleConn(c net.Conn) {
	// defer调用关闭服务器的连接，然后返回到主函数，继续等待下一个连接请求。
	defer c.Close()
	for {
		// net.Conn实现了io.Writer接口，我们可以直接向其写入内容。
		_, err := io.WriteString(c, time.Now().Format("15:05:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
