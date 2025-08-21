---
type: posts
title: "Rust所有权系统"
date: 2025-08-16T22:57:58+0800
slug: "38984233-be8b-4f04-bf7a-1a0f947eb699"
draft: false
author: 
  name: hao
  link: https://github.com/haochan1996
  email: espholychan@outllook.com
  avatar: https://avatars.githubusercontent.com/u/190246046?v=4
description: 深入理解Rust所有权系统，包括所有权规则、借用、生命周期等核心概念
keywords:
license:
comment: false
weight: 2
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

## 什么是所有权

所有权（ownership）是 Rust 最为与众不同的特性，它让 Rust 无需垃圾回收（garbage collector）即可保障内存安全。因此，理解所有权如何工作是十分重要的。

在其他系统编程语言中，如 C 和 C++，程序员需要手动管理内存的分配和释放，这容易导致内存泄漏、悬垂指针等问题。而在带有垃圾回收的语言中，如 Java、Python，虽然不用手动管理内存，但会有运行时性能开销。

Rust 选择了第三条道路：通过所有权系统管理内存，编译器在编译时会根据一系列的规则进行检查。如果违反了任何这些规则，程序都不能编译。在运行时，所有权系统的任何功能都不会减慢程序。

### 栈（Stack）与堆（Heap）

在深入所有权之前，让我们了解一下栈和堆。这些概念对于理解所有权系统至关重要。

**栈（Stack）**：
- 存储数据的顺序是后进先出（LIFO）
- 所有存储在栈上的数据都必须是已知且固定的大小
- 访问速度快，因为数据总是在栈顶添加或移除

**堆（Heap）**：
- 用于存储编译时大小未知或大小可能变化的数据
- 分配内存时，内存分配器在堆的某处找到一块足够大的空位，把它标记为已使用，并返回一个指针
- 访问速度较慢，因为必须先访问指针，然后跟踪指针到达数据

```rust
fn main() {
    // 这些值存储在栈上，因为它们有已知的固定大小
    let x = 5;
    let y = true;
    let z = 'a';
    
    // 这个字符串字面值存储在程序二进制文件中，变量存储的是指向它的引用
    let s1 = "hello";
    
    // 这个 String 的数据存储在堆上，变量存储的是指向堆数据的指针
    let s2 = String::from("hello");
}
```

## 所有权规则

Rust 的所有权有三个基本规则：

1. **Rust 中的每一个值都有一个被称为其所有者（owner）的变量**
2. **值在任一时刻有且只有一个所有者**
3. **当所有者（变量）离开作用域，这个值将被丢弃**

让我们通过例子来理解这些规则：

### 变量作用域

作用域是一个项（item）在程序中有效的范围：

```rust
fn main() {
    {                      // s 在这里无效，它尚未声明
        let s = "hello";   // 从此处起，s 是有效的

        // 使用 s
        println!("{}", s);
    }                      // 此作用域已结束，s 不再有效
}
```

当变量离开作用域，Rust 为我们调用一个特殊的函数 `drop`，在这个函数中可以放置释放内存的代码。

### String 类型

为了演示所有权规则，我们需要一个比之前介绍的都要复杂的数据类型。之前介绍的类型都是存储在栈上的，当离开作用域时被移出栈。而现在我们要看看存储在堆上的数据。

```rust
fn main() {
    // 字符串字面值是不可变的
    let s1 = "hello";
    
    // String 类型是可变的，存储在堆上
    let mut s2 = String::from("hello");
    s2.push_str(", world!"); // push_str() 在字符串后追加字面值
    
    println!("{}", s2); // 将打印 `hello, world!`
}
```

### 内存与分配

对于`String`类型，为了支持一个可变，可增长的文本片段，需要在堆上分配一块在编译时未知大小的内存来存放内容。这意味着：

1. 必须在运行时向内存分配器（memory allocator）请求内存
2. 需要一个当我们处理完`String`时将内存返回给分配器的方法

第一部分由我们完成：当调用`String::from`时，它的实现请求其所需的内存。

