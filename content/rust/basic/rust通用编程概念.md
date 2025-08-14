---
title: "Rust通用编程概念"
date: 2025-08-14T15:57:54+0800
slug: "e0e61349-11a8-4cef-bb97-8bd12d84e769"
draft: true
author: 
  name: hao
  link: https://github.com/haochan1996
  email: espholychan@outllook.com
  avatar: https://avatars.githubusercontent.com/u/190246046?v=4
description:
keywords:
license:
comment: false
weight: 0
tags:
  - Rust
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

## 变量

在 Rust 中，变量默认是不可变的（**immutable**）。这意味着一旦你给变量赋了值，就不能再改变它。这种设计是 Rust 保证内存安全和并发安全的核心思想之一。

### 不可变变量（Immutable Variables）

当你声明一个变量时，如果不做任何特殊处理，它就是不可变的。

```rust
fn main() {
    let x = 5;
    println!("The value of x is: {}", x);

    // 下面这行代码会导致编译错误，因为 x 是不可变的
    // x = 6;
}
```

编译器会报错，提示你不能对一个不可变的变量二次赋值。

### 可变变量（Mutable Variables）

如果你需要一个可以改变值的变量，需要在声明时使用 **`mut`** 关键字，使其变为可变（**mutable**）。

```rust
fn main() {
    let mut x = 5;
    println!("The value of x is: {}", x);

    // 现在可以修改 x 的值了
    x = 6;
    println!("The new value of x is: {}", x);
}
```

使用 `mut` 关键字的好处是，它能清晰地向代码的读者表明，这个变量的值在未来的某个时刻是可能被改变的。

-----

### 常量（Constants）

除了不可变变量，Rust 还有 **常量（constants）**。常量和不可变变量有些相似，但有几个关键区别：

1.  **关键字**: 声明常量需要使用 **`const`** 关键字，而不是 `let`。
2.  **可变性**: 常量总是不可变的，你不能给它们加上 `mut`。
3.  **命名规范**: 常量名通常使用全大写字母，并用下划线分隔单词，例如 `MAX_POINTS`。
4.  **表达式限制**: 常量只能被设置为常量表达式的结果，而不能是函数调用的结果或任何在运行时计算的值。

<!-- end list -->

```rust
const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;

fn main() {
    println!("Three hours in seconds is: {}", THREE_HOURS_IN_SECONDS);
}
```

### 变量遮蔽（Shadowing）

Rust 允许你用一个同名的新变量来“遮蔽”（**shadow**）旧的变量。当新的变量被声明后，旧的变量就无法再访问了。

这和可变变量（`mut`）的区别在于：

  * **`mut`**: 修改的是同一个变量的值。
  * **Shadowing**: 创建了一个新的、同名的变量，它完全覆盖了旧的变量。

遮蔽的好处是，你可以改变变量的类型，而这是 `mut` 无法做到的。

```rust
fn main() {
    let x = 5;

    // 遮蔽 x
    let x = x + 1;

    {
        // 在内部作用域中再次遮蔽 x
        let x = x * 2;
        println!("The value of x in the inner scope is: {}", x); // 输出 12
    }

    println!("The value of x is: {}", x); // 输出 6
}
```

在这里，内部作用域的 `x` 遮蔽了外部的 `x`，但当内部作用域结束时，外部的 `x` 依然是有效且未被改变的。


## 数据类型

在 Rust 中，每个变量都有一个特定的数据类型。Rust 是一种**静态类型**语言，这意味着它在编译时就必须知道所有变量的类型。编译器通常能够根据你赋值的方式推断出类型，但有时你可能需要明确地标注类型。

### 基本数据类型 (Primitive Data Types)

Rust 的基本数据类型可以分为两大类：**标量 (Scalar)** 和 **复合 (Compound)**。

#### 标量类型 (Scalar Types)

