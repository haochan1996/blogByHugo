---
type: posts
title: "Rust枚举和模式匹配"
date: 2025-08-21T15:54:21+0800
slug: "48f5da79-a32a-4845-bc90-bcbb574ce0e9"
draft: false
author: 
  name: hao
  link: https://github.com/haochan1996
  email: espholychan@outllook.com
  avatar: https://avatars.githubusercontent.com/u/190246046?v=4
description: 深入学习Rust枚举和模式匹配，掌握函数式编程的核心特性
keywords:
license:
comment: false
weight: 4
tags:
  - Rust
  - Enum
  - Pattern Matching
categories:
  - Rust基础入门
hiddenFromHomePage: false
hiddenFromSearch: false
hiddenFromRelated: false
hiddenFromFeed: false
summary:
resources:
  - name: featured-image
    src: featured-image.jpg
  - name: featured-image-preview
    src: featured-image-preview.jpg
toc: true
math: false
lightgallery: false
password:
message:
repost:
  enable: true
  url:
---

## 什么是枚举

**枚举**（enumerations，也被称为 enums）允许你通过列举可能的**成员**（variants）来定义一个类型。枚举给了你一个灵活且强大的方式来表示数据可能是几种不同类型中的一种。

在很多语言中，枚举只能是简单的值列表，但 Rust 的枚举更加强大，可以存储数据，并且每个变体可以有不同类型和数量的关联数据。

## 定义枚举

让我们看一个需要枚举的场景：处理 IP 地址。目前使用的 IP 地址有两个主要标准：版本四（IPv4）和版本六（IPv6）。这些是我们的程序可能会遇到的所有可能的 IP 地址类型，所以可以**枚举**出所有可能的值，这也正是此枚举名字的由来。

```rust
enum IpAddrKind {
    V4,
    V6,
}
```

现在 `IpAddrKind` 就是一个可以在代码中使用的自定义数据类型了。

### 枚举值

可以像这样创建 `IpAddrKind` 两个不同成员的实例：

```rust
enum IpAddrKind {
    V4,
    V6,
}

fn main() {
    let four = IpAddrKind::V4;
    let six = IpAddrKind::V6;

    route(IpAddrKind::V4);
    route(IpAddrKind::V6);
}

fn route(ip_kind: IpAddrKind) {
    // 处理 IP 地址类型
}
```

注意枚举的成员位于其标识符的命名空间中，并使用两个冒号分开。这么设计的益处是现在 `IpAddrKind::V4` 和 `IpAddrKind::V6` 都是 `IpAddrKind` 类型的。

### 将数据附加到枚举变体

我们可以将数据直接放进每一个枚举变体中，这样就不需要额外的结构体了：

```rust
enum IpAddr {
    V4(String),
    V6(String),
}

fn main() {
    let home = IpAddr::V4(String::from("127.0.0.1"));
    let loopback = IpAddr::V6(String::from("::1"));
}
```

我们直接将数据附加到枚举的每个变体上，这样就不需要额外的结构体了。这里我们用 `String` 来存储 IP 地址。

枚举的另一个优势：每个变体可以处理不同类型和数量的数据。IPv4 版本的 IP 地址总是含有四个值在 0 和 255 之间的数字组件。如果我们想要将 `V4` 地址存储为四个 `u8` 值而 `V6` 地址仍然表现为一个 `String`，这就不能用结构体了：

```rust
enum IpAddr {
    V4(u8, u8, u8, u8),
    V6(String),
}

fn main() {
    let home = IpAddr::V4(127, 0, 0, 1);
    let loopback = IpAddr::V6(String::from("::1"));
}
```

### 更复杂的枚举示例

来看一个更复杂的例子，展示枚举变体的多样性：

```rust
enum Message {
    Quit,                       // 没有关联任何数据
    Move { x: i32, y: i32 },    // 类似结构体包含命名字段
    Write(String),              // 包含单独一个 String
    ChangeColor(i32, i32, i32), // 包含三个 i32
}
```

这个枚举有四个含有不同类型的变体：

- `Quit` 没有关联任何数据
- `Move` 类似结构体包含命名字段
- `Write` 包含单独一个 `String`
- `ChangeColor` 包含三个 `i32`

