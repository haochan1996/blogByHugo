# Rust语言以及课程介绍


令狐大神学习视频链接，点击跳转播放。

[Rust语言编程基础教程](https://www.bilibili.com/video/av78062009?vd_source=939ae5b13ea25e42d7ce7f25bd855603&spm_id_from=333.788.player.switch&p=2)

[Rust 程序设计语言 中文版](https://rustwiki.org/zh-CN/book/title-page.html)

基础学习主要结合上面两个资料。

## 环境搭建

Windows 系统​：

访问 [Rust 官网](https://www.rust-lang.org/zh-CN/tools/install)，下载 `rustup-init.exe`。

运行安装程序，选择默认选项（按回车）。安装过程会自动配置环境变量，需重启终端生效。

依赖项​：安装过程中需勾选 ​​“Visual Studio C++ Build Tools”​​（包含MSVC编译器）。

macOS/Linux 系统​：

打开终端，执行以下命令：

```bash
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
```

安装完成后重启终端，或运行 `source $HOME/.cargo/env` 加载环境变量。
​依赖项​：

macOS：安装 Xcode 命令行工具：`xcode-select --install`
Ubuntu/Debian：安装编译工具链：`sudo apt install build-essential`


## Cargo 介绍

Cargo 是 Rust 官方提供的**构建系统（Build System）和包管理器（Package Manager）**，与 Rust 编译器 `rustc` 深度集成，负责管理项目的全生命周期：

- 包管理：自动下载、编译、链接依赖库（称为 “crates”），解决版本冲突。
- 构建流程：编译源代码，生成可执行文件或库，支持调试与发布模式。
- 项目标准化：统一项目结构、测试框架、文档生成与发布流程，提升协作效率。

> 💡 **定位类比**：
> 类似 JavaScript 的 `npm`、Python 的 `pip`、Java 的 `Maven`。

### Cargo核心功能详解

#### 项目初始化

- 创建新项目：

  ```bash
  cargo new my_project     # 二进制项目
  cargo new my_lib --lib   # 库项目
  ```

  生成标准结构：

  ```markdown
  my_project/
  ├── Cargo.toml    # 项目配置
  ├── src/
  │   └── main.rs   # 入口文件（或 lib.rs 库项目）
  └── .gitignore    # 默认 Git 配置
  ```

#### 依赖管理

- 在`Cargo.toml` 中声明依赖：

  ```toml
  [dependencies]
  serde = "1.0"                   # 指定版本
  tokio = { version = "1.0", features = ["full"] } # 启用特性[6](@ref)
  ```

- Cargo 自动从`crates.io`（Rust 官方包仓库）下载依赖，并生成版本锁文件`Cargo.lock`。

#### 构建与运行

- 调试构建：

  ```bash
  cargo build        # 输出到 target/debug/
  cargo run          # 编译后立即运行
  ```

- 快速检查代码确保其可以编译，但并不产生可执行文件：

```bash
cargo check
```

- 发布构建（优化性能）：

  ```bash
  cargo build --release  # 输出到 target/release/
  ```

#### 测试与文档

- 运行测试：

  ```bash
  cargo test  # 执行所有标记 #[test] 的测试函数
  ```

- 生成文档：

  ```bash
  cargo doc --open  # 自动生成 HTML 文档并打开浏览器
  ```

#### 发布与共享

- 发布项目到`crates.io`：

  ```bash
  cargo publish  # 需提前登录并配置 API Key
  ```

### 在 Cargo 项目中运行指定二进制文件

若项目包含多个入口（如多个 `main.rs`），可通过配置实现：

1. **配置 `Cargo.toml`**：

   在 `Cargo.toml`中添加多个 `[[bin]]`目标：

   ```toml
   [[bin]]
   name = "cli"
   path = "src/cli/main.rs"
   
   [[bin]]
   name = "server"
   path = "src/server/main.rs"
   ```

2. **运行指定目标**：

   ```bash
   cargo run --bin cli   # 运行 cli 程序
   cargo run --bin server # 运行 server 程序
   ```

   此方法需预先定义路径，适合结构化项目

## 创建hello项目

```bash
cargo new hello # 创建项目
cargo run # 编译，运行
```

生成的项目结构：

```bash
hello/
├── Cargo.toml    # 项目配置
├── src/
│   └── main.rs   # 入口文件（或 lib.rs 库项目）
└── .gitignore    # 默认 Git 配置
```

默认的`main.rs`文件：

```rust
fn main() {
    println!("Hello, world!");
}
```

main 函数（也称为主函数）很特殊：它始终是每个可执行 Rust 程序中运行的第一个代码。第一行声明一个名为 main 的函数，不带参数也没有返回值。如果有参数，那么它们的名字会放到括号内，它们将放在括号 () 内。。

`cargo run`运行结果：

```
D:\rs_learn\hello>cargo run
   Compiling hello v0.1.0 (D:\rs_learn\hello)
    Finished `dev` profile [unoptimized + debuginfo] target(s) in 1.48s
     Running `target\debug\hello.exe`
Hello, world!
```

程序会先编译生成二进制文件`target\debug\hello.exe`，然后运行。你看到 "Hello, world!" 字符串。我们将这个字符串作为参数传递给 println!，接着 println! 将字符串打印到屏幕上。


---

> 作者: [hao](https://github.com/haochan1996)  
> URL: http://localhost:1313/rust/basic/96e9d38d-4156-45d0-942f-2c7c1c25e1aa/  