标量类型代表一个单一的值。

  * **整数 (Integers)**
    整数类型用于存储没有小数部分的数字。Rust 提供了多种有符号（`i`）和无符号（`u`）的整数类型，其名称表示其位数。

      * **有符号整数**：`i8`, `i16`, `i32`, `i64`, `i128`
      * **无符号整数**：`u8`, `u16`, `u32`, `u64`, `u128`
      * **`isize` 和 `usize`**：这两种类型的大小取决于你的计算机架构。在 64 位系统上是 64 位，在 32 位系统上是 32 位。它们主要用于索引集合或进行内存地址运算。

    默认情况下，Rust 编译器会推断为 `i32`，这是最常见的类型。

    ```rust
    let x = 42;    // 编译器推断为 i32
    let y: u64 = 10000; // 显式类型标注
    ```

  * **浮点数 (Floating-Point Numbers)**
    浮点数用于存储带小数点的数字。Rust 有两种主要的浮点数类型：

      * **`f32`**：单精度浮点数
      * **`f64`**：双精度浮点数 (默认)

    <!-- end list -->

    ```rust
    let pi = 3.14; // 编译器推断为 f64
    let z: f32 = 2.718;
    ```

  * **布尔值 (Booleans)**
    布尔类型只有两个可能的值：`true` 和 `false`。它的大小为 1 个字节。

    ```rust
    let is_rust_fun = true;
    let is_raining: bool = false;
    ```

  * **字符 (Characters)**
    `char` 类型是 Rust 最基础的字母类型。它以 Unicode 标量值表示，可以存储比 ASCII 更丰富的字符，包括中文、日文、表情符号等。`char` 类型的大小为 4 个字节。

    ```rust
    let heart_eye = '😻';
    let a_letter = 'a';
    ```

#### 复合类型 (Compound Types)

复合类型可以将多个值组合成一个类型。

  * **元组 (Tuples)**
    元组可以将多个不同类型的值打包成一个复合类型。元组的长度是固定的，一旦声明就不能改变。

    ```rust
    let tup: (i32, f64, char) = (500, 6.4, 'Z');
    let (x, y, z) = tup; // 解构元组
    println!("The value of y is: {}", y);
    
    // 也可以通过索引访问元组元素
    let five_hundred = tup.0;
    ```

  * **数组 (Arrays)**
    数组用于存储一组相同类型的值。数组的长度是固定的。与元组不同，数组中的每个元素必须是相同的类型。

    ```rust
    let a = [1, 2, 3, 4, 5];
    let months = ["Jan", "Feb", "Mar"];
    let first = a[0];
    
    // 你也可以用 `[类型; 长度]` 的形式声明
    let b: [i32; 5] = [1, 2, 3, 4, 5];
    
    // 或者创建一个包含相同值的数组
    let c = [3; 5]; // 等价于 [3, 3, 3, 3, 3]
    ```

    访问数组元素时，如果索引超出范围，Rust 会在运行时报错（panic），从而避免了许多 C/C++ 等语言中的缓冲区溢出问题。

### 类型转换 (Type Casting)

在 Rust 中，类型转换不是隐式进行的，你需要显式地进行转换。

```rust
let integer = 10;
let float = 5.5;

// 这行代码会报错，因为 i32 和 f64 不能直接相加
// let sum = integer + float;

let sum = integer as f64 + float; // 将 integer 转换为 f64
println!("Sum is: {}", sum);
```

通过这种方式，Rust 保证了类型安全，让你能清楚地知道代码中发生的类型转换。

## 函数

函数是 Rust 程序的基本构建块。你已经见过了最重要的函数 `main`，它是很多程序的入口点。你也见过了 `fn` 关键字，它用来声明新函数。

### 函数定义

Rust 代码中的函数定义以 `fn` 开始，后跟函数名和一对圆括号。大括号告诉编译器哪里是函数体的开始和结尾。

```rust
fn main() {
    println!("Hello, world!");
    
    another_function();
}

fn another_function() {
    println!("Another function.");
}
```

在 Rust 中，函数名使用下划线命名法（snake case），即所有字母都是小写并使用下划线分隔单词。

注意在 `main` 函数中调用了 `another_function`。我们能够调用它是因为它在源码中的某个地方被定义过了。Rust 不关心你在哪里定义函数，只要你在某个地方定义了它们。

### 函数参数

我们可以定义为拥有 **参数（parameters）** 的函数，参数是特殊变量，是函数签名的一部分。当函数拥有参数（形参）时，可以为这些参数提供具体的值（实参）。

```rust
fn main() {
    print_labeled_measurement(5, 'h');
}

fn print_labeled_measurement(value: i32, unit_label: char) {
    println!("The measurement is: {value}{unit_label}");
}
```