### 为枚举定义方法

就像我们能够在结构体上定义方法那样，也可以在枚举上定义方法：

```rust
impl Message {
    fn call(&self) {
        // 在这里定义方法体
        match self {
            Message::Quit => println!("Quit message"),
            Message::Move { x, y } => println!("Move to ({}, {})", x, y),
            Message::Write(text) => println!("Write: {}", text),
            Message::ChangeColor(r, g, b) => println!("Change color to ({}, {}, {})", r, g, b),
        }
    }
}

fn main() {
    let m = Message::Write(String::from("hello"));
    m.call();
}
```

## Option 枚举

`Option` 是标准库定义的另一个非常有用的枚举。`Option` 类型应用广泛因为它编码了一个非常普遍的场景，即一个值要么有值要么没值。

在很多语言中，空值（null）是一个特殊的值，表示"没有值"。但是在 Rust 中没有空值功能。不过，Rust 确实拥有一个可以编码存在或不存在概念的枚举。这个枚举是 `Option<T>`：

```rust
enum Option<T> {
    None,
    Some(T),
}
```

`Option<T>` 枚举是如此有用以至于它甚至被包含在了 prelude 之中，你不需要将其显式引入作用域。另外，它的成员也是如此，可以不需要 `Option::` 前缀来直接使用 `Some` 和 `None`。

```rust
fn main() {
    let some_number = Some(5);
    let some_char = Some('e');

    let absent_number: Option<i32> = None;
}
```

如果使用 `None` 而不是 `Some`，需要告诉 Rust `Option<T>` 是什么类型的，因为编译器只通过 `None` 值无法推断出 `Some` 变体保存的值的类型。

### Option 的优势

为什么有 `Option<T>` 比有空值要好呢？简而言之，因为 `Option<T>` 和 `T`（这里 `T` 可以是任何类型）是不同的类型，编译器不允许像一个肯定有效的值那样使用 `Option<T>`。

```rust
fn main() {
    let x: i8 = 5;
    let y: Option<i8> = Some(5);

    // let sum = x + y; // 这行会编译错误！
}
```

强制显式处理可能为空的情况。为了使用 `Option<T>` 值，需要将其转换为 `T`。通常这能帮助我们捕获到空值最常见的问题之一：假设某值不为空但实际上为空。

## match 控制流结构

Rust 有一个叫做 `match` 的极为强大的控制流运算符，它允许我们将一个值与一系列的模式相比较，并根据相匹配的模式执行相应代码。模式可由字面值、变量、通配符和许多其他内容构成。

### 匹配枚举

让我们写一个函数，它可以接受一个未知的硬币，并以一种类似验钞机的方式，确定它是何种硬币并返回它的美分值：

```rust
#[derive(Debug)]
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}

fn main() {
    let coin = Coin::Penny;
    println!("Value: {}", value_in_cents(coin));
}
```

### 绑定值的模式

匹配分支的另一个有用的功能是可以绑定匹配的模式的部分值。这也就是如何从枚举变体中提取值。

```rust
#[derive(Debug)]
enum UsState {
    Alabama,
    Alaska,
    // --snip--
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u8 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        }
    }
}

fn main() {
    value_in_cents(Coin::Quarter(UsState::Alaska));
}
```

### 匹配 Option<T>

让我们使用 `match` 来处理 `Option<T>`！我们想要编写一个函数，它获取一个 `Option<i32>` 并且如果其中含有一个值，将其加一。如果其中没有值，函数应该返回 `None` 值，而不尝试执行任何操作。

```rust
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        None => None,
        Some(i) => Some(i + 1),
    }
}

fn main() {
    let five = Some(5);
    let six = plus_one(five);
    let none = plus_one(None);
    
    println!("{:?}", six);  // Some(6)
    println!("{:?}", none); // None
}
```

### 匹配是穷尽的

`match` 还有另一个方面需要讨论：这些分支必须覆盖了所有的可能性。Rust 中的匹配是**穷尽的**（exhaustive）：必须穷举到最后的可能性来使代码有效。

### 通配模式和 _ 占位符