第二部分在大部分语言中，要么有垃圾回收器，要么我们必须手动释放。Rust采取了一个不同的策略：内存在拥有它的变量离开作用域后就被自动释放。

```rust
fn main() {
    {
        let s = String::from("hello"); // 从此处起，s 是有效的

        // 使用 s
        println!("{}", s);
    }                                  // 此作用域已结束，
                                       // s 不再有效，内存被释放
}
```

## 变量与数据的交互方式

### 移动（Move）

在Rust中，多个变量可以采用不同的方式与同一数据进行交互：

```rust
fn main() {
    // 对于简单的标量值，这样做没问题
    let x = 5;
    let y = x; // x 的值被复制给 y，现在栈上有两个值都是 5
    
    println!("x = {}, y = {}", x, y); // 这样使用是可以的
}
```

但是对于`String`类型：

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1; // s1 的值移动到了 s2
    
    // println!("{}", s1); // 这行代码会报错！s1 不再有效
    println!("{}", s2); // 这是可以的
}
```

为什么会这样？当我们将`s1`赋给`s2`，`String`的数据被复制了，这意味着我们从栈上拷贝了它的指针、长度和容量。我们并没有复制指针指向的堆上数据。

<img src="https://doc.rust-lang.org/book/img/trpl04-04.svg" alt="String移动示意图" />

当变量离开作用域后，Rust自动调用`drop`函数并清理变量的堆内存。但是当`s2`和`s1`离开作用域，它们都会尝试释放相同的内存。这是一个叫做二次释放（double free）的错误。

为了确保内存安全，在`let s2 = s1`之后，Rust认为`s1`不再有效，因此Rust不需要在`s1`离开作用域后清理任何东西。这在Rust中被称为**移动（move）**。

### 克隆（Clone）

如果我们确实需要深度复制`String`中堆上的数据，而不仅仅是栈上的数据，可以使用一个叫做`clone`的通用函数：

```rust
fn main() {
    let s1 = String::from("hello");
    let s2 = s1.clone(); // 深拷贝堆上的数据
    
    println!("s1 = {}, s2 = {}", s1, s2); // 现在两个都有效
}
```

### 拷贝（Copy）

对于像整型这样的类型，它们完全存储在栈上，所以拷贝其实际的值是快速的：

```rust
fn main() {
    let x = 5;
    let y = x;
    
    println!("x = {}, y = {}", x, y); // 都有效
}
```

这里没有调用`clone`，不过`x`依然有效且没有被移动到`y`中。原因是像整型这样的类型在编译时是已知大小的，会被存储在栈上，所以拷贝其实际的值是快速的。

Rust 有一个叫做`Copy`trait 的特殊注解，可以用在类似整型这样的存储在栈上的类型上。如果一个类型实现了`Copy`trait，那么一个旧的变量在将其赋值给其他变量后仍然可用。

以下是一些 `Copy` 的类型：
- 所有整数类型，比如`u32`
- 布尔类型，`bool`
- 所有浮点数类型，比如`f64`
- 字符类型，`char`
- 元组，当且仅当其包含的类型也都实现`Copy`的时候。比如，`(i32,i32)`实现了`Copy`，但`(i32,String)`就没有

## 所有权与函数

将值传递给函数在语义上与给变量赋值相似。向函数传递值可能会移动或者复制，就像赋值语句一样：

```rust
fn main() {
    let s = String::from("hello");  // s 进入作用域

    takes_ownership(s);             // s 的值移动到函数里
                                    // 所以到这里不再有效

    let x = 5;                      // x 进入作用域

    makes_copy(x);                  // x 应该移动函数里，
                                    // 但 i32 是 Copy 的，
                                    // 所以在后面可继续使用 x

    println!("x = {}", x);          // 这里可以使用 x
    // println!("s = {}", s);       // 这里不能使用 s，会编译错误

} // 这里，x 先移出了作用域，然后是 s。但因为 s 的值已经被移动，没什么特殊之处

fn takes_ownership(some_string: String) { // some_string 进入作用域
    println!("{}", some_string);
} // 这里，some_string 移出作用域并调用 `drop` 方法。占用的内存被释放

