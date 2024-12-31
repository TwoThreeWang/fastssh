package internal

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// copyToClipboard 将内容复制到剪贴板
func copyToClipboard(content string) error {
	var cmd *exec.Cmd

	// 根据操作系统选择不同的命令
	switch runtime.GOOS {
	case "windows":
		// Windows 使用 clip 命令
		cmd = exec.Command("clip")
		cmd.Stdin = strings.NewReader(content)
	case "darwin":
		// macOS 使用 pbcopy 命令
		cmd = exec.Command("pbcopy")
		cmd.Stdin = strings.NewReader(content)
	case "linux":
		// Linux 使用 xclip 或 xsel 命令
		cmd = exec.Command("xclip", "-selection", "clipboard")
		cmd.Stdin = strings.NewReader(content)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	// 执行命令
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to copy to clipboard: %v", err)
	}

	return nil
}
