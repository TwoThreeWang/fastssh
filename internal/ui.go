package internal

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os"
	"strings"
)

// addNewServer 添加新服务器逻辑
func addNewServer(servers []Server) {
	var newServer Server

	validateHost := func(input string) error {
		if len(input) == 0 {
			return fmt.Errorf("服务器地址不能为空")
		}
		if !strings.Contains(input, "@") {
			return fmt.Errorf("格式错误，应为用户名@服务器ip，例如：root@192.168.1.1")
		}
		return nil
	}
	hostPrompt := promptui.Prompt{
		Label:    "输入用户名&服务器地址 (例如: root@192.168.1.1)",
		Validate: validateHost,
	}

	var err error
	newServer.Host, err = hostPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}

	parts := strings.SplitN(newServer.Host, "@", 2)
	newServer.User = parts[0]
	newServer.Host = parts[1]

	portPrompt := promptui.Prompt{
		Label: "端口号（可选，默认为 22）",
	}

	newServer.Port, err = portPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}
	if newServer.Port == "" {
		newServer.Port = "22"
	}

	namePrompt := promptui.Prompt{
		Label: "设置服务器别名 (可选，留空则使用服务器地址@用户名)",
	}

	newServer.Name, err = namePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}

	if newServer.Name == "" {
		newServer.Name = newServer.Host
		if newServer.User != "" {
			newServer.Name = newServer.Host + "@" + newServer.User
		}
	}

	loginTypePrompt := promptui.Select{
		Label: "选择登录方式",
		Items: []string{"密码登录", "密钥登录"},
	}

	index, _, err := loginTypePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}

	if index == 0 {
		newServer.LoginType = "password"
		passwordPrompt := promptui.Prompt{
			Label: "输入密码",
			Mask:  '*',
		}
		newServer.Password, err = passwordPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		// 保存加密后的密码
		secretKey := os.Getenv("SECRET_KEY")
		// 加密密码
		newServer.Password, err = Encrypt([]byte(secretKey), newServer.Password)
		if err != nil {
			log.Fatalf("密码加密失败: %v", err)
			return
		}
	} else {
		newServer.LoginType = "private_key"
		privateKeyPathPrompt := promptui.Prompt{
			Label: "输入密钥路径",
			Validate: func(input string) error {
				if _, err := os.Stat(input); os.IsNotExist(err) {
					return fmt.Errorf("密钥文件不存在")
				}
				return nil
			},
		}
		newServer.PrivateKeyPath, err = privateKeyPathPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
	}

	servers = append(servers, newServer)
	err = SaveServers(servers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务器已保存。")
	prompt := promptui.Select{
		Label: "选择操作",
		Items: []string{"返回主菜单", "添加新服务器"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}
	if result == "返回主菜单" {
		CmdUi(servers) // 返回主菜单
	} else if result == "添加新服务器" {
		addNewServer(servers)
	}
}

// selectServer 选择已有服务器
func selectServer(servers []Server) {
	if len(servers) == 0 {
		fmt.Println("没有已保存服务器。")
		// 提示并添加“返回上一页”选项
		addNewServerPrompt := promptui.Select{
			Label: "请选择操作",
			Items: []string{"返回主菜单", "添加新服务器"},
		}
		_, result, err := addNewServerPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		if result == "返回主菜单" {
			CmdUi(servers)
		} else if result == "添加新服务器" {
			addNewServer(servers)
		}
	} else {
		serverStrings := make([]string, len(servers))
		for i, server := range servers {
			serverStrings[i] = server.Name
		}
		serverStrings = append(serverStrings, "返回主菜单")
		serverPrompt := promptui.Select{
			Label: "选择要连接的服务器",
			Items: serverStrings,
		}
		index, _, err := serverPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		if serverStrings[index] == "返回主菜单" {
			CmdUi(servers)
		}
		fmt.Printf("正在连接服务器: %s\n", servers[index].Name)
		if err := SshConnect(servers[index]); err != nil {
			fmt.Println("连接失败:", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func delServer(servers []Server) {
	if len(servers) == 0 {
		fmt.Println("没有可删除的服务器。")
		addNewServerPrompt := promptui.Select{
			Label: "请选择操作",
			Items: []string{"返回主菜单"},
		}
		_, _, err := addNewServerPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		CmdUi(servers)
	} else {
		serverStrings := make([]string, len(servers))
		for i, server := range servers {
			serverStrings[i] = server.Name
		}
		serverStrings = append(serverStrings, "返回主菜单")

		deletePrompt := promptui.Select{
			Label: "选择要删除的服务器",
			Items: serverStrings,
		}
		index, result, err := deletePrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		if result == "返回主菜单" {
			CmdUi(servers)
		}
		confirmPrompt := promptui.Select{
			Label: fmt.Sprintf("确认要删除服务器: %s 吗？(y/n)", servers[index].Name),
			Items: []string{"确认", "取消"},
		}
		_, confirm, err := confirmPrompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}

		if confirm == "确认" {
			// 从切片中删除服务器
			servers = append(servers[:index], servers[index+1:]...)
			err := SaveServers(servers)
			if err != nil {
				fmt.Println("保存配置文件失败:", err)
				os.Exit(1)
			}
			fmt.Println("服务器已删除。")
		} else {
			fmt.Println("取消删除。")
		}
		prompt := promptui.Select{
			Label: "选择操作",
			Items: []string{"返回主菜单", "继续删除"},
		}
		_, result, err = prompt.Run()
		if err != nil {
			log.Fatalf("Prompt failed: %v", err)
			return
		}
		if result == "返回主菜单" {
			CmdUi(servers) // 返回主菜单
		} else if result == "继续删除" {
			delServer(servers)
		}
	}
}

func CmdUi(servers []Server) {
	// 主菜单
	prompt := promptui.Select{
		Label: "选择操作",
		Items: []string{"连接服务器", "添加新服务器", "删除服务器"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v", err)
		return
	}
	// 主菜单选项处理
	switch result {
	case "连接服务器":
		selectServer(servers)
	case "添加新服务器":
		addNewServer(servers)
	case "删除服务器":
		delServer(servers)
	}
}
