# C#概述


C# 是一款开源且免费的编程语言，基于流行的 C 风格语法（如 C、C++、Java），让开发者可以快速上手。它强大的跨平台能力支持开发多样化的应用，包括桌面软件、Web 应用、移动应用、微服务、物联网设备和游戏主机平台。得益于其开源特性，开发者不仅能自由使用，还能查看、修改源代码，并向社区贡献改进。

## C#语法基础

### **关键字（Keywords）**

C#关键字是编译器保留的特殊单词，分为两类：

- **保留关键字**：如 `class`、`if`、`for`、`void` 等，用于定义程序结构。
- **上下文关键字**：如 `var`、`async`、`await`，仅在特定场景下被视为关键字。

**关键作用**：

- 声明类型（`class`、`struct`、`enum`）
- 控制流程（`if`、`switch`、`while`）
- 定义访问权限（`public`、`private`）
- 处理异常（`try`、`catch`）

**示例**：

```C#
public class Program  // 'public'和'class'为关键字
{
    static void Main() 
    {
        if (true) { }  // 'if'控制逻辑分支
    }
}
```

### **标识符（Identifiers）**

标识符是开发者自定义的名称（如变量、类名），规则如下：

- **组成**：字母、数字、下划线`_`或`@`符号（`@`仅用于转义关键字，如 `@class`）
- **开头字符**：字母或下划线（**禁止数字开头**）
- **大小写敏感**：`myVar` 与 `MyVar` 不同
- **命名规范**：
  - 类名/方法名 → **PascalCase**（`MyClass`、`CalculateSum()`）
  - 变量/参数名 → **camelCase**（`userAge`、`isValid`）
  - 常量 → **全大写**（`MAX_SIZE`）

**示例**：

```c#
string firstName;    // 变量：camelCase
const int MAX_USERS = 100; // 常量：全大写
class CustomerOrder { }   // 类名：PascalCase
```

### **Main方法（Entry Point）**

`Main` 是程序执行的起点，必须满足：

- **固定签名**：`static void Main(string[] args)`（`args` 可省略）
- **静态方法**：用 `static` 修饰，无需实例化类
- **唯一性**：每个程序仅允许一个 `Main` 方法[1,3,10](https://tencent.yuanbao/@ref)

**示例**：

```c#
using System;
class App
{
    static void Main()  // 入口方法
    {
        Console.WriteLine("Hello, C#!"); 
    }
}
```

### **语句与语句分隔符**

- 语句（Statements）：程序执行的最小单元，如：
  - **声明语句**：`int x = 5;`
  - **赋值语句**：`x = 10;`
  - **控制流语句**：`if`、`for`、`return`
- 分隔符：
  - **分号 `;`**：标记语句结束（**不可省略**）
  - **大括号 `{}`**：定义代码块（如类体、循环体）[4,5](https://tencent.yuanbao/@ref)

**示例**：

```c#
if (x > 0)           // if条件控制
{                    // 代码块开始
    Console.WriteLine(x);
    return;          // 返回语句
}                    // 代码块结束
```

### **空白（Whitespace）**

空白（空格、制表符、换行）用于提升代码可读性：

- **编译器忽略**：不影响程序逻辑（如 `int x=5;` 等同 `int x = 5;`）
- 规范建议：
  - 操作符两侧加空格（`int sum = a + b;`）
  - 逗号后加空格（`Method(arg1, arg2)`）
  - 缩进代码块（通常4个空格或1个制表符）

**对比示例**：

```c#
// 紧凑写法（合法但难读）
int x=5;if(x>0){Console.WriteLine(x);}

// 规范写法
int x = 5;
if (x > 0) 
{
    Console.WriteLine(x);
}
```





---

> 作者: [hobby](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/basic/d14b1543-d65d-40cb-a7c4-f3bed1e0dfc6/  