使用枚举，我们也可以对一些特定的值采取特殊操作，但是对其他的所有值采取默认操作：

```rust
fn main() {
    let dice_roll = 9;
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        other => move_player(other), // 捕获其他所有值
    }
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
fn move_player(num_spaces: u8) {}
```

如果我们不想使用通配模式获取的值，可以使用 `_`，这是一个特殊的模式，可以匹配任意值而不绑定到该值：

```rust
fn main() {
    let dice_roll = 9;
    match dice_roll {
        3 => add_fancy_hat(),
        7 => remove_fancy_hat(),
        _ => (), // 匹配其他所有值，但不做任何事
    }
}

fn add_fancy_hat() {}
fn remove_fancy_hat() {}
```

## if let 简洁控制流

`if let` 语法让我们以一种不那么冗长的方式结合 `if` 和 `let`，来处理只匹配一个模式的值而忽略其他模式的情况。

考虑这样的代码，它匹配一个 `Option<u8>` 值并只希望当值为 3 时执行代码：

```rust
fn main() {
    let config_max = Some(3u8);
    match config_max {
        Some(max) => println!("The maximum is configured to be {}", max),
        _ => (),
    }
}
```

我们可以使用 `if let` 这种更短的方式编写：

```rust
fn main() {
    let config_max = Some(3u8);
    if let Some(max) = config_max {
        println!("The maximum is configured to be {}", max);
    }
}
```

我们也可以在 `if let` 中包含一个 `else`：

```rust
fn main() {
    let coin = Coin::Penny;
    let mut count = 0;
    
    if let Coin::Quarter(state) = coin {
        println!("State quarter from {:?}!", state);
    } else {
        count += 1;
    }
}
```

## 模式的语法

模式是 Rust 中特殊的语法，它用来匹配类型中的结构。结合使用模式和 `match` 表达式以及其他结构可以提供更多对程序控制流的支持。

### 字面值和变量

你可以直接匹配字面值模式，也可以使用命名变量：

```rust
fn main() {
    let x = 1;

    match x {
        1 => println!("one"),       // 字面值模式
        2 => println!("two"),
        3 => println!("three"),
        _ => println!("anything"),  // 通配符模式
    }
    
    let y = Some(5);
    match y {
        Some(50) => println!("Got 50"),
        Some(n) => println!("Matched, n = {}", n), // 变量模式
        _ => (),
    }
}
```

### 多个模式和范围

在 `match` 表达式中，可以使用 `|` 语法匹配多个模式，使用 `..=` 匹配范围：

```rust
fn main() {
    let x = 1;

    match x {
        1 | 2 => println!("one or two"),    // 多个模式
        3 => println!("three"),
        _ => println!("anything"),
    }
    
    let y = 5;
    match y {
        1..=5 => println!("one through five"),  // 范围模式
        _ => println!("something else"),
    }
    
    let z = 'c';
    match z {
        'a'..='j' => println!("early ASCII letter"),
        'k'..='z' => println!("late ASCII letter"),
        _ => println!("something else"),
    }
}
```

### 解构以分解值

我们也可以使用模式来解构结构体、枚举和元组：

#### 解构结构体

```rust
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let p = Point { x: 0, y: 7 };

    let Point { x: a, y: b } = p;
    assert_eq!(0, a);
    assert_eq!(7, b);
    
    // 简写形式
    let Point { x, y } = p;
    assert_eq!(0, x);
    assert_eq!(7, y);
    
    match p {
        Point { x, y: 0 } => println!("On the x axis at {x}"),
        Point { x: 0, y } => println!("On the y axis at {y}"),
        Point { x, y } => println!("On neither axis: ({x}, {y})"),
    }
}
```

#### 解构枚举

```rust
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn main() {
    let msg = Message::ChangeColor(0, 160, 255);

    match msg {
        Message::Quit => {
            println!("The Quit variant has no data to destructure.");
        }
        Message::Move { x, y } => {
            println!("Move in the x direction {x} and in the y direction {y}");
        }
        Message::Write(text) => {
            println!("Text message: {text}");
        }
        Message::ChangeColor(r, g, b) => {
            println!("Change the color to red {r}, green {g}, and blue {b}")
        }
    }
}
```

