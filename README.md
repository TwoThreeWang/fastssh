<div align="center">
  <img src="https://cdn.wangtwothree.com/imgur/QJna1jH.png" alt="FastSSH一个简单、轻量、高效管理服务器连接的命令行利器！">
  <h1 align="center">FastSSH</h1>
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
</div>

FastSSH 是一个轻量、简单的命令行工具，用于管理您的服务器连接信息。通过简单的指令，您可以轻松地添加新的服务器连接（包括主机地址、端口、用户名等），删除不再需要的连接，并快速建立与服务器的 SSH 连接。无需繁琐的配置文件或图形界面操作，一切尽在命令行掌控。

## 功能特性

*   **快速连接：** 通过预先配置的连接信息，一键连接服务器。
*   **便捷管理：** 在命令行中快速添加和删除服务器连接信息。
*   **安全易用：** 软件功能无需联网，操作皆在本地；不需要复杂的命令，一看就懂易于上手。
*   **轻量简洁：** 软件打包后仅 5M，一个可执行文件即可运行。

## 为什么开发 FastSSH

前段时间电脑重装，安装 SSH 工具的时候，想着看看有没有界面美观、轻量好用的 SSH 软件，之前一直使用的是 Tabby，界面很喜欢，但是打开速度有点慢，而且有点占用资源，测试好几款之后目前在使用 iTerm2。

在测试过程中，由于需要下载切换多款 SSH 工具，每次都得配置服务器连接，而且有些工具配置还比较复杂，比如 iTerm2，突发奇想，能不能做一个命令行工具，可以保存服务器的连接信息，并且直接连接，这样不管用哪一款 SSH 工具，都不需要配置了，直接输入一个简短的命令就可以登录、管理已经保存的服务器信息。

于是说做就做，FastSSH 由此诞生，一个简单、轻量、高效管理服务器连接的命令行利器！

目前第一个版本，在 Windows 和 Mac 上测试是没有问题的，只测试了账号密码登录的方式，密钥方式还没测试，有感兴趣的小伙伴可以下载试一下，有任何问题可以 issues 反馈。

**为什么登录的时候还需要粘贴密码？**

起初的想法是自动粘贴密码登录的，但是目前了解到的方式都是要通过第三方软件实现，比如 sshpass，为了保持工具的轻量化，还是想尽量使用系统原生的功能，目前想到的方法就是登录的时候工具自动将密码放到剪贴板里，这样就可以直接粘贴密码，回车登录就可以了，如果有大神有更好的方法，欢迎交流。

### 上手指南

1. 下载打包好的可执行文件：https://github.com/TwoThreeWang/fastssh/releases/latest
2. 将文件路径添加到环境变量（可选）
3. 打开终端工具（如果没有添加环境变量，需要在命令行中导航到 fastssh 所在目录）
4. 运行 FastSSH
```
MAC 或者 Linux 直接运行 fastssh 即可
Windows 执行 fastssh.exe
```

程序内使用方向键选择，回车确定，Ctrl+C 退出程序

![FastSSH一个简单、轻量、高效管理服务器连接的命令行利器！](https://cdn.wangtwothree.com/imgur/b5H0l14.png)

### Mac 设置环境变量

#### 临时添加到路径

你可以在终端会话中临时将工具所在的目录添加到 PATH 环境变量中。假设你的工具位于 /Users/yourusername/fastssh/ 目录下，你可以在终端中运行以下命令：
```
export PATH="$PATH:/Users/yourusername/fastssh/"
```
这样，在当前终端会话中，你就可以在任意目录下调用你的工具了。但是，这种方法只对当前终端会话有效，一旦关闭终端，设置就会丢失。

#### 永久添加到路径

**1. 确定你使用的 Shell:**

Mac 默认的 Shell 已经从 Bash 切换到了 Zsh。你需要确定你当前使用的是哪个 Shell，这会影响你需要编辑的文件。

*   打开终端。
*   输入 `echo $SHELL` 并回车。
*   如果输出 `/bin/bash` 或类似内容，则你使用的是 Bash。
*   如果输出 `/bin/zsh` 或类似内容，则你使用的是 Zsh。

**2. 编辑相应的配置文件:**

根据你使用的 Shell，选择需要编辑的文件：

*   **Bash:** 编辑 `~/.bash_profile` 文件，在终端输入 `vi ~/.bash_profile` 并回车。
*   **Zsh:** 编辑 `~/.zshrc` 文件，在终端输入 `vi ~/.zshrc` 并回车。

**3. 添加环境变量:**

在打开的文件中，添加以下行，将 `/path/to/your/directory/` 替换为你想要添加的目录的实际路径：

```bash
export PATH="$PATH:/path/to/your/directory/"
```

*   `export PATH`: 这是设置环境变量的命令。
*   `/path/to/your/directory/`: 这是你想要添加到 `PATH` 的目录的完整路径。例如：`/usr/local/bin`、`/Users/你的用户名/bin`、`/Applications/MyApplication/Contents/MacOS` 等。
*   `:$PATH`: 这会将你添加的目录放在 `PATH` 的前面，这样系统会优先搜索这个目录。保留 `$PATH` 可以确保原有的环境变量不会丢失。

**示例：**

假设你想将 `/Users/myuser/fastssh` 目录添加到 `PATH`，则应添加以下行：

```bash
export PATH="/Users/myuser/fastssh/:$PATH"
```

**4. 保存并关闭文件:**

*   **nano:** 按下 `Ctrl + X`，然后按 `Y` 保存，最后按回车退出。
*   **vim & vi:** 按下 `Esc` 键，输入 `:wq` 并回车保存并退出。
*   **TextEdit:** 保存文件（确保格式为纯文本），然后关闭。

**5. 使更改生效:**

你需要重新加载配置文件才能使更改生效。在终端输入以下命令：

*   **Bash:** `source ~/.bash_profile`
*   **Zsh:** `source ~/.zshrc`

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
│  ├── encryption.go  // 加解密函数
│  ├── server.go  // 配置文件操作函数
│  ├── ssh.go  // ssh 连接函数
│  └── ui.go  // 程序主界面
├── main.go
└── go.mod
```

### 完全卸载

1. 删除程序所在目录
2. 如果有设置环境变量，删除相关环境变量

