package login

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gothicrush/simple-chat/common/message"
	"github.com/gothicrush/simple-chat/utils"
	"log"
	"net"
)

func Login(userid string, password string) (err error){

	// 1.连接到服务器
	conn, err := net.Dial("tcp", "0.0.0.0:27721")

	if err != nil {
		return err
	}

	defer conn.Close()

	// 2.准备 Message
	var msg message.Message
	msg.Type = message.TypeOfLoginMsg

	// 3.创建LoginMsg
	var loginMsg message.MsgLogin
	loginMsg.UesrName = "username"
	loginMsg.UserId = userid
	loginMsg.Password = password

	// 4.将LoginMsg序列化
	data,err := json.Marshal(loginMsg)
	if err != nil {
		return err
	}

	// 5.将 LoginMsg序列化填充到 msg.Data
	msg.Data = string(data)

	// 6.将msg序列化
	pkg, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = utils.WritePkg(conn, pkg)

	if err != nil {
		return err
	}

	// 10.接收响应
	msgByte, err := utils.ReadPkg(conn)

	if err != nil {
		log.Println("Login(userid string, password string) Error", err)
		return err
	}

	msgRec, err := utils.BytesToMessage(msgByte)

	if err != nil {
		log.Println("Login(userid string, password string) Error", err)
		return err
	}

	// 11. 接收实际MsgLoginRes
	var loginResMsg message.MsgLoginRes

	err = json.Unmarshal([]byte(msgRec.Data), &loginResMsg)

	if err != nil {
		log.Println("Login(userid string, password string) Error", err)
		return err
	}

	if loginResMsg.Code != 200 {
		fmt.Println(loginResMsg.ErrorInfo)
		return errors.New("登录失败")
	}

	return nil
}
