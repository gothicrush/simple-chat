package main

import (
	"fmt"
	"github.com/gothicrush/simple-chat/client/processor"
	"os"
)

func main() {
	// 接收用户的选择
	var key int

	// 判断是否继续显示菜单
	var loop bool = true

	for loop {
		fmt.Println("---------欢迎登录多人聊天系统---------")
		fmt.Println("1.登录聊天室")
		fmt.Println("2.注册用户")
		fmt.Println("3.退出系统")
		fmt.Println("请选择 1-3")

		// 获取用户的选择
		fmt.Scanln(&key)

		switch key {
		case 1, 2, 3:
		    loop = false
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	// 根据用户输入，显示新的提示信息
	if key == 1 {

		var userid string
		var password string

		fmt.Println("欢迎登录聊天室")

		// 用户输入账号id和密码
		fmt.Println("请输入用户id号")
		fmt.Scanln(&userid)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&password)

		_, err := processor.LoginProcessor(userid, password)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			for {
				fmt.Println("---------欢迎登录多人聊天系统---------")
				fmt.Println("1.显示在线用户列表")
				fmt.Println("2.发送消息")
				fmt.Println("3.信息列表")
				fmt.Println("4.退出程序")
				fmt.Println("请选择 1-4")

				var key int

				fmt.Scanln(&key)

				switch key {
				case 1:
					fmt.Println("TODO::显示在线用户列表")
					os.Exit(0)
				case 2:
					fmt.Println("TODO::发送消息")
					os.Exit(0)
				case 3:
					fmt.Println("TODO::信息列表")
					os.Exit(0)
				case 4:
					fmt.Println("退出程序")
				    os.Exit(0)
				default:
					fmt.Println("输入有误，请重新输入")
				}
			}
		}
	} else if key == 2 {
		fmt.Println("进行用户注册")
	} else {
		fmt.Println("退出系统")
	}
}