fn makes_copy(some_integer: i32) { // some_integer 进入作用域
    println!("{}", some_integer);
} // 这里，some_integer 移出作用域。没什么特殊之处
```

### 返回值与作用域

返回值也可以转移所有权：

```rust
fn main() {
    let s1 = gives_ownership();         // gives_ownership 将返回值
                                        // 移给 s1

    let s2 = String::from("hello");     // s2 进入作用域

    let s3 = takes_and_gives_back(s2);  // s2 被移动到
                                        // takes_and_gives_back 中,
                                        // 它也将返回值移给 s3
    
    println!("s1 = {}, s3 = {}", s1, s3);
    // println!("s2 = {}", s2);         // s2 已经被移动，不能使用
} // 这里，s3 移出作用域并被丢弃。s2 也移出作用域，但已被移走，
  // 所以什么也不会发生。s1 移出作用域并被丢弃

fn gives_ownership() -> String {             // gives_ownership 将返回值移动给
                                             // 调用它的函数

    let some_string = String::from("yours"); // some_string 进入作用域

    some_string                              // 返回 some_string 并移出给调用的函数
}

// takes_and_gives_back 将传入字符串并返回该值
fn takes_and_gives_back(a_string: String) -> String { // a_string 进入作用域

    a_string  // 返回 a_string 并移出给调用的函数
}
```

变量的所有权总是遵循相同的模式：将值赋给另一个变量时移动它。当持有堆中数据值的变量离开作用域时，其值将通过`drop`被清理掉，除非数据被移动为另一个变量所有。

## 引用与借用

如果我们想要函数使用一个值但不获取所有权该怎么办？如果我们想在调用函数后还能继续使用它，那传入一个引用可能是更好的选择。

### 引用（References）

**引用**（reference）像一个指针，因为它是一个地址，我们可以由此访问储存于该地址的属于其他变量的数据。但是引用确保指向某个特定类型的有效值。

```rust
fn main() {
    let s1 = String::from("hello");

    let len = calculate_length(&s1);

    println!("The length of '{}' is {}.", s1, len);
    // s1 仍然有效，因为我们只是借用了它
}

fn calculate_length(s: &String) -> usize { // s 是对 String 的引用
    s.len()
} // 这里，s 离开了作用域。但因为它并不拥有引用值的所有权，
  // 所以什么也不会发生
```

注意我们传递 `&s1` 给 `calculate_length`，同时在函数定义中，我们获取 `&String` 而不是 `String`。这些 `&` 符号就是**引用**，它们允许你使用值但不获取其所有权。

![引用示意图](https://doc.rust-lang.org/book/img/trpl04-05.svg)

我们将创建一个引用的行为称为**借用**（borrowing）。正如现实生活中，如果一个人拥有某样东西，你可以从他那里借来。当你使用完毕，必须还回去。

### 可变引用

默认情况下，引用是不可变的。但我们也可以创建可变引用：

```rust
fn main() {
    let mut s = String::from("hello");

    change(&mut s);
    
    println!("{}", s); // 输出 "hello, world"
}

