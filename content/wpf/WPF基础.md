---
title: "WPF基础"
date: 2025-10-12T16:32:43+0800
slug: "b33113a1-613e-499c-b440-d349011391a3"
draft: false
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
  - draft
categories:
  - draft
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
type: posts
---

## WPF介绍

WPF（Windows Presentation Foundation）是微软推出的一种用于构建桌面应用程序的图形子系统。它提供了一种统一的编程模型，用于创建具有丰富用户界面的应用程序。WPF基于XAML（eXtensible Application Markup Language）语言，使得开发者可以通过声明式的方式定义用户界面，同时支持数据绑定、动画、样式和模板等功能。

WPF的主要特点包括：

1. **矢量图形**：WPF使用矢量图形渲染界面，使得应用程序在不同分辨率下都能保持清晰。
2. **数据绑定**：WPF支持强大的数据绑定功能，可以轻松地将UI元素与数据源连接起来，实现动态更新。
3. **样式和模板**：WPF允许开发者定义样式和控制模板，使得界面可以高度定制化。
4. **动画和媒体支持**：WPF内置了对动画和多媒体的支持，可以创建丰富的用户体验。
5. **布局系统**：WPF提供了灵活的布局系统，可以根据窗口大小自动调整UI元素的位置和大小。    

WPF广泛应用于各种桌面应用程序的开发，特别适合需要复杂用户界面的应用，如图形设计工具、数据可视化应用等。通过WPF，开发者可以创建现代化、响应迅速且具有吸引力的桌面应用程序。

### WPF的基本概念

1. **XAML（eXtensible Application Markup Language）**：WPF使用XAML作为主要的用户界面定义语言。XAML是一种基于XML的标记语言，允许开发者以声明式的方式定义UI元素、布局和行为。通过XAML，开发者可以清晰地描述界面的结构和外观。
2. **控件（Controls）**：WPF提供了丰富的内置控件，如按钮、文本框、列表框等，这些控件可以直接在XAML中使用。控件是用户界面的基本构建块，开发者可以通过设置属性和事件来定制它们的行为和外观。
3. **布局（Layout）**：WPF提供了多种布局容器，如Grid、StackPanel、DockPanel等，用于组织和排列UI元素。布局容器可以根据窗口大小自动调整子元素的位置和大小，确保界面在不同分辨率下都能良好显示。
4. **数据绑定（Data Binding）**：WPF支持强大的数据绑定功能，允许UI元素与数据源（如对象、集合等）进行连接。通过数据绑定，UI可以自动反映数据的变化，实现动态更新。
5. **样式和模板（Styles and Templates）**：WPF允许开发者定义样式和控制模板，以实现界面的高度定制化。样式可以统一设置控件的外观，而模板则可以完全改变控件的结构和行为。
6. **事件（Events）**：WPF使用事件驱动模型，允许开发者响应用户交互（如点击、输入等）。WPF支持路由事件机制，可以在控件树中传播事件，方便事件处理。
7. **命令（Commands）**：WPF引入了命令概念，允许将用户操作与应用程序逻辑解耦。命令可以绑定到控件上，实现统一的操作处理。
8. **动画（Animation）**：WPF内置了对动画的支持，允许开发者创建丰富的用户体验。通过动画，可以实现控件的平滑过渡、移动、缩放等效果。   
9. **资源（Resources）**：WPF允许开发者定义和使用资源，如颜色、样式、模板等。资源可以在应用程序的不同部分共享，促进代码的重用和维护。
10. **依赖属性（Dependency Properties）**：WPF引入了依赖属性机制，允许属性值通过数据绑定、样式等方式动态变化。依赖属性支持属性值的继承和变化通知，增强了属性的灵活性。
11. **视觉树（Visual Tree）和逻辑树（Logical Tree）**：WPF中的UI元素组织成两种树结构。视觉树表示实际渲染的元素，而逻辑树表示元素的逻辑关系。理解这两种树结构有助于处理事件和布局。
12. **命令绑定（Command Binding）**：WPF允许将命令与命令处理程序绑定，简化了用户操作的处理。通过命令绑定，可以实现统一的操作逻辑。
13. **多线程支持（Multithreading Support）**：WPF支持多线程编程，允许在后台线程中执行耗时操作，同时保持UI的响应性。WPF提供了Dispatcher机制，用于在UI线程上执行操作。
14. **3D图形支持（3D Graphics Support）**：WPF内置了对3D图形的支持，允许开发者创建三维用户界面和图形效果。通过3D图形，可以实现更丰富的视觉体验。
15. **文档和打印支持（Document and Printing Support）**：WPF提供了对文档和打印的支持，允许开发者创建和打印复杂的文档内容。WPF支持固定文档和流文档两种文档类型。
16. **国际化和本地化（Internationalization and Localization）**：WPF支持国际化和本地化，  允许开发者创建多语言应用程序。通过资源文件，可以轻松地管理不同语言的文本和资源。
17. **触摸和手势支持（Touch and Gesture Support）**：WPF支持触摸屏和手势操作，允许开发者创建适应现代硬件的应用程序。通过触摸和手势，可以实现更自然的用户交互。

