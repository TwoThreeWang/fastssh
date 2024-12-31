package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Server struct {
	Name           string `json:"name,omitempty"`             // 服务器名称（可选，默认为 主机地址_用户名）
	User           string `json:"user,omitempty"`             // 用户名
	Host           string `json:"host"`                       // 主机地址
	Port           string `json:"port,omitempty"`             // 端口号（可选，默认为 22）
	Password       string `json:"password,omitempty"`         // 密码
	PrivateKeyPath string `json:"private_key_path,omitempty"` // 私钥路径
	LoginType      string `json:"login_type,omitempty"`       // 登陆方式（0密码；1密钥）
}

func getConfigPath() string {
	//const configFile = "./servers.json"
	// 获取程序的执行路径
	exePath, err := os.Executable()
	if err != nil {
		fmt.Errorf("failed to get executable path: %v", err)
		return ""
	}

	// 获取程序所在目录
	exeDir := filepath.Dir(exePath)

	// 在程序目录下查找配置文件 servers.json
	return filepath.Join(exeDir, "servers.json")
}

func LoadServers() ([]Server, error) {
	configFile := getConfigPath()
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(configFile)
			if err != nil {
				return nil, fmt.Errorf("创建配置文件失败: %w", err)
			}
			return []Server{}, nil
		}
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}
	if data == nil || len(data) == 0 {
		return []Server{}, nil
	}

	var servers []Server
	err = json.Unmarshal(data, &servers)
	if err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %w", err)
	}

	return servers, nil
}

func SaveServers(servers []Server) error {
	data, err := json.MarshalIndent(servers, "", "    ")
	if err != nil {
		return fmt.Errorf("序列化服务器列表失败: %w", err)
	}
	configFile := getConfigPath()
	err = os.WriteFile(configFile, data, 0644)
	if err != nil {
		return fmt.Errorf("保存配置文件失败: %w", err)
	}

	return nil
}
