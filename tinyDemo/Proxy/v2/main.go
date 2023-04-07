package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	socks5Ver = 0x05
	comBind   = 0x01
	atypIPV4  = 0x01
	atypeHOXT = 0x03
	atypeIPV6 = 0x04
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
	err := auth(reader, conn)
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err)
		return
	}
	log.Println("auth success")
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method)

	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

// curl --socks5 127.0.0.1:10050 -v http://www.qq.com
