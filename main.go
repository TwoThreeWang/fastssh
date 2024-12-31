package main

import (
	"fmt"
	"temp/internal"
)

func main() {
	fmt.Println("--------------------------------------------------------------------\n    __  __     ____         ______           __     __________ __  __\n   / / / /__  / / /___     / ____/___ ______/ /_   / ___/ ___// / / /\n  / /_/ / _ \\/ / / __ \\   / /_  / __ `/ ___/ __/   \\__ \\\\__ \\/ /_/ / \n / __  /  __/ / / /_/ /  / __/ / /_/ (__  ) /_    ___/ /__/ / __  /  \n/_/ /_/\\___/_/_/\\____/  /_/    \\__,_/____/\\__/   /____/____/_/ /_/   \n\nAuthor: WangTwoThree\nGithub: https://github.com/TwoThreeWang/fastssh \n--------------------------------------------------------------------\n欢迎使用，将程序添加到系统环境变量，就可以在任意目录下直接调用 FastSSH 登陆服务器\n")
	servers, err := internal.LoadServers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 主程序
	internal.CmdUi(servers)
}
