# C#数据类型


## 1. 什么是数据类型？

数据类型定义了变量可以存储的数据种类、占用内存的大小以及可以对变量执行的操作。C# 中的数据类型主要分为两类：

- **值类型 (Value Types)**：直接存储数据，如整数、浮点数、布尔值等。
- **引用类型 (Reference Types)**：存储数据的引用（内存地址），如字符串、数组、类等。

## 2. 值类型 (Value Types)

值类型变量直接在栈内存中存储实际数据。以下是常见的内置值类型：

### 2.1 整数类型

整数类型用于存储没有小数部分的数字。C# 提供多种整数类型，区别在于大小和是否支持负数：

| 类型     | 占用字节 | 范围                                                    | 描述             |
| -------- | -------- | ------------------------------------------------------- | ---------------- |
| `byte`   | 1        | 0 到 255                                                | 无符号 8 位整数  |
| `sbyte`  | 1        | -128 到 127                                             | 有符号 8 位整数  |
| `short`  | 2        | -32,768 到 32,767                                       | 有符号 16 位整数 |
| `ushort` | 2        | 0 到 65,535                                             | 无符号 16 位整数 |
| `int`    | 4        | -2,147,483,648 到 2,147,483,647                         | 有符号 32 位整数 |
| `uint`   | 4        | 0 到 4,294,967,295                                      | 无符号 32 位整数 |
| `long`   | 8        | -9,223,372,036,854,775,808 到 9,223,372,036,854,775,807 | 有符号 64 位整数 |
| `ulong`  | 8        | 0 到 18,446,744,073,709,551,615                         | 无符号 64 位整数 |

**示例代码**：

```csharp
int age = 25;
long bigNumber = 1234567890123L; // 使用 L 表示 long 类型
byte smallNumber = 255;
```

### 2.2 浮点类型

浮点类型用于存储带小数部分的数字，适用于需要高精度的场景：

| 类型      | 占用字节 | 范围                  | 精度             |
| --------- | -------- | --------------------- | ---------------- |
| `float`   | 4        | ±1.5e-45 到 ±3.4e38   | 7 位有效数字     |
| `double`  | 8        | ±5.0e-324 到 ±1.7e308 | 15-16 位有效数字 |
| `decimal` | 16       | ±1.0e-28 到 ±7.9e28   | 28-29 位有效数字 |

**注意**：

- `float` 需要在数字后加 `f`（如 `3.14f`）。
- `decimal` 需要加 `m`（如 `3.14m`），适合金融计算等高精度场景。

**示例代码**：

```csharp
float temperature = 36.6f;
double pi = 3.14159265359;
decimal balance = 12345.6789m;
```

### 2.3 布尔类型

- `bool`：占用 1 字节，值只能是 `true` 或 `false`。
- 用于逻辑判断。

**示例代码**：

```csharp
bool isStudent = true;
bool hasLicense = false;
```

### 2.4 字符类型

- `char`：占用 2 字节，存储单个 Unicode 字符（用单引号 `''` 表示）。

**示例代码**：

```csharp
char grade = 'A';
char symbol = '\u0041'; // Unicode 表示 'A'
```

### 2.5 结构体 (Struct)

结构体是用户定义的值类型，包含多个字段。例如，`System.DateTime` 是一个结构体。

**示例代码**：

```csharp
struct Point
{
    public int X;
    public int Y;
}

Point p = new Point { X = 10, Y = 20 };
```

## 3. 引用类型 (Reference Types)

引用类型变量存储的是数据的内存地址，数据本身存储在堆内存中。常见的引用类型包括：

### 3.1 字符串 (String)

- `string`：表示不可变的 Unicode 字符序列。
- 使用双引号 `""` 定义。

**示例代码**：

```csharp
string name = "Alice";
string greeting = $"Hello, {name}!"; // 字符串插值
```

### 3.2 数组 (Array)

数组是固定大小的元素集合，元素类型必须一致。

**示例代码**：

```csharp
int[] numbers = new int[3] { 1, 2, 3 };
string[] names = { "Alice", "Bob", "Charlie" };
```

### 3.3 类 (Class)

类是用户定义的引用类型，可以包含字段、属性、方法等。

**示例代码**：

```csharp
class Person
{
    public string Name { get; set; }
    public int Age { get; set; }
}

Person person = new Person { Name = "Bob", Age = 30 };
```

### 3.4 对象 (Object)

- `object` 是所有类型的基类，任何类型都可以转换为 `object`。
- 使用时需注意装箱和拆箱（见下文）。

**示例代码**：

```csharp
object obj = 42; // 装箱
int number = (int)obj; // 拆箱
```

## 4. 值类型与引用类型的区别

