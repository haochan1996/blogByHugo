---
type: posts
title: "Rust结构体"
date: 2025-08-19T22:41:13+0800
slug: "800ae850-1128-4994-b419-72cb33047087"
draft: false
author: 
  name: hao
  link: https://github.com/haochan1996
  email: espholychan@outllook.com
  avatar: https://avatars.githubusercontent.com/u/190246046?v=4
description: 深入学习Rust结构体，包括定义、实例化、方法、关联函数等核心概念
keywords:
license:
comment: false
weight: 3
tags:
  - Rust
categories:
  - "Rust基础入门"
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

## 什么是结构体

**结构体**（struct）是一种自定义数据类型，允许你将多个相关联的值组合在一起，形成一个有意义的组合。如果你熟悉面向对象语言，结构体就像对象的数据属性。

结构体和元组类似，都可以包含多个不同类型的值。但与元组不同的是，结构体需要为每个数据片段命名，这样就更清楚各个值的意义，也不需要依赖数据的顺序来访问结构体的值。

## 定义和实例化结构体

### 定义结构体

使用 `struct` 关键字来定义结构体，并为整个结构体提供一个名字。结构体的名字需要描述它所组合的数据的意义。接着，在大括号中，定义每一部分数据的名字和类型，我们称为**字段**（field）：

```rust
struct User {
    active: bool,
    username: String,
    email: String,
    sign_in_count: u64,
}
```

### 实例化结构体

一旦定义了结构体，我们可以通过为每个字段指定具体值来创建这个结构体的**实例**：

```rust
fn main() {
    let user1 = User {
        active: true,
        username: String::from("someusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };
    
    println!("User: {}", user1.username);
}
```

为了从结构体中获取某个值，我们使用点号。如果结构体的实例是可变的，我们也可以使用点号并为对应的字段赋值。注意整个实例必须是可变的；Rust 并不允许只将某个字段标记为可变。

```rust
fn main() {
    let mut user1 = User {
        active: true,
        username: String::from("someusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };

    user1.email = String::from("anotheremail@example.com");
}
```

### 使用字段初始化简写语法

当变量名与字段名相同时，可以使用**字段初始化简写语法**（field init shorthand）：

```rust
fn build_user(email: String, username: String) -> User {
    User {
        active: true,
        username, // 简写，等同于 username: username,
        email,    // 简写，等同于 email: email,
        sign_in_count: 1,
    }
}
```

### 使用结构体更新语法从其他实例创建实例

使用**结构体更新语法**（struct update syntax）可以基于其他实例创建新实例：

```rust
fn main() {
    let user1 = User {
        active: true,
        username: String::from("someusername123"),
        email: String::from("someone@example.com"),
        sign_in_count: 1,
    };

    let user2 = User {
        active: user1.active,
        username: user1.username,
        email: String::from("another@example.com"),
        sign_in_count: user1.sign_in_count,
    };

    // 使用结构体更新语法，更简洁的写法
    let user3 = User {
        email: String::from("another@example.com"),
        ..user1 // 必须放在最后，表示剩余字段应从 user1 获取值
    };
    
    // 注意：user1 在这里不再可用，因为 username 被移动了
    // println!("{}", user1.username); // 这行会编译错误
}
```

注意，结构体更新语法就像带有 `=` 的赋值，因为它移动了数据。在上面的例子中，我们在创建 `user3` 后不能再使用 `user1`，因为 `user1` 的 `username` 字段中的 `String` 被移到 `user3` 中。如果我们给 `user3` 的 `email` 和 `username` 都赋予新的 `String` 值，而只使用 `user1` 的 `active` 和 `sign_in_count` 值，那么 `user1` 在创建 `user3` 后仍然有效。

## 元组结构体

Rust 还支持类似元组的结构体，称为**元组结构体**（tuple structs）。元组结构体有着结构体名称提供的含义，但没有具体的字段名，只有字段的类型：

```rust
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
    let black = Color(0, 0, 0);
    let origin = Point(0, 0, 0);
    
    // 访问元组结构体的值
    println!("Black color: ({}, {}, {})", black.0, black.1, black.2);
}
```

注意 `black` 和 `origin` 值的类型不同，因为它们是不同的元组结构体的实例。你定义的每一个结构体有其自己的类型，即使结构体中的字段有着相同的类型。

