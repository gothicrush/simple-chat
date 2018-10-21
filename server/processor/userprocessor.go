package processor

import (
	"github.com/gothicrush/simple-chat/common/message"
	"github.com/gothicrush/simple-chat/common/transfer"
	"log"
)

func loginProcessor(tr transfer.Transfer, msg *message.Message) {

	messageLogin, err := message.PackageToMessageLogin([]byte(msg.Data))

	if err != nil {
		log.Println("processor.loginProcessor Error 001 : ", err)
		return
	}

	// 创建MessageLoginResponse结构体
	var messageLoginResponse message.MessageLoginResponse

	if messageLogin.UserId == "abc" && messageLogin.Password == "123" {
		messageLoginResponse.Code = 200
	} else {
		messageLoginResponse.Code = 500
		messageLoginResponse.Info = "账号密码出错"
	}

	// 序列化MessageLoginResponse结构体
	data, err := message.MessageLoginResponseToPackage(&messageLoginResponse)

	if err != nil {
		log.Println("processor.loginProcessor Error 002 : ", err)
		return
	}

	// 创建 Message，接收序列化后的MessageLoginResponse
	var messageReturn message.Message

	messageReturn.Data = string(data)
	messageReturn.Type = message.TypeMessageLoginResponse

	// 序列化 messageReturn
	messageReturnPackage, err := message.MessageToPackage(&messageReturn)

	if err != nil {
		log.Println("processor.loginProcessor Error 003 : ", err)
		return
	}

	// 发送 messageReturnPackage
	err = tr.SendPackage(messageReturnPackage)

	if err != nil {
		log.Println("processor.loginProcessor Error 004 : ", err)
		return
	}
}
