# WPF入门



## WPF概述

Windows Presentation Foundation 桌面指南（WPF），这是一个独立于分辨率的 UI 框架，它使用基于矢量的呈现引擎，旨在利用现代图形硬件。 WPF 提供了一组全面的应用程序开发功能，其中包括可扩展应用程序标记语言（XAML）、控件、数据绑定、布局、2D 和 3D 图形、动画、样式、模板、文档、媒体、文本和版式。 WPF 是 .NET 的一部分，因此可以生成包含 .NET API 其他元素的应用程序。

WPF 有两个实现：

1. **.NET** 版本（本指南）：
2. **.NET Framework 4** 版本：

尽管.Net是一种跨平台技术，但是WPF仅能在Windows上运行。

## 什么是XAML？

XAML 是基于 XML 的标记语言，以声明方式实现应用程序的外观。 通常使用它来定义窗口、对话框、页面和用户控件，并用控件、形状和图形填充它们。

以下示例使用 XAML 实现包含单个按钮的窗口的外观：

```xaml
<Window
    xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
    Title="Window with button"
    Width="250" Height="100">

  <!-- Add button to window -->
  <Button Name="button">Click Me!</Button>

</Window>
```

具体而言，此 XAML 使用 Window 和 Button 元素定义窗口和按钮。 每个元素都配置了属性，例如 Window 元素的 Title 属性来指定窗口的标题栏文本。 在运行时，WPF 将标记中定义的元素和属性转换为 WPF 类的实例。 例如，Window 元素转换为 Window 类的实例，该类 Title 属性是 Title 特性的值。

下图显示了上一示例中 XAML 定义的用户界面（UI）：

![包含按钮的窗口](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/markup-window-button.png)

## WPF的项目结构

创建WPF应用程序，VS Studio新建项目，项目模板选择`WPF应用程序`。

![image-20250802142651992](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250802142651992.png)

这里的项目名称是`01WPF入门`，解决方案名称是`WPF框架学习`，创建后的项目结构如下：

![image-20250802142931386](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250802142931386.png)

双击解决方案名称

![image-20250802143717257](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250802143717257.png)

```
<Project Sdk="Microsoft.NET.Sdk">

  <PropertyGroup>
    <OutputType>WinExe</OutputType>	输出文件类型，widows可执行文件
    <TargetFramework>net8.0-windows</TargetFramework>	目标框架
    <RootNamespace>_01WPF入门</RootNamespace> 根命名空间
    <Nullable>enable</Nullable>
    <ImplicitUsings>enable</ImplicitUsings>
    <UseWPF>true</UseWPF>
  </PropertyGroup>

</Project>
```



---

> 作者: [hobby](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/wpf/892c051/  