fn change(some_string: &mut String) {
    some_string.push_str(", world");
}
```

不过可变引用有一个很大的限制：**在同一时间只能有一个对某一特定数据的可变引用**：

```rust
fn main() {
    let mut s = String::from("hello");

    let r1 = &mut s;
    // let r2 = &mut s; // 错误！不能同时有两个可变引用

    println!("{}", r1);
}
```

这个限制的好处是 Rust 可以在编译时就避免数据竞争。**数据竞争**（data race）类似于竞态条件，它可由这三个行为造成：

1. 两个或更多指针同时访问同一数据
2. 至少有一个指针被用来写入数据
3. 没有同步数据访问的机制

我们也不能在拥有不可变引用的同时拥有可变引用：

```rust
fn main() {
    let mut s = String::from("hello");

    let r1 = &s; // 没问题
    let r2 = &s; // 没问题
    // let r3 = &mut s; // 大问题！不能在有不可变引用时创建可变引用

    println!("{} and {}", r1, r2);
}
```

但是，多个不可变引用是可以的，因为没有哪个只能读取数据的人有能力影响其他人读取到的数据。

注意一个引用的作用域从声明的地方开始一直持续到最后一次使用为止：

```rust
fn main() {
    let mut s = String::from("hello");

    let r1 = &s; // 没问题
    let r2 = &s; // 没问题
    println!("{} and {}", r1, r2);
    // 此位置之后 r1 和 r2 不再使用

    let r3 = &mut s; // 没问题
    println!("{}", r3);
}
```

### 悬垂引用（Dangling References）

在具有指针的语言中，很容易通过释放内存时保留指向它的指针而错误地生成一个**悬垂指针**（dangling pointer），所谓悬垂指针是其指向的内存可能已经被分配给其它持有者。相比之下，在 Rust 中编译器确保引用永远也不会变成悬垂引用：当你拥有一些数据的引用，编译器确保数据不会在其引用之前离开作用域。

```rust
fn main() {
    // let reference_to_nothing = dangle(); // 这会编译错误
}

fn dangle() -> &String { // dangle 返回一个字符串的引用
    let s = String::from("hello"); // s 是一个新字符串

    &s // 返回字符串 s 的引用
} // 这里 s 离开作用域并被丢弃。其内存被释放。
  // 危险！返回的引用指向了无效的内存
```

解决方案是直接返回 `String`：

```rust
fn no_dangle() -> String {
    let s = String::from("hello");

    s // 返回 String，所有权被移出
}
```

## 切片（Slice）类型

slice 允许你引用集合中一段连续的元素序列，而不用引用整个集合。slice 是一类引用，所以它没有所有权。

### 字符串 slice

**字符串 slice**（string slice）是 `String` 中一部分值的引用：

```rust
fn main() {
    let s = String::from("hello world");

    let hello = &s[0..5];  // 或 &s[..5]
    let world = &s[6..11]; // 或 &s[6..]
    let whole = &s[..];    // 整个字符串的 slice

    println!("{} {}", hello, world); // hello world
}
```

slice 的语法：`&s[starting_index..ending_index]`，其中 `starting_index` 是 slice 的第一个位置，`ending_index` 则是 slice 最后一个位置的后一个值。

![字符串slice示意图](https://doc.rust-lang.org/book/img/trpl04-06.svg)

让我们用 slice 重写之前的 `first_word` 函数：

```rust
fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