在函数签名中，**必须** 声明每个参数的类型。这是 Rust 设计中一个经过深思熟虑的决定：要求在函数定义中提供类型注解，意味着编译器不需要你在代码的其他地方注明类型来推断你的意图。

当定义多个参数时，使用逗号分隔参数声明，像这样：

```rust
fn greet(name: &str, age: i32) {
    println!("Hello, {}! You are {} years old.", name, age);
}

fn main() {
    greet("Alice", 30);
}
```

### 语句和表达式

函数体由一系列的语句和一个可选的结尾表达式构成。到目前为止，我们提到的函数还不包含结尾表达式，不过你已经见过作为语句一部分的表达式。因为 Rust 是一门基于表达式（expression-based）的语言，这是一个需要理解的重要区别。

- **语句（Statements）** 是执行一些操作但不返回值的指令。
- **表达式（Expressions）** 计算并产生一个值。

让我们看一些例子：

```rust
fn main() {
    let y = 6; // 这是一个语句
}
```

函数定义也是语句，上面整个例子本身就是一个语句。

语句不返回值。因此，不能把 `let` 语句赋值给另一个变量，比如下面的例子尝试做的，会产生一个错误：

```rust
fn main() {
    let x = (let y = 6); // 错误！
}
```

`let y = 6` 语句并不返回值，所以没有可以绑定到 `x` 上的值。

表达式计算出一个值，并且你将编写的大部分 Rust 代码是由表达式组成的。考虑一个数学运算，比如 `5 + 6`，这是一个表达式并计算出值 `11`。表达式可以是语句的一部分：在语句 `let y = 6;` 中，`6` 是一个表达式，它计算出的值是 `6`。

用大括号创建的一个新的块作用域也是一个表达式，例如：

```rust
fn main() {
    let y = {
        let x = 3;
        x + 1  // 注意这里没有分号
    };

    println!("The value of y is: {y}");
}
```

这个表达式：

```rust
{
    let x = 3;
    x + 1
}
```

是一个代码块，它的值是 `4`。这个值作为 `let` 语句的一部分被绑定到 `y` 上。注意 `x + 1` 这一行在结尾没有分号，与你见过的大部分代码行不同。**表达式的结尾没有分号。如果在表达式的结尾加上分号，它就变成了语句，而语句不会返回值。**

### 具有返回值的函数

函数可以向调用它的代码返回值。我们并不对返回值命名，但要在箭头（`->`）后声明它的类型。在 Rust 中，函数的返回值等同于函数体最后一个表达式的值。使用 `return` 关键字和指定值，可从函数中提前返回；但大部分函数隐式的返回最后的表达式。

```rust
fn five() -> i32 {
    5  // 这是一个表达式，没有分号
}

fn main() {
    let x = five();
    println!("The value of x is: {x}");
}
```

在 `five` 函数中没有函数调用、宏、甚至没有 `let` 语句——只有数字 `5`。这在 Rust 中是一个完全有效的函数。注意函数的返回类型也被指定好了，就是 `-> i32`。

让我们看看另一个例子：

```rust
fn plus_one(x: i32) -> i32 {
    x + 1  // 表达式，没有分号
}

fn main() {
    let x = plus_one(5);
    println!("The value of x is: {x}");
}
```

运行代码会打印出 `The value of x is: 6`。但如果在包含 `x + 1` 的行尾加上一个分号，把它从表达式变成语句，我们将看到一个错误：

```rust
fn plus_one(x: i32) -> i32 {
    x + 1;  // 语句，有分号 - 这会导致错误！
}
```

编译这段代码会产生一个错误，因为函数被定义为返回 `i32`，但语句并不会返回值。

### 函数的高级特性

#### 多返回值

虽然 Rust 函数只能返回一个值，但可以通过元组返回多个值：

```rust
fn calculate(x: i32, y: i32) -> (i32, i32, i32) {
    (x + y, x - y, x * y)
}

fn main() {
    let (sum, diff, product) = calculate(10, 5);
    println!("Sum: {}, Diff: {}, Product: {}", sum, diff, product);
}
```

#### 函数作为参数

Rust 支持高阶函数，可以将函数作为参数传递：

