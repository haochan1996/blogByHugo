# C#概述


### 1. C# 简介
C# 是一种简单、现代、通用、面向对象的编程语言，由微软开发，旨在运行于 .NET 框架之上。它支持跨平台开发，适用于 Windows、macOS 和 Linux。C# 的设计目标包括类型安全、垃圾回收和丰富的类库，适合初学者快速上手。

- **历史与发展**：C# 1.0 于 2002 年发布，最新版本为 C# 13（2024 年 11 月发布），与 .NET 9（2025 年 8 月 5 日发布）兼容。版本历史见 [https://learn.microsoft.com/zh-cn/dotnet/csharp/whats-new/csharp-version-history](https://learn.microsoft.com/zh-cn/dotnet/csharp/whats-new/csharp-version-history)。
- **关键特性**：自动垃圾回收、面向对象编程支持（封装、继承、多态）、强大的标准库、异步编程支持等。

### 2. 开发环境设置
学习 C# 的第一步是设置开发环境。以下是详细步骤：

- **下载 .NET SDK**：.NET SDK 是 C# 开发的核心组件。最新版本为 .NET 9.0.8（2025-08-05，标准支持）和 .NET 8.0.19（2025-08-05，长期支持）。下载地址：[https://dotnet.microsoft.com/zh-cn/download](https://dotnet.microsoft.com/zh-cn/download)。安装后，您可以通过命令行运行 `dotnet --version` 验证。
- **选择 IDE**：
  - **Visual Studio Code**：轻量级跨平台编辑器，适合初学者。下载地址：[https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)。安装 C# Dev Kit 扩展：[https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csdevkit](https://marketplace.visualstudio.com/items?itemName=ms-dotnettools.csdevkit)。
  - **Visual Studio**：功能强大的 IDE，适合 Windows 和 macOS 用户。下载地址：[https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=Community&rel=17](https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?sku=Community&rel=17)。