| 特性         | 值类型                    | 引用类型                   |
| ------------ | ------------------------- | -------------------------- |
| **存储位置** | 栈内存                    | 堆内存（引用在栈中）       |
| **赋值行为** | 复制整个值                | 复制引用（指向同一对象）   |
| **默认值**   | 0 或相应类型的默认值      | `null`                     |
| **示例**     | `int`, `double`, `struct` | `string`, `class`, `array` |

**示例代码**（值类型与引用类型的赋值行为）：

```csharp
int a = 10;
int b = a;
b = 20; // a 仍然是 10
Console.WriteLine(a); // 输出 10

string[] arr1 = { "A", "B" };
string[] arr2 = arr1;
arr2[0] = "C"; // arr1[0] 也变为 "C"
Console.WriteLine(arr1[0]); // 输出 C
```

## 5. 类型转换

C# 中类型转换分为以下几种：

### 5.1 隐式转换

低精度类型到高精度类型的自动转换，无需显式声明。

**示例代码**：

```csharp
int i = 100;
double d = i; // 隐式转换
```

### 5.2 显式转换

高精度类型到低精度类型需要强制转换，可能丢失数据。

**示例代码**：

```csharp
double d = 123.45;
int i = (int)d; // 显式转换，i = 123
```

### 5.3 使用 `Convert` 类

`System.Convert` 类提供多种类型转换方法。

**示例代码**：

```csharp
string str = "123";
int num = Convert.ToInt32(str);
```

### 5.4 使用 `Parse` 和 `TryParse`

- `Parse`：将字符串转换为指定类型，若失败抛出异常。
- `TryParse`：尝试转换，若失败返回 `false`，不抛异常。

**示例代码**：

```csharp
string str = "123";
int result;
bool success = int.TryParse(str, out result); // success = true, result = 123
```

## 6. 装箱与拆箱 (Boxing and Unboxing)

- **装箱**：将值类型转换为 `object` 或接口类型，存储到堆内存。
- **拆箱**：从 `object` 类型转换回值类型。

**示例代码**：

```csharp
int i = 42;
object obj = i; // 装箱
int j = (int)obj; // 拆箱
```

**注意**：装箱和拆箱会影响性能，尽量避免在性能敏感的代码中使用。

## 7. 可空类型 (Nullable Types)

值类型默认不能为 `null`，但可通过可空类型（如 `int?`）支持 `null` 值。

**示例代码**：

```csharp
int? nullableInt = null;
if (nullableInt.HasValue)
{
    Console.WriteLine(nullableInt.Value);
}
else
{
    Console.WriteLine("Value is null");
}
```

**可空值类型简写**：

- `int?` 等价于 `Nullable<int>`。

## 8. 动态类型 (Dynamic Type)

`dynamic` 类型允许在运行时确定类型，绕过编译时类型检查。适用于与动态语言交互或未知类型的场景。

**示例代码**：

```csharp
dynamic value = 42;
value = "Now a string"; // 动态类型允许更改类型
Console.WriteLine(value);
```

**注意**：使用 `dynamic` 会降低代码可读性和性能，建议谨慎使用。

## 9. 字面量 (Literals)

字面量是直接在代码中书写的常量值，用于初始化变量或表示特定值。C# 支持多种类型的字面量，并允许使用前缀或后缀来指定数据类型或进制。

### 9.1 常见字面量类型

- 整数字面量：
  - 十进制：`123`
  - 十六进制：`0x7B`（以 `0x` 开头）
  - 二进制：`0b01111011`（以 `0b` 开头，C# 7.0+ 支持）
  - 后缀：`L`（`long`）、`UL`（`ulong`）、`U`（`uint`）。
- 浮点字面量：
  - 普通浮点：`3.14`
  - 科学计数法：`3.14e2`（表示 3.14 × 10² = 314）
  - 后缀：`f`（`float`）、`d`（`double`）、`m`（`decimal`）。
- **字符字面量**：`'A'` 或 `'\u0041'`（Unicode 字符）。
- 字符串字面量：
  - 普通字符串：`"Hello"`
  - 逐字字符串：`@"C:\Path"`（忽略转义字符）。
  - 插值字符串：`$"Value: {x}"`（C# 6.0+）。
- **布尔字面量**：`true`、`false`。
- **空字面量**：`null`（仅限引用类型或可空值类型）。

**示例代码**：

```csharp
int dec = 123;          // 十进制
int hex = 0x7B;         // 十六进制，等价于 123
int bin = 0b01111011;   // 二进制，等价于 123
long big = 123L;        // long 类型
float f = 3.14f;        // float 类型
decimal m = 3.14m;      // decimal 类型
string s = $"Value: {dec}"; // 插值字符串
```

### 9.2 字面量后缀和前缀

