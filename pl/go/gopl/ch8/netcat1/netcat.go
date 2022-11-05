package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// 从连接中读取数据，并将读到的内容写到标准输出中
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

/*
	启动服务端clock1
	让我们同时运行两个客户端来进行一个测试，这里可以开两个终端窗口，
	下面上边的是其中的一个的输出，下边的是另一个的输出：
	13:58:54
	13:58:55
	13:58:56
	^C
	13:58:57
	13:58:58
	13:58:59
	^C

	第二个客户端必须等待第一个客户端完成工作，这样服务端才能继续向后执行；
	因为我们这里的服务器程序同一时间只能处理一个客户端连接。
 	---------------------------------------------------------------
	启动服务端clock2
	同时运行两个客户端
	11:37:37
	11:38:38
	11:39:39
	^C
	11:37:37
	11:38:38
	11:39:39
	^C
	多个客户端可以同时接收到时间,实现并发。

*/