- **在线编译器**：如果不希望安装本地环境，可以使用在线工具，如菜鸟教程提供的编译器：[https://www.runoob.com/try/showcs.php?filename=HelloWorld](https://www.runoob.com/try/showcs.php?filename=HelloWorld)，适合初学者快速测试代码。

### 3. 基本语法
C# 的基本语法包括变量声明、数据类型、运算符和控制结构。以下是详细内容：

#### 3.1 Hello World 程序
这是 C# 的入门程序，展示如何输出文本：

```csharp
using System;

class Program
{
    static void Main()
    {
        Console.WriteLine("Hello, World!");
    }
}
```

- `using System;` 引入系统命名空间。
- `Main` 是程序的入口点。
- `Console.WriteLine` 用于输出文本到控制台。

从 Microsoft Learn 的交互式教程中，扩展学习包括字符串操作，如变量声明和插值：

```csharp
string aFriend = "Bill";
Console.WriteLine(aFriend); // 输出 Bill
aFriend = "Maira";
Console.WriteLine($"Hello {aFriend}"); // 输出 Hello Maira
```

字符串方法如 `Trim`、`Replace`、`ToUpper` 等也非常实用，详见 [https://learn.microsoft.com/zh-cn/dotnet/csharp/tour-of-csharp/tutorials/hello-world](https://learn.microsoft.com/zh-cn/dotnet/csharp/tour-of-csharp/tutorials/hello-world)。

#### 3.2 数据类型和变量
C# 支持多种数据类型，初学者需掌握以下常见类型：

- **整型**：`int`（32 位整数，如 `int a = 18;`），范围见 `int.MaxValue` 和 `int.MinValue`。
- **浮点型**：`double`（双精度浮点数，如 `double b = 4.5;`），适合科学计算；`decimal`（高精度，如 `decimal c = 1.0M;`），用于金融计算。
- **字符串**：`string`（如 `string name = "Alice";`），支持插值和方法如 `Length`、`Contains`。

数学运算示例：

```csharp
int a = 18; int b = 6;
int c = a + b; // c = 24
Console.WriteLine(c);
```

注意溢出问题，如 `int.MaxValue + 3` 会导致负数结果，详见 [https://learn.microsoft.com/zh-cn/dotnet/csharp/tour-of-csharp/tutorials/numbers](https://learn.microsoft.com/zh-cn/dotnet/csharp/tour-of-csharp/tutorials/numbers)。

#### 3.3 运算符
C# 支持算术运算符（`+`, `-`, `*`, `/`, `%`）、比较运算符（`==`, `!=`, `>`, `<`）和逻辑运算符（`&&`, `||`, `!`）。运算顺序遵循数学规则，可用括号调整优先级。

示例：

```csharp
int d = a + b * c; // 先乘后加
int e = (a + b) % c; // 取余
```

#### 3.4 控制结构
控制结构包括条件判断和循环：
- **if 语句**：
  ```csharp
  if (age >= 18)
  {
      Console.WriteLine("成年人");
  }
  else
  {
      Console.WriteLine("未成年人");
  }
  ```
- **switch 语句**：用于多分支选择，示例见上文。
- **循环**：`for` 循环示例：
  ```csharp
  for (int i = 0; i < 5; i++)
  {
      Console.WriteLine(i); // 输出 0 到 4
  }
  ```

### 4. 面向对象编程
C# 是面向对象的，支持封装、继承和多态。以下是详细介绍：

#### 4.1 类和对象
类是对象的蓝图，对象是类的实例。示例：

```csharp
public class Person
{
    public string Name { get; set; }
    public int Age { get; set; }

    public void SayHello()
    {
        Console.WriteLine($"你好，我叫 {Name}，今年 {Age} 岁。");
    }
}
```

创建对象：

```csharp
Person person = new Person { Name = "Alice", Age = 30 };
person.SayHello(); // 输出：你好，我叫 Alice，今年 30 岁。
```

#### 4.2 继承
继承允许子类复用父类的属性和方法。示例：

```csharp
public class Student : Person
{
    public string StudentId { get; set; }
}
```

从 Microsoft Learn 的银行系统示例中，`InterestEarningAccount` 和 `LineOfCreditAccount` 继承自 `BankAccount`，利率分别为 2% 和 7%，详见 [https://learn.microsoft.com/zh-cn/dotnet/csharp/fundamentals/tutorials/oop](https://learn.microsoft.com/zh-cn/dotnet/csharp/fundamentals/tutorials/oop)。

#### 4.3 多态
多态通过 `virtual` 和 `override` 实现方法重写。示例：

```csharp
public class Animal
{
    public virtual void MakeSound()
    {
        Console.WriteLine("动物叫");
    }
}

public class Dog : Animal
{
    public override void MakeSound()
    {
        Console.WriteLine("汪汪");
    }
}
```

#### 4.4 封装
通过访问修饰符（如 `public`、`private`）控制成员访问，保护内部实现。

### 5. 高级主题
初学者可逐步学习以下高级主题：
- **异常处理**：使用 `try-catch` 捕获异常，例如：
  ```csharp
  try
  {
      int result = 10 / 0;
  }
  catch (DivideByZeroException e)
  {
      Console.WriteLine("除数不能为零：" + e.Message);
  }
  ```
- **文件 I/O**：读写文件，使用 `System.IO` 命名空间。
- **LINQ**：语言集成查询，用于数据操作，示例见官方文档。
- **异步编程**：使用 `async` 和 `await` 处理异步任务，适合网络操作。

### 6. 进一步学习资源
以下是推荐的学习资源：
- **官方文档**：[https://learn.microsoft.com/zh-cn/dotnet/csharp/](https://learn.microsoft.com/zh-cn/dotnet/csharp/)，提供交互式教程和参考资料。
- **在线教程**：[https://www.runoob.com/csharp/csharp-tutorial.html](https://www.runoob.com/csharp/csharp-tutorial.html)，适合初学者，包含在线编译器。
- **书籍**：推荐《C# 7.0 核心技术指南》和《Head First C#》，深入学习高级概念。
- **社区与动态**：关注 C# 技术大会（如 2025 年会议，主题为智能创新）和 GitHub 仓库（如 [https://github.com/dotnet/csharplang](https://github.com/dotnet/csharplang)），了解最新动态。 [https://learn.microsoft.com/zh-cn/dotnet/csharp/whats-new/csharp-version-history](https://learn.microsoft.com/zh-cn/dotnet/csharp/whats-new/csharp-version-history)。


---

> 作者: [hao](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/basic/d14b1543-d65d-40cb-a7c4-f3bed1e0dfc6/  

