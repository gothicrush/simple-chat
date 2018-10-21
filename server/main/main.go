package main

import (
	"github.com/gothicrush/simple-chat/server/processor"
	"log"
	"net"
)

func main() {

	// 提示服务器已经开启
	log.Println("server start...")

	// 创建监听器
	listener, err := net.Listen("tcp", "0.0.0.0:27721")

	// 如果监听器创建失败则报错
	if err != nil {
		panic(err)
	}

	// 延时关闭监听器
	defer listener.Close()

	// 一直进行监听
	for {
		// 接收访问连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Server accept connection failed: ",err)
			continue
		}

		// 开启新协程并将连接传递给总处理器进行处理
		go processor.CentralProcessor(conn)
	}
}