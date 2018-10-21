package processor

import (
	"errors"
	"github.com/gothicrush/simple-chat/common/message"
	"github.com/gothicrush/simple-chat/common/transfer"
	"log"
	"net"
)

func LoginProcessor(userid string, password string) (bool, error){

	// 连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:27721")

	if err != nil {
		log.Println("processor.LoginProcessor Error 001 : ", err)
		return false, err
	}

	// 延时关闭连接
	defer conn.Close()

	// 创建传输工具
	tr := transfer.Transfer{Conn: conn}

	// 准备 Message
	var messageSend message.Message
	messageSend.Type = message.TypeMessageLogin

	// 准备 MessageLogin
	var messageLogin message.MessageLogin
	messageLogin.UserId = userid
	messageLogin.Password = password

	// 将 MessageLogin 序列化
	messageLoginPackage, err := message.MessageLoginToPackage(&messageLogin)

	if err != nil {
		log.Println("processor.LoginProcessor Error 002 : ", err)
		return false, err
	}

	// 将 messageLoginPackage 填充到 msg.Data
	messageSend.Data = string(messageLoginPackage)

	// 将 Message 序列化
	messageSendPackage, err := message.MessageToPackage(&messageSend)

	if err != nil {
		log.Println("processor.LoginProcessor Error 003 : ", err)
		return false, err
	}

	// 发送 Message
	err = tr.SendPackage(messageSendPackage)

	if err != nil {
		log.Println("processor.LoginProcessor Error 004 : ", err)
		return false, err
	}

	///////////////////////////////////////////////////////////////////////

	// 接收响应
	messageAcceptPackage, err := tr.AcceptPackage()

	if err != nil {
		log.Println("processor.LoginProcessor Error 005 : ", err)
		return false, err
	}

	// 将获取的 Message 数据包转为 Message
	messageAccept, err := message.PackageToMessage(messageAcceptPackage)

	if err != nil {
		log.Println("processor.LoginProcessor Error 006 : ", err)
		return false, err
	}

	// 获取 MessageLoginResponse
	messageLoginResponse, err := message.PackageToMessageLoginResponse([]byte(messageAccept.Data))

	if err != nil {
		log.Println("processor.LoginProcessor Error 007 : ", err)
		return false, err
	}

	if messageLoginResponse.Code != 200 {
		log.Println("processor.LoginProcessor Error 008: ", messageLoginResponse.Info)
		return false, errors.New(messageLoginResponse.Info)
	}

	return true, nil
}