## 类单元结构体

你也可以定义一个没有任何字段的结构体！它们被称为**类单元结构体**（unit-like structs）因为它们类似于 `()`，即 unit 类型：

```rust
struct AlwaysEqual;

fn main() {
    let subject = AlwaysEqual;
}
```

类单元结构体常常在你想要在某个类型上实现 trait 但不需要在类型中存储数据的时候发挥作用。

## 结构体数据的所有权

在上面的 `User` 结构体的定义中，我们使用了自身拥有所有权的 `String` 类型而不是 `&str` 字符串 slice 类型。这是一个有意而为之的选择，因为我们想要这个结构体拥有它所有的数据，为此只要整个结构体是有效的话其数据也是有效的。

可以使结构体存储被其他对象拥有的数据的引用，不过这么做的话需要用上**生命周期**（lifetimes），这是一个后面会讨论的 Rust 功能。生命周期确保结构体引用的数据有效性跟结构体本身保持一致。如果你尝试在结构体中存储一个引用而不指定生命周期将是无效的：

```rust
// 这个不能编译！
struct User {
    active: bool,
    username: &str, // 缺少生命周期说明符
    email: &str,    // 缺少生命周期说明符
    sign_in_count: u64,
}

fn main() {
    let user1 = User {
        active: true,
        username: "someusername123",
        email: "someone@example.com",
        sign_in_count: 1,
    };
}
```

## 方法语法

**方法**（method）与函数类似：它们使用 `fn` 关键字和名称声明，可以拥有参数和返回值，同时包含在某处调用该方法时会执行的代码。不过方法与函数是不同的，因为它们在结构体的上下文中被定义，并且它们第一个参数总是 `self`，它代表调用该方法的结构体实例。

### 定义方法

让我们把前面实现的求面积的函数改写成 `Rectangle` 结构体上的 `area` 方法：

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!(
        "The area of the rectangle is {} square pixels.",
        rect1.area()
    );
}
```

为了使函数定义于 `Rectangle` 的上下文中，我们开始了一个 `impl`（implementation）块。这个 `impl` 块中的所有内容都将与 `Rectangle` 类型相关联。

在 `area` 的签名中，我们使用 `&self` 来替代 `rectangle: &Rectangle`。`&self` 实际上是 `self: &Self` 的缩写。在一个 `impl` 块中，`Self` 类型是 `impl` 块的类型的别名。方法的第一个参数必须有一个名为 `self` 的 `Self` 类型的参数，所以 Rust 让你在第一个参数位置上只用 `self` 这个名字来缩写。

我们仍然需要在 `self` 前面使用 `&` 来表示这个方法借用了 `Self` 实例，就像我们在 `rectangle: &Rectangle` 中做的那样。方法可以选择获得 `self` 的所有权，或者像我们这里一样不可变地借用 `self`，或者可变地借用 `self`，就跟其他参数一样。

### 带有更多参数的方法

让我们通过实现 `Rectangle` 结构体上的另一方法来练习使用方法。这次，我们让一个 `Rectangle` 的实例获取另一个 `Rectangle` 实例，如果 `self` 能完全包含第二个长方形则返回 `true`；否则返回 `false`：

```rust
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }

    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };
    let rect2 = Rectangle {
        width: 10,
        height: 40,
    };
    let rect3 = Rectangle {
        width: 60,
        height: 45,
    };

    println!("Can rect1 hold rect2? {}", rect1.can_hold(&rect2));
    println!("Can rect1 hold rect3? {}", rect1.can_hold(&rect3));
}
```

### 关联函数

所有在 `impl` 块中定义的函数被称为**关联函数**（associated functions），因为它们与 `impl` 后面命名的类型相关。我们可以定义不以 `self` 为第一参数的关联函数（因此不是方法），因为它们并不作用于一个结构体的实例。我们已经使用过这样的函数了：在 `String` 类型上定义的 `String::from` 函数。

不是方法的关联函数经常被用作返回一个结构体新实例的构造函数。这些函数的名称通常为 `new`，但 `new` 并不是一个关键字。例如，我们可以提供一个叫做 `square` 的关联函数，它接受一个维度参数并且同时作为宽和高，这样可以更轻松地创建一个正方形 `Rectangle` 而不必指定两次同样的值：

```rust
impl Rectangle {
    fn square(size: u32) -> Self {
        Self {
            width: size,
            height: size,
        }
    }
}