### WPF的项目结构

一个典型的WPF项目结构通常包括以下主要部分：

1. **App.xaml**：这是应用程序的入口文件，定义了应用程序的资源和启动逻辑。App.xaml.cs文件包含应用程序的代码逻辑。
2. **MainWindow.xaml**：这是应用程序的主窗口文件，定义了主界面的布局和控件。MainWindow.xaml.cs文件包含主窗口的代码逻辑。
3. **Views文件夹**：用于存放应用程序的其他窗口和用户控件的XAML文件及其对应的代码文件。
4. **Models文件夹**：用于存放应用程序的数据模型类，定义应用程序的数据结构和业务逻辑。
5. **ViewModels文件夹**：用于存放视图模型类，负责连接视图和模型，实现数据绑定和命令处理。
6. **Resources文件夹**：用于存放应用程序的资源文件，如样式、模板、图像等。
7. **Converters文件夹**：用于存放值转换器类，负责在数据绑定过程中转换数据格式。
8. **Services文件夹**：用于存放应用程序的服务类，负责处理数据访问、网络请求等功能。
9. **Properties文件夹**：包含应用程序的属性设置文件，如AssemblyInfo.cs和Settings.settings。
10. **Packages文件夹**：用于存放通过NuGet安装的第三方库和依赖项。
11. **App.config**：应用程序的配置文件，用于存储应用程序的设置和配置信息。
12. **项目文件（.csproj）**：包含项目的配置信息，如引用的库、编译选项等。 
13. **其他文件**：根据项目需求，可能还会包含其他文件和文件夹，如文档、测试代码等。

这个结构可以根据具体项目的需求进行调整和扩展，但通常会包含上述主要部分，以便组织和管理WPF应用程序的代码和资源。

## XAML基础

XAML（eXtensible Application Markup Language）是一种基于XML的标记语言，用于定义WPF应用程序的用户界面。通过XAML，开发者可以以声明式的方式描述UI元素、布局和行为，使得界面的结构和外观更加清晰和易于维护。

### XAML的基本语法
1. **元素（Elements）**：XAML中的每个UI组件都表示为一个元素。例如，按钮表示为<Button>元素，文本框表示为<TextBox>元素。

```xml
<Button Content="Click Me" Width="100" Height="30"/>
<TextBox Text="Enter text here" Width="200"/>
```

2. **属性（Attributes）**：元素可以通过属性来设置其外观和行为。例如，设置按钮的内容、宽度和高度。
```xml
<Button Content="Click Me" Width="100" Height="30"/>
```

3. **嵌套（Nesting）**：XAML允许元素嵌套在其他元素内部，以表示层次结构。例如，将按钮放在StackPanel中。

```xml
<StackPanel>
    <Button Content="Button 1" Width="100" Height="30"/>
    <Button Content="Button 2" Width="100" Height="30"/>
</StackPanel>
```
4. **命名空间（Namespaces）**：XAML使用XML命名空间来区分不同的元素和属性。通常，WPF元素使用默认命名空间。

```xml
xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
```

5. **注释（Comments）**：XAML支持注释，可以使用`<!-- -->`来添加注释。

```xml
<!-- This is a comment -->
<Button Content="Click Me" Width="100" Height="30"/>
```
6. **事件（Events）**：可以在XAML中为控件绑定事件处理程序。例如，为按钮的Click事件绑定一个方法。

```xml
<Button Content="Click Me" Click="Button_Click"/>
```

7. **资源（Resources）**：XAML允许定义和使用资源，如样式、颜色等。资源可以在应用程序的不同部分共享。

```xml
<Window.Resources>
    <SolidColorBrush x:Key="MyBrush" Color="LightBlue"/>
</Window.Resources>
<Button Background="{StaticResource MyBrush}" Content="Click Me"/>
```

8. **数据绑定（Data Binding）**：XAML支持数据绑定，可以将UI元素的属性绑定到数据源。

```xml
<TextBox Text="{Binding UserName}"/>  
```

9. **布局（Layout）**：XAML提供了多种布局容器，如Grid、StackPanel、DockPanel等，用于组织和排列UI元素。

```xml
<Grid>
    <Button Content="Button 1" Grid.Row="0" Grid.Column="0"/>
    <Button Content="Button 2" Grid.Row="0" Grid.Column="1"/>
</Grid>
```

