# Rust语言以及课程介绍


学习视频链接，点击跳转播放。

{{< bilibili av78062009 >}} 

## 环境搭建

### 安装

​Windows 系统​：

访问 Rust 官网，下载 `rustup-init.exe`。

运行安装程序，选择默认选项（按回车）。安装过程会自动配置环境变量，需重启终端生效。

​依赖项​：安装过程中需勾选 ​​“Visual Studio C++ Build Tools”​​（包含MSVC编译器）。

​macOS/Linux 系统​：

打开终端，执行以下命令：

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

安装完成后重启终端，或运行 `source $HOME/.cargo/env` 加载环境变量。
​依赖项​：

macOS：安装 Xcode 命令行工具：`xcode-select --install`
Ubuntu/Debian：安装编译工具链：`sudo apt install build-essential`


### Cargo 介绍







---

> 作者: [hao](https://github.com/haochan1996)  
> URL: http://localhost:1313/rust/96e9d38d-4156-45d0-942f-2c7c1c25e1aa/  

