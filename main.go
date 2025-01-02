package main

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
	"temp/internal"
)

func getEnv() bool {
	// 检查 .env 文件是否存在
	// 获取程序的执行路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Errorf("failed to get executable path: %v", err)
		return false
	}

	// 获取程序所在目录
	exeDir := filepath.Dir(exePath)

	// 在程序目录下查找配置文件 servers.json
	envPath := filepath.Join(exeDir, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		fmt.Println(".env 文件不存在，是否创建？ (y/n)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "y" {
			fmt.Println("请设置一个私人密钥，用于加密服务器密码")
			keyReader := bufio.NewReader(os.Stdin)
			keyInput, _ := keyReader.ReadString('\n')
			keyInput = strings.TrimSpace(keyInput)
			// 创建 .env 文件并写入示例内容
			err := os.WriteFile(envPath, []byte("SECRET_KEY="+keyInput), 0644)
			if err != nil {
				fmt.Println("创建 .env 文件失败:", err)
				return false
			}
			fmt.Println(".env 文件已创建。")
		} else {
			fmt.Println("未创建 .env 文件，程序将退出。请创建 .env 文件并设置环境变量。")
			return false
		}
	}

	// 加载 .env 文件
	err = godotenv.Load(envPath)
	if err != nil {
		fmt.Println("加载 .env 文件失败:", err)
		return false
	}
	return true
}

func main() {
	fmt.Println("--------------------------------------------------------------------\n    __  __     ____         ______           __     __________ __  __\n   / / / /__  / / /___     / ____/___ ______/ /_   / ___/ ___// / / /\n  / /_/ / _ \\/ / / __ \\   / /_  / __ `/ ___/ __/   \\__ \\\\__ \\/ /_/ / \n / __  /  __/ / / /_/ /  / __/ / /_/ (__  ) /_    ___/ /__/ / __  /  \n/_/ /_/\\___/_/_/\\____/  /_/    \\__,_/____/\\__/   /____/____/_/ /_/   \n\nAuthor: WangTwoThree\nGithub: https://github.com/TwoThreeWang/fastssh \n--------------------------------------------------------------------\n欢迎使用，将程序添加到系统环境变量，就可以在任意目录下直接调用 FastSSH 登陆服务器\n")
	servers, err := internal.LoadServers()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// 加载配置文件
	if !getEnv() {
		return
	}
	// 主程序
	internal.CmdUi(servers)
}