### 忽略模式中的值

#### 使用 _ 忽略整个值

```rust
fn foo(_: i32, y: i32) {
    println!("This code only uses the y parameter: {}", y);
}

fn main() {
    foo(3, 4);
}
```

#### 用 .. 忽略剩余值

```rust
struct Point {
    x: i32,
    y: i32,
    z: i32,
}

fn main() {
    let origin = Point { x: 0, y: 0, z: 0 };

    match origin {
        Point { x, .. } => println!("x is {}", x),
    }
    
    let numbers = (2, 4, 8, 16, 32);
    match numbers {
        (first, .., last) => {
            println!("Some numbers: {first}, {last}");
        }
    }
}
```

### 匹配守卫

**匹配守卫**（match guard）是一个指定于 `match` 分支模式之后的额外 `if` 条件：

```rust
fn main() {
    let num = Some(4);

    match num {
        Some(x) if x % 2 == 0 => println!("The number {} is even", x),
        Some(x) => println!("The number {} is odd", x),
        None => (),
    }
}
```

### @ 绑定

`@` 运算符允许我们在创建一个存放值的变量的同时测试其值是否匹配模式：

```rust
enum Message {
    Hello { id: i32 },
}

fn main() {
    let msg = Message::Hello { id: 5 };

    match msg {
        Message::Hello {
            id: id_variable @ 3..=7,
        } => println!("Found an id in range: {}", id_variable),
        Message::Hello { id: 10..=12 } => {
            println!("Found an id in another range")
        }
        Message::Hello { id } => println!("Found some other id: {}", id),
    }
}
```

## 实践示例

### 1. 状态机实现

```rust
#[derive(Debug)]
enum State {
    Waiting,
    Processing { progress: u8 },
    Completed { result: String },
    Error { message: String },
}

struct Task {
    id: u32,
    state: State,
}

impl Task {
    fn new(id: u32) -> Self {
        Task {
            id,
            state: State::Waiting,
        }
    }
    
    fn start_processing(&mut self) {
        match self.state {
            State::Waiting => {
                self.state = State::Processing { progress: 0 };
                println!("Task {} started processing", self.id);
            }
            _ => println!("Task {} is not in waiting state", self.id),
        }
    }
    
    fn update_progress(&mut self, progress: u8) {
        match &mut self.state {
            State::Processing { progress: ref mut p } => {
                *p = progress;
                println!("Task {} progress: {}%", self.id, progress);
                if progress >= 100 {
                    self.state = State::Completed {
                        result: format!("Task {} completed successfully", self.id),
                    };
                }
            }
            _ => println!("Task {} is not currently processing", self.id),
        }
    }
    
    fn get_status(&self) -> String {
        match &self.state {
            State::Waiting => "Waiting to start".to_string(),
            State::Processing { progress } => format!("Processing: {}%", progress),
            State::Completed { result } => result.clone(),
            State::Error { message } => format!("Error: {}", message),
        }
    }
}

fn main() {
    let mut task = Task::new(1);
    println!("Status: {}", task.get_status());
    
    task.start_processing();
    task.update_progress(50);
    task.update_progress(100);
    
    println!("Final status: {}", task.get_status());
}
```

### 2. HTTP 请求处理器