- 后缀明确数据类型：
  - `123L`：`long`
  - `123.45f`：`float`
  - `123.45m`：`decimal`
- 前缀指定进制：
  - `0x`：十六进制
  - `0b`：二进制（C# 7.0+）

## 10. 进制转换

C# 提供了多种方法在不同进制（如十进制、十六进制、二进制）之间转换数值。

### 10.1 使用字面量直接表示

如上所述，可以直接使用 `0x` 或 `0b` 前缀定义十六进制或二进制值。

**示例代码**：

```csharp
int hexValue = 0xFF; // 十六进制 FF = 十进制 255
int binValue = 0b11111111; // 二进制 11111111 = 十进制 255
```

### 10.2 使用 `Convert.ToString` 转换到指定进制

`Convert.ToString` 方法可以将数字转换为指定进制的字符串表示。

**示例代码**：

```csharp
int number = 255;
string binary = Convert.ToString(number, 2);  // 转换为二进制: "11111111"
string hex = Convert.ToString(number, 16);    // 转换为十六进制: "ff"
string octal = Convert.ToString(number, 8);   // 转换为八进制: "377"
```

### 10.3 使用 `Convert.ToInt32` 等从字符串解析

可以将字符串形式的其他进制数解析为十进制。

**示例代码**：

```csharp
string hexString = "FF";
int decValue = Convert.ToInt32(hexString, 16); // 解析十六进制 FF 为 255
string binString = "11111111";
int decValue2 = Convert.ToInt32(binString, 2); // 解析二进制 11111111 为 255
```

### 10.4 格式化输出

使用字符串格式化控制进制输出。

**示例代码**：

```csharp
int number = 255;
Console.WriteLine($"Hex: {number:X}"); // 输出 "FF"
Console.WriteLine($"Binary: {Convert.ToString(number, 2)}"); // 输出 "11111111"
```

## 11. 高级主题：元组 (Tuple)

C# 7.0 引入了元组，允许将多个值组合为一个轻量级的数据结构。

**示例代码**：

```csharp
(int, string) person = (25, "Alice");
Console.WriteLine($"Age: {person.Item1}, Name: {person.Item2}");

// 命名元组
var namedTuple = (Age: 25, Name: "Alice");
Console.WriteLine($"Age: {namedTuple.Age}, Name: {namedTuple.Name}");
```

## 12. 实践建议

1. 选择合适的数据类型：
   - 使用 `int` 或 `double` 进行常规计算。
   - 使用 `decimal` 进行金融计算。
   - 使用 `string` 处理文本。
2. **避免不必要的装箱/拆箱**：尽量使用泛型（如 `List<T>`）代替 `ArrayList`。
3. **善用可空类型**：在需要表示“无值”时使用 `int?` 等。
4. **小心类型转换**：使用 `TryParse` 避免异常。
5. **动态类型谨慎使用**：仅在必要时使用 `dynamic`，如与 COM 对象交互。
6. 字面量和进制转换：
   - 使用 `0x` 和 `0b` 前缀简化十六进制和二进制输入。
   - 使用 `Convert.ToString` 和 `Convert.ToInt32` 进行进制转换。

## 13. 示例程序

以下是一个综合示例，展示多种数据类型、字面量和进制转换的用法：

```csharp
using System;

class Program
{
    static void Main()
    {
        // 基本数据类型
        int age = 25;
        double height = 1.75;
        string name = "Alice";
        bool isStudent = true;

        // 可空类型
        int? score = null;
        Console.WriteLine(score.HasValue ? score.Value : "No score");

        // 数组
        int[] grades = { 85, 90, 95 };
        Console.WriteLine($"Average grade: {grades.Average()}");

        // 元组
        var person = (Age: age, Name: name);
        Console.WriteLine($"Person: {person.Name}, {person.Age} years old");

        // 类型转换
        string input = "123";
        if (int.TryParse(input, out int number))
        {
            Console.WriteLine($"Parsed number: {number}");
        }

        // 动态类型
        dynamic value = 42;
        Console.WriteLine(value);
        value = "Dynamic string";
        Console.WriteLine(value);

        // 字面量和进制转换
        int dec = 255;
        int hex = 0xFF; // 十六进制 FF = 255
        int bin = 0b11111111; // 二进制 11111111 = 255
        Console.WriteLine($"Decimal: {dec}, Hex: {hex:X}, Binary: {Convert.ToString(bin, 2)}");
        string hexString = "FF";
        int parsedHex = Convert.ToInt32(hexString, 16);
        Console.WriteLine($"Parsed hex {hexString} to decimal: {parsedHex}");
    }
}
```



---

> 作者: [hao](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/basic/c6a99510-88ae-4d53-9fbb-a175cc2b0a07/  