fn main() {
    let sq = Rectangle::square(3);
}
```

关键字 `Self` 在函数的返回类型中代指在 `impl` 关键字后出现的类型，在这里是 `Rectangle`。

使用结构体名和 `::` 语法来调用这个关联函数：比如 `let sq = Rectangle::square(3);`。这个方法位于结构体的命名空间中，`::` 语法用于关联函数和模块创建的命名空间。

### 多个 impl 块

每个结构体都允许拥有多个 `impl` 块：

```rust
impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}

impl Rectangle {
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}
```

这里没有理由将这些方法分散在多个 `impl` 块中，不过这是有效的语法。

## 自动推导 trait

Rust 提供了许多可以通过 `derive` 注解来使用的 trait，它们可以为我们的自定义类型增加实用的行为。

### Debug trait

为了能够打印 `Rectangle` 实例来调试，我们可以在结构体定义之前加上外部属性 `#[derive(Debug)]`：

```rust
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

fn main() {
    let rect1 = Rectangle {
        width: 30,
        height: 50,
    };

    println!("rect1 is {:?}", rect1);
    println!("rect1 is {:#?}", rect1); // 更好看的格式
}
```

### 其他常用的可推导 trait

```rust
#[derive(Debug, Clone, PartialEq)]
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let p1 = Point { x: 1, y: 2 };
    let p2 = p1.clone(); // Clone trait
    
    println!("p1: {:?}", p1); // Debug trait
    println!("p1 == p2: {}", p1 == p2); // PartialEq trait
}
```

常用的可推导 trait：
- `Debug`：允许使用 `{:?}` 格式化输出
- `Clone`：允许显式复制值
- `Copy`：允许类型进行简单的位复制
- `PartialEq` 和 `Eq`：允许比较相等性
- `PartialOrd` 和 `Ord`：允许比较大小
- `Hash`：允许类型作为 HashMap 的键

## 结构体的高级特性

### 1. 结构体与模式匹配

结构体可以在模式匹配中使用：

```rust
struct Point {
    x: i32,
    y: i32,
}

fn main() {
    let p = Point { x: 0, y: 7 };

    match p {
        Point { x, y: 0 } => println!("On the x axis at {}", x),
        Point { x: 0, y } => println!("On the y axis at {}", y),
        Point { x, y } => println!("On neither axis: ({}, {})", x, y),
    }
}
```

### 2. 结构体解构

可以使用 `let` 语句解构结构体：

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
}
```

### 3. 结构体与泛型

结构体可以使用泛型参数：

```rust
struct Point<T> {
    x: T,
    y: T,
}

impl<T> Point<T> {
    fn new(x: T, y: T) -> Self {
        Point { x, y }
    }
    
    fn x(&self) -> &T {
        &self.x
    }
}

// 为特定类型实现方法
impl Point<f32> {
    fn distance_from_origin(&self) -> f32 {
        (self.x.powi(2) + self.y.powi(2)).sqrt()
    }
}

fn main() {
    let integer = Point::new(5, 10);
    let float = Point::new(1.0, 4.0);
    
    println!("float distance: {}", float.distance_from_origin());
}
```

## 实践示例

### 1. 构建一个简单的用户管理系统

```rust
#[derive(Debug, Clone, PartialEq)]
struct User {
    id: u32,
    username: String,
    email: String,
    active: bool,
}

impl User {
    fn new(id: u32, username: String, email: String) -> Self {
        User {
            id,
            username,
            email,
            active: true,
        }
    }
    
    fn deactivate(&mut self) {
        self.active = false;
    }
    
    fn is_active(&self) -> bool {
        self.active
    }
    
    fn update_email(&mut self, new_email: String) {
        self.email = new_email;
    }
}

fn main() {
    let mut user = User::new(
        1, 
        String::from("alice"), 
        String::from("alice@example.com")
    );
    
    println!("User: {:?}", user);
    println!("Is active: {}", user.is_active());
    
    user.update_email(String::from("alice@newdomain.com"));
    user.deactivate();
    
    println!("Updated user: {:?}", user);
    println!("Is active: {}", user.is_active());
}
```

### 2. 几何图形计算系统

```rust
#[derive(Debug)]
struct Rectangle {
    width: f64,
    height: f64,
}