10. **样式和模板（Styles and Templates）**：XAML允许定义样式和控制模板，以实现界面的高度定制化。

```xml
<Window.Resources>
    <Style TargetType="Button">
        <Setter Property="Background" Value="LightGreen"/>
        <Setter Property="FontSize" Value="16"/>
    </Style>
</Window.Resources>
<Button Content="Styled Button"/>
```

11. **命令（Commands）**：可以在XAML中绑定命令，以实现用户操作的处理。

```xml
<Button Content="Save" Command="{Binding SaveCommand}"/>
```

12. **转换器（Converters）**：XAML支持值转换器，可以在数据绑定过程中转换数据格式。

```xml
<TextBlock Text="{Binding IsEnabled, Converter={StaticResource BoolToTextConverter}}"/>
```

13. **触发器（Triggers）**：XAML允许使用触发器来响应属性变化或事件，从而动态改变UI元素的外观。

```xml
<Style TargetType="Button">
    <Style.Triggers>
        <Trigger Property="IsMouseOver" Value="True">
            <Setter Property="Background" Value="Yellow"/>
        </Trigger>
    </Style.Triggers>
</Style>
<Button Content="Hover Me"/>
```
14. **布局属性（Layout Properties）**：XAML中的布局容器提供了多种属性来控制子元素的排列方式，如Margin、Padding、HorizontalAlignment、VerticalAlignment等。

```xml
<StackPanel Margin="10" Padding="5" HorizontalAlignment="Center" VerticalAlignment="Top">
    <Button Content="Button 1" Margin="5"/>
    <Button Content="Button 2" Margin="5"/>
</StackPanel>
```

15. **命名元素（Naming Elements）**：可以为XAML中的元素指定名称，以便在代码中引用它们。使用`x:Name`属性来命名元素。

```xml
<Button x:Name="MyButton" Content="Click Me" Click="MyButton_Click"/>
```

16. **布局容器（Layout Containers）**：XAML提供了多种布局容器，用于组织和排列UI元素。常见的布局容器包括：

- **Grid**：用于创建网格布局，可以通过行和列来组织子元素。

```xml
<Grid>
    <Grid.RowDefinitions>
        <RowDefinition Height="Auto"/>
        <RowDefinition Height="*"/>
    </Grid.RowDefinitions>
    <Grid.ColumnDefinitions>
        <ColumnDefinition Width="*"/>
        <ColumnDefinition Width="*"/>
    </Grid.ColumnDefinitions>   
    <Button Content="Button 1" Grid.Row="0" Grid.Column="0"/>
    <Button Content="Button 2" Grid.Row="0" Grid.Column="1"/>
    <TextBox Grid.Row="1" Grid.ColumnSpan="2" Text="Enter text here"/>
</Grid>
```

- **StackPanel**：用于垂直或水平堆叠子元素。

```xml
<StackPanel Orientation="Vertical">
    <Button Content="Button 1" Width="100" Height="30"/>
    <Button Content="Button 2" Width="100" Height="30"/>
</StackPanel>
```

- **DockPanel**：用于将子元素停靠在容器的边缘。

```xml
<DockPanel>
    <Button Content="Top" DockPanel.Dock="Top"/>
    <Button Content="Bottom" DockPanel.Dock="Bottom"/>
    <Button Content="Left" DockPanel.Dock="Left"/>
    <Button Content="Right" DockPanel.Dock="Right"/>
    <TextBox Text="Center" />
</DockPanel>
```
- **Canvas**：用于绝对定位子元素。

```xml
<Canvas>
    <Button Content="Button 1" Canvas.Left="50" Canvas.Top="30"/>
    <Button Content="Button 2" Canvas.Left="150" Canvas.Top="80"/>
</Canvas>
```

- **WrapPanel**：用于在水平或垂直方向上自动换行排列子元素。

```xml
  <WrapPanel Orientation="Horizontal">
    <Button Content="Button 1" Width="100" Height="30"/>
    <Button Content="Button 2" Width="100" Height="30"/>
    <Button Content="Button 3" Width="100" Height="30"/>
    <Button Content="Button 4" Width="100" Height="30"/>
</WrapPanel>
``` 
- **UniformGrid**：用于创建均匀分布的网格布局。

```xml
<UniformGrid Rows="2" Columns="2">
    <Button Content="Button 1"/>
    <Button Content="Button 2"/>
    <Button Content="Button 3"/>
    <Button Content="Button 4"/>
</UniformGrid>
```

- **ScrollViewer**：用于在内容超出可视区域时提供滚动功能。

