package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type TagResponse struct {
	TagName string `json:"tag_name"` // 将 JSON 中的 tag_name 映射到 Go 结构体的 TagName 字段
}

// Version 当前程序版本号
const Version = "1.4"

func CheckUpdate() {
	fmt.Println("正在检查是否有更新...")
	// 发送 GET 请求获取最新的版本号
	url := "https://api.github.com/repos/TwoThreeWang/fastssh/releases/latest"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error sending GET request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 解析 JSON 数据
	var responseData TagResponse
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	if responseData.TagName != Version {
		fmt.Println("程序有更新啦！当前版本：" + Version + "，最新版本：" + responseData.TagName)
		fmt.Println("去下载最新版：https://github.com/TwoThreeWang/fastssh/releases/latest")
	} else {
		fmt.Println("已经是最新版啦！")
	}
}