```rust
#[derive(Debug)]
enum HttpMethod {
    Get,
    Post,
    Put,
    Delete,
}

#[derive(Debug)]
enum HttpStatus {
    Ok,
    NotFound,
    BadRequest,
    InternalServerError,
}

#[derive(Debug)]
struct HttpRequest {
    method: HttpMethod,
    path: String,
    body: Option<String>,
}

#[derive(Debug)]
struct HttpResponse {
    status: HttpStatus,
    body: String,
}

fn handle_request(request: HttpRequest) -> HttpResponse {
    match (&request.method, request.path.as_str()) {
        (HttpMethod::Get, "/") => HttpResponse {
            status: HttpStatus::Ok,
            body: "Welcome to the home page!".to_string(),
        },
        (HttpMethod::Get, "/users") => HttpResponse {
            status: HttpStatus::Ok,
            body: "List of users".to_string(),
        },
        (HttpMethod::Post, "/users") => {
            match &request.body {
                Some(body) if !body.is_empty() => HttpResponse {
                    status: HttpStatus::Ok,
                    body: format!("Created user with data: {}", body),
                },
                _ => HttpResponse {
                    status: HttpStatus::BadRequest,
                    body: "Missing user data".to_string(),
                },
            }
        },
        (HttpMethod::Get, path) if path.starts_with("/users/") => {
            let user_id = &path[7..]; // 提取用户ID
            if user_id.chars().all(|c| c.is_numeric()) {
                HttpResponse {
                    status: HttpStatus::Ok,
                    body: format!("User details for ID: {}", user_id),
                }
            } else {
                HttpResponse {
                    status: HttpStatus::BadRequest,
                    body: "Invalid user ID".to_string(),
                }
            }
        },
        _ => HttpResponse {
            status: HttpStatus::NotFound,
            body: "Page not found".to_string(),
        },
    }
}

fn main() {
    let requests = vec![
        HttpRequest {
            method: HttpMethod::Get,
            path: "/".to_string(),
            body: None,
        },
        HttpRequest {
            method: HttpMethod::Post,
            path: "/users".to_string(),
            body: Some("John Doe".to_string()),
        },
        HttpRequest {
            method: HttpMethod::Get,
            path: "/users/123".to_string(),
            body: None,
        },
    ];
    
    for request in requests {
        let response = handle_request(request);
        println!("Response: {:?}", response);
    }
}
```

## 最佳实践

### 1. 优先使用枚举而不是布尔值

```rust
// 不好的做法
fn process_user(is_admin: bool, is_active: bool) {
    // 难以理解参数含义
}

// 好的做法
#[derive(Debug)]
enum UserRole {
    Admin,
    User,
    Guest,
}

#[derive(Debug)]
enum UserStatus {
    Active,
    Inactive,
    Suspended,
}

fn process_user(role: UserRole, status: UserStatus) {
    match (role, status) {
        (UserRole::Admin, UserStatus::Active) => {
            // 管理员且活跃
        }
        (_, UserStatus::Suspended) => {
            // 任何被暂停的用户
        }
        _ => {
            // 其他情况
        }
    }
}
```

### 2. 使用 Result 进行错误处理

```rust
#[derive(Debug)]
enum ParseError {
    InvalidFormat,
    NumberTooLarge,
    EmptyInput,
}

fn parse_number(input: &str) -> Result<i32, ParseError> {
    if input.is_empty() {
        return Err(ParseError::EmptyInput);
    }
    
    match input.parse::<i32>() {
        Ok(num) => {
            if num > 1000 {
                Err(ParseError::NumberTooLarge)
            } else {
                Ok(num)
            }
        }
        Err(_) => Err(ParseError::InvalidFormat),
    }
}

fn main() {
    let inputs = vec!["123", "abc", "", "2000"];
    
    for input in inputs {
        match parse_number(input) {
            Ok(num) => println!("Parsed: {}", num),
            Err(ParseError::EmptyInput) => println!("Error: Input is empty"),
            Err(ParseError::InvalidFormat) => println!("Error: Invalid format"),
            Err(ParseError::NumberTooLarge) => println!("Error: Number too large"),
        }
    }
}
```

## 总结

枚举和模式匹配是 Rust 中极其强大的特性：

1. **类型安全**：枚举提供了类型安全的方式来表示可能的值
2. **数据关联**：枚举变体可以携带不同类型和数量的数据
3. **穷尽匹配**：`match` 确保处理所有可能的情况
4. **模式匹配**：强大的模式语法支持复杂的数据解构
5. **错误处理**：`Option` 和 `Result` 枚举提供了安全的错误处理机制
6. **控制流**：`if let` 和 `while let` 提供简洁的控制流语法

掌握枚举和模式匹配是学习 Rust 函数式编程风格的关键，它们让代码更加安全、表达力更强，并帮助编译器发现潜在的错误。这些特性使得 Rust 能够在编译时捕获许多其他语言在运行时才会发现的错误。

