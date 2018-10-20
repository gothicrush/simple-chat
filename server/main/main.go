package main

import (
	"encoding/json"
	"github.com/gothicrush/simple-chat/common/message"
	"github.com/gothicrush/simple-chat/utils"
	"log"
	"net"
)

func main() {

	log.Println("server start...")

	listener, err := net.Listen("tcp", "0.0.0.0:27721")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Server accept connection failed: ",err)
			continue
		}

		go processConnection(conn)
	}
}

func processConnection(conn net.Conn) {

	defer conn.Close()

	msgByte, err := utils.ReadPkg(conn)

	if err != nil {
		log.Println("processConnection：", err)
		return
	}

	msg, err := utils.BytesToMessage(msgByte)

	if err != nil {
		log.Println("processConnection：", err)
		return
	}

	err = ServerProcessMessage(conn, msg)

	if err != nil {
		log.Println("serverProcessMessage(conn, &msg) error", err)
		return
	}
}

func ServerProcessMessage(conn net.Conn, msg *message.Message) error {

	switch msg.Type {
	case message.TypeOfLoginMsg:
		var loginMsg message.MsgLogin

	    err := json.Unmarshal([]byte(msg.Data), &loginMsg)
	    if err != nil {
	    	return err
		}

	    var loginResMsg message.MsgLoginRes

	    if loginMsg.UserId == "abc" && loginMsg.Password == "123" {
	    	loginResMsg.Code = 200
		} else {
			loginResMsg.Code = 500
			loginResMsg.ErrorInfo = "账号密码出错"
		}

	    data, err := json.Marshal(loginResMsg)

	    if err != nil {
	    	return err
		}

	    var newpkg message.Message

		newpkg.Data = string(data)
		newpkg.Type = message.TypeOfLoginResMsg

		pkg, err := json.Marshal(newpkg)

		if err != nil {
			return err
		}

	    err = utils.WritePkg(conn, pkg)

	    return err
	}

	return nil
}