fn main() {
    let mut s = String::from("hello world");

    let word = first_word(&s);

    // s.clear(); // 错误！不能在有不可变引用时修改字符串

    println!("the first word is: {}", word);
}
```

### 字符串字面值就是 slice

还记得我们讲到过字符串字面值被储存在二进制文件中吗？现在知道 slice 了，我们就可以正确了解字符串字面值了：

```rust
let s = "Hello, world!";
```

这里 `s` 的类型是 `&str`：它是一个指向二进制程序特定位置的 slice。这也就是为什么字符串字面值是不可变的；`&str` 是一个不可变引用。

### 字符串 slice 作为参数

如果有一个字符串 slice，可以直接传递它。如果有一个 `String`，则可以传递整个 `String` 的 slice 或对 `String` 的引用。这种灵活性利用了 **解引用强制转换**（deref coercions）：

```rust
fn first_word(s: &str) -> &str { // 改为接受 &str 类型
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

fn main() {
    let my_string = String::from("hello world");

    // first_word 中传入 `String` 的 slice
    let word = first_word(&my_string[0..6]);
    let word = first_word(&my_string[..]);
    // first_word 也接受 `String` 的引用，
    // 这等同于 `String` 的 slice
    let word = first_word(&my_string);

    let my_string_literal = "hello world";

    // first_word 中传入字符串字面值的 slice
    let word = first_word(&my_string_literal[0..6]);
    let word = first_word(&my_string_literal[..]);

    // 因为字符串字面值**就是**字符串 slice，
    // 这样写也可以，即不使用 slice 语法！
    let word = first_word(my_string_literal);
}
```

### 其他类型的 slice

字符串 slice，正如你想象的那样，是针对字符串的。不过也有更通用的 slice 类型：

```rust
fn main() {
    let a = [1, 2, 3, 4, 5];

    let slice = &a[1..3]; // 类型是 &[i32]

    assert_eq!(slice, &[2, 3]);
}
```

这个 slice 的类型是 `&[i32]`。它跟字符串 slice 的工作方式一样，通过存储第一个集合元素的引用和一个集合长度。

## 实践示例和最佳实践

### 1. 所有权转移的实践

```rust
#[derive(Debug)]
struct User {
    name: String,
    email: String,
}

fn create_user(name: String, email: String) -> User {
    User { name, email } // 所有权转移到新的 User 实例
}

fn print_user_info(user: &User) { // 借用而不获取所有权
    println!("User: {} <{}>", user.name, user.email);
}

fn update_email(user: &mut User, new_email: String) {
    user.email = new_email; // 修改通过可变引用进行
}

fn main() {
    let name = String::from("Alice");
    let email = String::from("alice@example.com");
    
    let mut user = create_user(name, email); // name 和 email 的所有权被转移
    // println!("{}", name); // 错误！name 已经被移动
    
    print_user_info(&user); // 借用 user
    
    update_email(&mut user, String::from("alice@newdomain.com"));
    print_user_info(&user);
}
```

### 2. 避免不必要的克隆

```rust
// 不好的做法
fn process_data_bad(data: Vec<String>) -> Vec<String> {
    let mut result = Vec::new();
    for item in data {
        let processed = item.clone(); // 不必要的克隆
        result.push(processed);
    }
    result
}

// 好的做法
fn process_data_good(data: Vec<String>) -> Vec<String> {
    data.into_iter()
        .map(|item| item) // 直接移动所有权
        .collect()
}

// 或者如果需要保留原始数据
fn process_data_with_borrow(data: &[String]) -> Vec<String> {
    data.iter()
        .map(|item| item.clone()) // 只在必要时克隆
        .collect()
}
```

### 3. 合理使用引用

```rust
fn analyze_text(text: &str) -> (usize, usize, usize) {
    let word_count = text.split_whitespace().count();
    let char_count = text.chars().count();
    let line_count = text.lines().count();
    
    (word_count, char_count, line_count)
}

fn main() {
    let content = String::from("Hello world\nThis is Rust\nOwnership is powerful");
    
    let (words, chars, lines) = analyze_text(&content); // 借用而不移动
    
    println!("Words: {}, Chars: {}, Lines: {}", words, chars, lines);
    println!("Original content is still available: {}", content);
}
```

### 4. 使用切片提高灵活性

```rust
fn find_max_value(numbers: &[i32]) -> Option<i32> {
    if numbers.is_empty() {
        None
    } else {
        Some(*numbers.iter().max().unwrap())
    }
}

fn main() {
    let vec = vec![1, 5, 3, 9, 2];
    let array = [10, 20, 30, 40];
    
    // 可以接受 Vec 的引用
    if let Some(max) = find_max_value(&vec) {
        println!("Max in vec: {}", max);
    }
    
    // 也可以接受数组的引用
    if let Some(max) = find_max_value(&array) {
        println!("Max in array: {}", max);
    }
    
    // 甚至可以接受切片
    if let Some(max) = find_max_value(&vec[1..4]) {
        println!("Max in slice: {}", max);
    }
}
```

## 总结

所有权是 Rust 的核心概念，它通过编译时检查确保内存安全。理解并掌握所有权系统是学习 Rust 的关键：

1. **所有权规则**：每个值都有唯一的所有者，当所有者离开作用域时值被丢弃
2. **移动语义**：赋值和函数调用会转移所有权，避免了数据竞争
3. **借用系统**：通过引用可以使用值而不获取所有权
4. **生命周期**：编译器确保引用总是有效的
5. **切片类型**：提供了对集合部分数据的引用

掌握这些概念将帮助你编写高效、安全的 Rust 代码，避免常见的内存管理问题。
