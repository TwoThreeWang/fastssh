<div align="center">
<h2 align="center">FastSSH</h2>
  <p align="center">
    一个简单、轻量、高效管理服务器连接的命令行利器！
    <br />
    <a href="https://github.com/TwoThreeWang/fastssh/wiki"><strong>探索本项目的文档 »</strong></a>
    <br />
    <br />
    <a href="https://github.com/TwoThreeWang/fastssh/issues">报告Bug</a>
    ·
    <a href="https://github.com/TwoThreeWang/">提出新特性</a>
  </p>
  <img src="https://cdn.wangtwothree.com/imgur/QJna1jH.png" alt="Logo">
</div>


该命令行工具用于管理您的服务器连接信息。通过简单的命令行指令，您可以轻松地添加新的服务器连接（包括主机地址、端口、用户名等），删除不再需要的连接，并快速建立与服务器的SSH连接。无需繁琐的配置文件或图形界面操作，一切尽在命令行掌控。

## 功能特性

*   **快速连接：** 通过预先配置的连接信息，一键连接服务器。
*   **便捷管理：** 在命令行中添加和删除服务器连接信息。
*   **简洁易用：** 不需要复杂的命令，一看就懂易于上手。

### 上手指南

1. 下载打包好的可执行文件：https://github.com/TwoThreeWang/fastssh/releases
2. 将文件路径添加到环境变量（可选）
3. 打开终端工具（如果没有添加环境变量，需要在命令行中导航到 fastssh 所在目录）
4. 运行 FastSSH
```
MAC 或者 Linux 直接 ./fastssh 即可
Windows 执行 ./fastssh.exe
```

程序内使用方向键选择，回车确定，Ctrl+C 退出程序

### 开发

1. 下载源代码
2. 执行 go mod tidy 安装包文件
3. 运行 main.go 文件

### 文件目录说明

```
filetree 
├── LICENSE.txt
├── README.md
├── /internal/
│  ├── clipboard.go  // 剪贴板操作函数
│  ├── server.go  // 配置文件操作函数
│  ├── ssh.go  // ssh 连接函数
│  └── ui.go  // 程序主界面
├── main.go
└── go.mod
```