```xml
<ScrollViewer Width="200" Height="100">
    <StackPanel>
        <Button Content="Button 1" Width="100" Height="30"/>
        <Button Content="Button 2" Width="100" Height="30"/>
        <Button Content="Button 3" Width="100" Height="30"/>
        <Button Content="Button 4" Width="100" Height="30"/>
        <Button Content="Button 5" Width="100" Height="30"/>
    </StackPanel>
</ScrollViewer>
``` 


### XAML与C#的结合

XAML与C#代码通常结合使用，以实现WPF应用程序的功能和交互。XAML负责定义用户界面的结构和外观，而C#代码则处理应用程序的逻辑和行为。以下是XAML与C#结合使用的一些常见方式：
1. **事件处理**：在XAML中为控件绑定事件处理程序，然后在C#代码中实现这些处理程序。例如，为按钮的Click事件绑定一个方法。
XAML:
```xml
<Button Content="Click Me" Click="Button_Click"/>
```
C#:
```csharp
private void Button_Click(object sender, RoutedEventArgs e)
{
    MessageBox.Show("Button clicked!");
}
```

2. **数据绑定**：在XAML中使用数据绑定将UI元素的属性绑定到C#中的数据源。通过实现INotifyPropertyChanged接口，可以实现数据的动态更新。
XAML:
```xml
<TextBox Text="{Binding UserName}"/>
```
C#:
```csharp
public class ViewModel : INotifyPropertyChanged
{
    private string userName;
    public string UserName
    {
        get { return userName; }
        set
        {
            userName = value;
            OnPropertyChanged("UserName");
        }
    }
    public event PropertyChangedEventHandler PropertyChanged;
    protected void OnPropertyChanged(string name)
    {
        PropertyChanged?.Invoke(this, new PropertyChangedEventArgs(name));
    }
}
```
3. **命令绑定**：在XAML中绑定命令到控件，然后在C#代码中实现命令逻辑。通过实现ICommand接口，可以创建自定义命令。
XAML:
```xml
<Button Content="Save" Command="{Binding SaveCommand}"/>
```
C#:
```csharp
public class ViewModel
{
    public ICommand SaveCommand { get; private set; }
    public ViewModel()
    {
        SaveCommand = new RelayCommand(Save);
    }
    private void Save()
    {
        // Save logic here
    }
}
```

4. **访问XAML元素**：在C#代码中通过元素的名称访问XAML定义的控件，以便操作它们的属性和方法。
XAML:
```xml
<Button x:Name="MyButton" Content="Click Me" Click="MyButton_Click"/>
```
C#:
```csharp
private void MyButton_Click(object sender, RoutedEventArgs e)
{
    MyButton.Content = "Clicked!";
}
```

5. **资源管理**：在XAML中定义资源（如样式、颜色等），然后在C#代码中访问和使用这些资源。
XAML:
```xml
<Window.Resources>
    <SolidColorBrush x:Key="MyBrush" Color="LightBlue"/>
</Window.Resources>
<Button Background="{StaticResource MyBrush}" Content="Click Me"/>
```
C#:
```csharp
var brush = (Brush)FindResource("MyBrush");
MyButton.Background = brush;
```

6. **导航和页面管理**：在WPF应用程序中，可以使用Frame和Page控件实现页面导航。在XAML中定义Frame和Page，然后在C#代码中控制导航逻辑。
XAML:
```xml
<Frame x:Name="MainFrame"/>
```
C#:
```csharp
MainFrame.Navigate(new Uri("Page1.xaml", UriKind.Relative));
```
7. **动画和视觉效果**：在XAML中定义动画和视觉效果，然后在C#代码中触发这些动画。

XAML:
```xml
<Button x:Name="MyButton" Content="Animate Me">
    <Button.Triggers>
        <EventTrigger RoutedEvent="Button.Click">
            <BeginStoryboard>
                <Storyboard>
                    <DoubleAnimation Storyboard.TargetProperty="Width" To="200" Duration="0:  0:0.5"/>
                </Storyboard>
            </BeginStoryboard>
        </EventTrigger>
    </Button.Triggers>
</Button>
```
C#:
```csharp
// Animation is triggered by the button click event defined in XAML
``` 

8. **多线程操作**：在C#代码中使用Dispatcher来在UI线程上执行操作，以确保UI的响应性。
```csharp
Application.Current.Dispatcher.Invoke(() =>
{
    MyButton.Content = "Updated from background thread";
});
```   

通过以上方式，XAML与C#代码紧密结合，共同实现WPF应用程序的功能和交互。开发者可以利用XAML的声明式语法定义界面，同时通过C#代码实现复杂的逻辑和行为，从而创建现代化、响应迅速且具有吸引力的桌面应用程序。