```rust
fn add(x: i32, y: i32) -> i32 {
    x + y
}

fn multiply(x: i32, y: i32) -> i32 {
    x * y
}

fn operate(x: i32, y: i32, operation: fn(i32, i32) -> i32) -> i32 {
    operation(x, y)
}

fn main() {
    let result1 = operate(10, 5, add);
    let result2 = operate(10, 5, multiply);
    
    println!("Add result: {}", result1);      // 15
    println!("Multiply result: {}", result2); // 50
}
```

#### 闭包（简介）

虽然这属于更高级的主题，但简单介绍一下闭包。闭包是可以捕获其环境的匿名函数：

```rust
fn main() {
    let x = 4;
    
    // 闭包
    let square = |num| num * num;
    let add_x = |num| num + x;  // 捕获了变量 x
    
    println!("Square of 5: {}", square(5));
    println!("5 + x: {}", add_x(5));
}
```

### 函数设计原则

编写好的函数时，应该遵循以下原则：

1. **单一职责**：每个函数应该只做一件事
2. **函数名清晰**：函数名应该清楚地表达其功能
3. **参数合理**：避免参数过多，一般不超过 3-4 个
4. **返回有意义的值**：如果函数有计算结果，应该返回它

```rust
// 好的函数设计示例
fn calculate_area(width: f64, height: f64) -> f64 {
    width * height
}

fn is_even(number: i32) -> bool {
    number % 2 == 0
}

fn format_greeting(name: &str) -> String {
    format!("Hello, {}!", name)
}

fn main() {
    let area = calculate_area(10.0, 5.0);
    let even = is_even(42);
    let greeting = format_greeting("Alice");
    
    println!("Area: {}", area);
    println!("Is 42 even? {}", even);
    println!("{}", greeting);
}
```

函数是 Rust 程序的核心组织单元。理解如何定义和使用函数，以及语句与表达式的区别，是掌握 Rust 编程的重要基础。

## 注释

注释是程序员在代码中写给自己和其他程序员看的说明文字。编译器会忽略注释，所以注释不会影响程序的功能。良好的注释习惯可以让代码更易于理解和维护。

### 行注释

在 Rust 中，最常用的注释方式是行注释，使用两个斜杠 `//` 开始：

```rust
fn main() {
    // 这是一个行注释
    println!("Hello, world!");
    
    let x = 5; // 这也是一个行注释，在代码行末尾
    
    // 你可以写多行注释
    // 每一行都以 // 开始
    // 这样就可以写很长的说明了
    let y = x * 2;
}
```

行注释从 `//` 开始一直到该行的末尾。你可以将行注释放在代码行的末尾，也可以单独占一行。

### 块注释

Rust 还支持块注释，使用 `/*` 开始，以 `*/` 结束：

```rust
fn main() {
    /*
     * 这是一个块注释
     * 可以跨越多行
     * 通常用于较长的说明
     */
    println!("Hello, world!");
    
    let x = /* 这个注释在代码中间 */ 42;
    
    /* 块注释也可以写在一行 */
    let y = x + 1;
}
```

块注释对于临时注释掉代码块特别有用，因为它们可以嵌套：

```rust
fn main() {
    /*
    let x = 5;
    /* 这是嵌套的注释 */
    let y = x + 1;
    println!("x = {}, y = {}", x, y);
    */
    
    println!("只有这行会执行");
}
```

### 文档注释

Rust 有特殊的文档注释语法，用于生成项目文档。这些注释使用三个斜杠 `///` 或 `//!`：

#### 外部文档注释 (`///`)

用于为紧跟在后面的项（函数、结构体、模块等）生成文档：

```rust
/// 计算两个数的加法
/// 
/// # 参数
/// 
/// * `a` - 第一个加数
/// * `b` - 第二个加数
/// 
/// # 返回值
/// 
/// 返回 `a` 和 `b` 的和
/// 
/// # 示例
/// 
/// ```
/// let result = add(2, 3);
/// assert_eq!(result, 5);
/// ```
fn add(a: i32, b: i32) -> i32 {
    a + b
}

/// 表示一个矩形的结构体
/// 
/// 包含宽度和高度两个字段
struct Rectangle {
    /// 矩形的宽度
    width: f64,
    /// 矩形的高度  
    height: f64,
}

impl Rectangle {
    /// 创建一个新的矩形
    /// 
    /// # 参数
    /// 
    /// * `width` - 矩形的宽度，必须大于 0
    /// * `height` - 矩形的高度，必须大于 0
    /// 
    /// # 示例
    /// 
    /// ```
    /// let rect = Rectangle::new(10.0, 20.0);
    /// ```
    fn new(width: f64, height: f64) -> Rectangle {
        Rectangle { width, height }
    }
    
