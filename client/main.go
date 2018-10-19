package main

import (
	"fmt"
	"github.com/gothicrush/simple-chat/client/login"
)

var userid string
var password string

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

		fmt.Println("登录聊天室")
		// 用户要登录
		fmt.Println("请输入用户id号")
		fmt.Scanln(&userid)
		fmt.Println("请输入用户密码")
		fmt.Scanln(&password)

		err := login.Login(userid, password)

		if err != nil {
			fmt.Println("登录失败")
		} else {
			fmt.Println("登录成功")
		}
	} else if key == 2 {
		fmt.Println("进行用户注册")
	} else {
		fmt.Println("退出系统")
	}
}