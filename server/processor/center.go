package processor

import (
	"github.com/gothicrush/simple-chat/common/message"
	"github.com/gothicrush/simple-chat/common/transfer"
	"log"
	"net"
)

func CentralProcessor(conn net.Conn) {

	// 延时关闭连接
	defer conn.Close()

	// 创建 Transfer
	tr := transfer.Transfer{Conn:conn}

	// 通过 transfer 获取数据包
	acceptPackage, err := tr.AcceptPackage()

	if err != nil {
		log.Println("processor.CentralProcessor Error 001 : ", err)
		return
	}

	// 将数据包转为 message
	msg, err := message.PackageToMessage(acceptPackage)

	if err != nil {
		log.Println("processor.CentralProcessor Error 002 : ", err)
		return
	}

	// 将 message 分派处理
	dispatchProcessor(tr, msg)
}

func dispatchProcessor(tr transfer.Transfer, msg *message.Message) {

	switch msg.Type {
	case message.TypeMessageLogin:
		loginProcessor(tr, msg)
	default:
		log.Println("processor.dispatchProcessor Error 001 : 无匹配的消息类型")
	}
}
