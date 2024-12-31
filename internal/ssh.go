package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func SshConnect(server Server) error {
	var cmd *exec.Cmd
	if server.LoginType == "password" {
		// 拼接SSH登录命令
		sshCommand := fmt.Sprintf("ssh -p %s %s@%s", server.Port, server.User, server.Host)
		// 创建一个*exec.Cmd结构体来表示要执行的命令
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("PowerShell", "-c", sshCommand)
		default:
			cmd = exec.Command("bash", "-c", sshCommand)
		}
		// 将密码粘贴到剪贴板
		err := copyToClipboard(server.Password)
		if err != nil {
			fmt.Printf("将密码复制到剪贴板失败: %v\n", err)
			return err
		} else {
			fmt.Println("密码已成功复制到剪贴板")
		}
	} else {
		// 拼接SSH登录命令
		sshCommand := fmt.Sprintf("ssh -p %s -i %s %s@%s", server.Port, server.PrivateKeyPath, server.User, server.Host)
		// 创建一个*exec.Cmd结构体来表示要执行的命令
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("PowerShell", "-c", sshCommand)
		default:
			cmd = exec.Command("bash", "-c", sshCommand)
		}
	}
	fmt.Println("服务器正在登录中...")
	fmt.Println("----------------------------------")
	// 将命令的标准输入、输出、错误输出关联到当前进程的对应部分
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// 执行命令，如果有错误则打印错误信息并退出
	return cmd.Run()
}