    /// 计算矩形的面积
    /// 
    /// # 返回值
    /// 
    /// 返回矩形的面积（宽度 × 高度）
    fn area(&self) -> f64 {
        self.width * self.height
    }
}
```

#### 内部文档注释 (`//!`)

用于为包含它的项（通常是模块或 crate）生成文档：

```rust
//! # 我的数学库
//! 
//! 这个库提供了一些基本的数学运算功能。
//! 
//! ## 特性
//! 
//! - 基本算术运算
//! - 几何图形计算
//! - 统计函数
//! 
//! ## 使用示例
//! 
//! ```
//! use my_math_lib::add;
//! 
//! let result = add(2, 3);
//! assert_eq!(result, 5);
//! ```

fn main() {
    // 函数体...
}
```

### 注释的最佳实践

#### 1. 解释"为什么"而不是"什么"

```rust
// 好的注释 - 解释为什么
fn calculate_discount(price: f64, customer_type: &str) -> f64 {
    // VIP 客户享受 20% 折扣以提高客户忠诚度
    if customer_type == "VIP" {
        price * 0.8
    } else {
        // 普通客户享受 5% 折扣以促进销售
        price * 0.95
    }
}

// 不好的注释 - 只是重复代码
fn bad_example(x: i32) -> i32 {
    // 将 x 乘以 2
    x * 2  // 这个注释没有添加任何有用信息
}
```

#### 2. 保持注释简洁明了

```rust
// 好的注释
fn process_user_input(input: &str) -> Result<i32, ParseIntError> {
    // 去除空白字符并转换为整数
    input.trim().parse()
}

// 不好的注释
fn bad_process(input: &str) -> Result<i32, ParseIntError> {
    // 这个函数接收一个字符串切片作为输入参数，然后我们调用 trim() 方法
    // 来移除字符串开头和结尾的空白字符，接着调用 parse() 方法尝试将
    // 清理后的字符串解析为一个 32 位有符号整数...
    input.trim().parse()
}
```

#### 3. 使用 TODO 和 FIXME 标记

```rust
fn incomplete_function() {
    // TODO: 实现错误处理
    // FIXME: 这里可能存在内存泄漏
    // HACK: 临时解决方案，需要重构
    
    println!("基本功能");
}
```

#### 4. 注释复杂的算法和业务逻辑

```rust
fn fibonacci(n: u32) -> u32 {
    // 使用迭代方式计算斐波那契数列，避免递归的指数时间复杂度
    if n <= 1 {
        return n;
    }
    
    let mut prev = 0;
    let mut curr = 1;
    
    // 从第三个数开始，每个数都是前两个数的和
    for _ in 2..=n {
        let next = prev + curr;
        prev = curr;
        curr = next;
    }
    
    curr
}
```

### 生成文档

当你使用文档注释时，可以使用 `cargo doc` 命令生成 HTML 格式的文档：

```bash
# 生成文档
cargo doc

# 生成文档并在浏览器中打开
cargo doc --open

# 包含私有项的文档
cargo doc --document-private-items
```

### 注释中的 Markdown

文档注释支持 Markdown 语法，让你可以创建格式丰富的文档：

```rust
/// # 数学工具函数
/// 
/// 这个模块包含各种数学计算函数。
/// 
/// ## 支持的操作
/// 
/// 1. **基本运算**
///    - 加法
///    - 减法  
///    - 乘法
///    - 除法
/// 
/// 2. **高级运算**
///    - 幂运算
///    - 开方
/// 
/// ## 代码示例
/// 
/// ```rust
/// use math_utils::power;
/// 
/// let result = power(2, 3);
/// assert_eq!(result, 8);
/// ```
/// 
/// > **注意**: 所有函数都会检查输入参数的有效性
/// 
/// 更多信息请参考 [官方文档](https://doc.rust-lang.org/)
fn power(base: i32, exponent: u32) -> i32 {
    base.pow(exponent)
}
```

良好的注释是高质量代码的重要组成部分。它们不仅帮助其他人理解你的代码，也帮助未来的你更快地回忆起代码的逻辑和意图。记住，代码告诉计算机做什么，而注释告诉人类为什么要这样做。