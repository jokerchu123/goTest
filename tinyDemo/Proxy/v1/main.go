package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:10050") //设置监听端口
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept() //接受请求，如果成功返回连接
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client) //启动goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn) //基于连接创建只读缓冲流
	for {
		b, err := reader.ReadByte() //每次读入一个字节
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b}) //将读入字节写入
		if err != nil {
			break
		}

	}
}

// nc 127.0.0.1 10050 与指定端口建立连接