#[derive(Debug)]
struct Circle {
    radius: f64,
}

trait Area {
    fn area(&self) -> f64;
    fn perimeter(&self) -> f64;
}

impl Area for Rectangle {
    fn area(&self) -> f64 {
        self.width * self.height
    }
    
    fn perimeter(&self) -> f64 {
        2.0 * (self.width + self.height)
    }
}

impl Area for Circle {
    fn area(&self) -> f64 {
        std::f64::consts::PI * self.radius * self.radius
    }
    
    fn perimeter(&self) -> f64 {
        2.0 * std::f64::consts::PI * self.radius
    }
}

impl Rectangle {
    fn new(width: f64, height: f64) -> Self {
        Rectangle { width, height }
    }
    
    fn square(size: f64) -> Self {
        Rectangle { width: size, height: size }
    }
    
    fn can_hold(&self, other: &Rectangle) -> bool {
        self.width > other.width && self.height > other.height
    }
}

impl Circle {
    fn new(radius: f64) -> Self {
        Circle { radius }
    }
}

fn print_area_info<T: Area + std::fmt::Debug>(shape: &T) {
    println!("Shape: {:?}", shape);
    println!("Area: {:.2}", shape.area());
    println!("Perimeter: {:.2}", shape.perimeter());
    println!("---");
}

fn main() {
    let rect = Rectangle::new(10.0, 20.0);
    let square = Rectangle::square(15.0);
    let circle = Circle::new(5.0);
    
    print_area_info(&rect);
    print_area_info(&square);
    print_area_info(&circle);
    
    println!("Can rect hold square? {}", rect.can_hold(&square));
}
```

## 最佳实践

### 1. 结构体设计原则

```rust
// 好的设计：字段有意义，方法职责单一
#[derive(Debug, Clone)]
struct BankAccount {
    account_number: String,
    balance: f64,
    owner: String,
}

impl BankAccount {
    fn new(account_number: String, owner: String) -> Self {
        BankAccount {
            account_number,
            balance: 0.0,
            owner,
        }
    }
    
    fn deposit(&mut self, amount: f64) -> Result<(), String> {
        if amount <= 0.0 {
            return Err("Amount must be positive".to_string());
        }
        self.balance += amount;
        Ok(())
    }
    
    fn withdraw(&mut self, amount: f64) -> Result<(), String> {
        if amount <= 0.0 {
            return Err("Amount must be positive".to_string());
        }
        if amount > self.balance {
            return Err("Insufficient funds".to_string());
        }
        self.balance -= amount;
        Ok(())
    }
    
    fn get_balance(&self) -> f64 {
        self.balance
    }
}
```

### 2. 合理使用所有权

```rust
#[derive(Debug)]
struct Library {
    name: String,
    books: Vec<Book>,
}

#[derive(Debug, Clone)]
struct Book {
    title: String,
    author: String,
    isbn: String,
}

impl Library {
    fn new(name: String) -> Self {
        Library {
            name,
            books: Vec::new(),
        }
    }
    
    fn add_book(&mut self, book: Book) {
        self.books.push(book);
    }
    
    fn find_book(&self, isbn: &str) -> Option<&Book> {
        self.books.iter().find(|book| book.isbn == isbn)
    }
    
    fn remove_book(&mut self, isbn: &str) -> Option<Book> {
        if let Some(pos) = self.books.iter().position(|book| book.isbn == isbn) {
            Some(self.books.remove(pos))
        } else {
            None
        }
    }
    
    fn book_count(&self) -> usize {
        self.books.len()
    }
}
```

## 总结

结构体是 Rust 中组织相关数据的强大工具：

1. **数据组织**：将相关数据组合成有意义的类型
2. **方法系统**：通过 `impl` 块为结构体添加行为
3. **关联函数**：提供构造函数和其他实用函数
4. **所有权集成**：与 Rust 的所有权系统完美配合
5. **trait 支持**：可以实现和推导各种 trait
6. **模式匹配**：支持解构和匹配

掌握结构体是学习 Rust 面向对象编程风格的基础，它为构建复杂的数据类型和系统提供了坚实的基础。

