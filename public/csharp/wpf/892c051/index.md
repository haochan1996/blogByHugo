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

创建WPF应用程序。

<img src="https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250803151842443.png" alt="image-20250803151842443" style="zoom: 33%;" />

**App.xaml 与 App.xaml.cs**

- 入口点：定义应用程序启动逻辑，通过StartupUri指定初始窗口（如StartupUri="MainWindow.xaml"）。
- 全局资源：在<Application.Resources>中定义样式、数据模板等资源，供整个应用复用。
- 生命周期事件：在App.xaml.cs中重写OnStartup()初始化服务，OnExit()执行清理逻辑，DispatcherUnhandledException捕获全局异常。

**MainWindow.xaml 与 MainWindow.xaml.cs**

- 主窗口UI：XAML文件定义布局和控件（如按钮、文本框）。
- 事件处理：代码文件（xaml.cs）实现交互逻辑（如按钮点击事件）。

## App.xaml

<img src="https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250803162107972.png" alt="image-20250803162107972" style="zoom: 50%;" />

App.xaml 是 WPF（Windows Presentation Foundation）应用程序的核心文件，承担着**应用程序入口点、全局资源配置中心**和**生命周期事件处理器**等关键角色。其核心作用可归纳为以下五类：

### 应用程序入口与启动控制

- **启动窗口定义**：通过`StartupUri`属性指定初始窗口（如 `StartupUri="MainWindow.xaml"`），简化启动流程。

- **自定义启动逻辑**：删除`StartupUri`后，可在`App.xaml.cs`中重写`OnStartup` 方法，实现动态初始化（如窗口预配置、参数解析、依赖注入容器初始化等操作）：

  ```c#
  protected override void OnStartup(StartupEventArgs e) 
  {
      base.OnStartup(e);
      // 示例：解析启动参数
      if (e.Args.Contains("debug")) 
          DebugMode.Enable();
      // 动态创建主窗口
      var mainWindow = new MainWindow();
      mainWindow.Title = "Custom Title";
      mainWindow.Show();
  }
  ```

### 全局资源集中管理

- **统一样式与模板**：在`<Application.Resources>`中定义应用级资源（样式、画笔、数据模板等），确保UI一致性：

  ```xaml
  <Application.Resources>
      <Style TargetType="Button">
          <Setter Property="Foreground" Value="Red"/>
          <Setter Property="FontSize" Value="14"/>
      </Style>
      <ResourceDictionary>
          <ResourceDictionary.MergedDictionaries>
              <ResourceDictionary Source="Styles/Icons.xaml"/>
          </ResourceDictionary.MergedDictionaries>
      </ResourceDictionary>
  </Application.Resources>
  ```

- **多语言支持**：通过合并不同语言的资源字典（如 `en-us.xaml`, `zh-cn.xaml`），实现动态切换界面语言。

### 应用程序生命周期管理

- **关键事件处理**：

  - `OnStartup`：初始化全局状态（如数据库连接、配置加载）。

  - `OnExit`：执行清理任务（如保存用户数据、释放资源）。

  - `OnSessionEnding`：拦截系统关机/注销事件，提示未保存数据：

    ```c#
    protected override void OnSessionEnding(SessionEndingCancelEventArgs e) 
    {
        if (HasUnsavedData) {
            e.Cancel = true;
            MessageBox.Show("数据未保存！");
        }
    }
    ```

- **未处理异常捕获**：订阅 `DispatcherUnhandledException` 事件，防止崩溃并记录错误。

### 统一 UI 行为与线程管理

- **全局样式生效范围**：在`App.xaml`中定义的隐式样式（无`x:Key`）自动应用于所有匹配控件，无需显式引用。
- **UI 线程安全**：通过`Application.Current.Dispatcher.Invoke`确保跨线程操作 UI 的安全性（但`OnStartup`中无需调用，因已在主线程）。

### 高级配置与扩展性

- 关闭模式控制：`ShutdownMode`属性决定应用退出时机：

  - `OnLastWindowClose`（默认）：所有窗口关闭后退出。
  - `OnMainWindowClose`：主窗口关闭即退出。
  - `OnExplicitShutdown`：需手动调用 `Application.Current.Shutdown()`。

- 依赖注入集成：在`OnStartup`中初始化容器（如 Unity、Autofac），注册全局服务：

  ```c#
  public static IUnityContainer Container;
  protected override void OnStartup(StartupEventArgs e) {
      Container = new UnityContainer();
      Container.RegisterType<IDataService, DataService>();
      base.OnStartup(e);
  }
  ```


## Application的生命周期

WPF（Windows Presentation Foundation）应用程序的生命周期由 Application 类管理，涵盖从启动到关闭的全过程，开发者可通过重写方法或订阅事件介入关键节点。以下是核心阶段及关键行为的解析：

### 启动阶段（Startup）
触发时机：应用程序入口点（Main 方法）调用 `Application.Run()` 后，主窗口显示前。

核心方法：`OnStartup(StartupEventArgs e)`：  用于初始化全局资源（如数据库连接、配置加载）、解析命令行参数（e.Args），或动态创建启动窗口（替代 StartupUri）。

```C#
protected override void OnStartup(StartupEventArgs e) 
{
    base.OnStartup(e);
    MainWindow = new CustomWindow(); // 动态创建主窗口
    MainWindow.Show();
}
```

启动画面（Splash Screen）：通过添加图像文件并设置生成操作为 SplashScreen，实现启动瞬间显示初始界面。

原生SplashScreen实现，适用于静态图片场景，性能最优，由系统级 API 支持。

**实现步骤：**

1. **添加图片资源**
   - 在项目中添加图片（支持 PNG、JPEG、BMP 等格式）。
   - 属性设置：右键图片 → 生成操作 选 SplashScreen（VS 自动生成代码）。
2. **代码控制显示逻辑**（可选高级配置）

```c#
protected override void OnStartup(StartupEventArgs e) 
{
    // 创建 SplashScreen 实例（图片路径需匹配资源名）
    var splash = new SplashScreen("SplashImage.png");  
    
    // 非自动关闭 + 置顶显示
    splash.Show(false, true);  
    
    // 设置超时关闭（防止主窗口卡死导致 Splash 滞留）
    var timer = new Timer(_ => 
    {
        Dispatcher.Invoke(() => splash.Close(TimeSpan.Zero));
    }, null, 15000, Timeout.Infinite);  // 最长显示 15 秒[4](@ref)
    
    // 初始化主窗口
    var mainWindow = new MainWindow();
    mainWindow.Loaded += (s, args) => splash.Close(TimeSpan.FromSeconds(0.5)); // 主窗口加载完成后淡出
    mainWindow.Show();
}
```

### 运行阶段（Active/Inactive）
激活与失焦：

- `OnActivated(EventArgs e)`：应用或任意窗口获焦时触发（如从后台切换至前台），适用于刷新实时数据。
- `OnDeactivated(EventArgs e)`：应用所有窗口失焦时触发，可暂停非关键任务或保存临时状态。

全局事件：

- `SessionEnding`：系统关机/注销时触发，通过`e.Cancel=true`可阻止关闭（例如提示未保存数据）。

```c#
protected override void OnSessionEnding(SessionEndingCancelEventArgs e) 
{
    if (HasUnsavedData) e.Cancel = true; // 取消系统关闭
}
```

### 关闭阶段（Shutdown）
关闭模式（ShutdownMode）：通过枚举值控制退出条件，可在App.xaml或代码中设置：

|       **模式**       |                    **触发条件**                    |
| :------------------: | :------------------------------------------------: |
| `OnLastWindowClose`  |           所有窗口关闭后退出（**默认**）           |
| `OnMainWindowClose`  |    主窗口（`Application.MainWindow`）关闭即退出    |
| `OnExplicitShutdown` | 需显式调用 `Application.Current.Shutdown()` 才退出 |

结束处理：`OnExit(ExitEventArgs e)`：应用退出前最后阶段，释放资源（如关闭数据库连接）、保存用户设置或记录日志。

```C#
protected override void OnExit(ExitEventArgs e) 
{
    SaveUserSettings();
    base.OnExit(e);
}
```
强制退出：`Environment.Exit(0)`立即终止进程（不执行清理），慎用。

### 全局异常处理
未捕获异常的捕获方式：

- `DispatcherUnhandledException`：捕获 UI 线程异常（如按钮事件中的错误），通过`e.Handled = true`标记为已处理。

- `AppDomain.CurrentDomain.UnhandledException`:捕获所有非 UI 线程异常（如后台线程崩溃），但无法阻止进程终止。

- `TaskScheduler.UnobservedTaskException`：专用于 Task 异步任务的异常（延迟至垃圾回收时触发）。 

```C#
public App() {
    DispatcherUnhandledException += (s, e) => LogError("UI线程异常", e.Exception);
    AppDomain.CurrentDomain.UnhandledException += (s, e) => LogError("非UI线程异常", e.ExceptionObject as Exception);
}
```

生命周期，可高效协调资源管理、异常恢复与用户体验，是构建健壮 WPF 应用的核心基础。

## 窗体的生命周期

WPF（Windows Presentation Foundation）窗体的生命周期涵盖了从创建到销毁的所有阶段。理解这些生命周期事件有助于更好地管理和控制应用程序的行为，特别是在资源管理、状态维护和用户交互方面。以下是WPF窗体的主要生命周期事件及其详细解释：

1. **构造函数**: 当窗体实例被创建时，构造函数会被调用。
2. **Loaded**: 在窗体首次显示在屏幕上之前触发。
3. **Initialized**: 当元素已完全初始化并且属性值应用后触发。
4. **Activated**: 窗体获得焦点时触发。
5. **Deactivated**: 窗体失去焦点时触发。
6. **StateChanged**: 窗体状态（如最大化、最小化、还原）改变时触发。
7. **Closing**: 窗体关闭前触发，允许取消关闭操作。
8. **Closed**: 窗体关闭后触发。

```c#
using System;
using System.Windows;

namespace WpfAppLifecycle
{
    public partial class MainWindow : Window
    {
        public MainWindow()
        {
            InitializeComponent();
            Loaded += OnMainWindowLoaded;
            Initialized += OnMainWindowInitialized;
            Activated += OnMainWindowActivated;
            Deactivated += OnMainWindowDeactivated;
            StateChanged += OnMainWindowStateChanged;
            Closing += OnMainWindowClosing;
            Closed += OnMainWindowClosed;
        }

        private void OnMainWindowLoaded(object sender, RoutedEventArgs e)
        {
            Console.WriteLine("MainWindow Loaded");
        }

        private void OnMainWindowInitialized(object sender, EventArgs e)
        {
            Console.WriteLine("MainWindow Initialized");
        }

        private void OnMainWindowActivated(object sender, EventArgs e)
        {
            Console.WriteLine("MainWindow Activated");
        }

        private void OnMainWindowDeactivated(object sender, EventArgs e)
        {
            Console.WriteLine("MainWindow Deactivated");
        }

        private void OnMainWindowStateChanged(object sender, EventArgs e)
        {
            Console.WriteLine($"MainWindow State Changed: {WindowState}");
        }

        private void OnMainWindowClosing(object sender, System.ComponentModel.CancelEventArgs e)
        {
            Console.WriteLine("MainWindow Closing");
            // 可以在这里取消关闭操作，例如：e.Cancel = true;
        }

        private void OnMainWindowClosed(object sender, EventArgs e)
        {
            Console.WriteLine("MainWindow Closed");
        }
    }
}
```

## 常用属性

在WPF（Windows Presentation Foundation）中，`Window` 类提供了许多常用的属性来控制窗口的外观和行为。以下是一些最常用且重要的 `Window` 属性：

### **Title**

- 描述: 设置窗口的标题栏文本。
- 示例: `<Window Title="My Application" />`

### **Width** 和 **Height**

- 描述: 设置窗口的宽度和高度。
- 示例: `<Window Width="500" Height="300" />`

### **MinWidth**, **MaxWidth**, **MinHeight**, **MaxHeight**

- 描述: 设置窗口的最小和最大宽度及高度。
- 示例: `<Window MinWidth="400" MaxWidth="800" MinHeight="200" MaxHeight="600" />`

### **ResizeMode**
   - 描述: 控制用户是否可以调整窗口大小。
   - 可能的值:
     - `NoResize`: 用户不能调整窗口大小。
     - `CanMinimize`: 用户只能最小化窗口。
     - `CanResize`: 用户可以调整窗口大小（默认值）。
     - `CanResizeWithGrip`: 用户可以通过拖动右下角的调整手柄来调整窗口大小。
   - 示例: `<Window ResizeMode="CanResizeWithGrip" />`

### **WindowState**

   - 描述: 设置窗口的初始状态。
   - 可能的值:
     - `Normal`: 正常显示窗口。
     - `Maximized`: 最大化窗口。
     - `Minimized`: 最小化窗口。
   - 示例: `<Window WindowState="Maximized" />`

### **ShowInTaskbar**
   - 描述: 控制窗口是否显示在任务栏上。
   - 默认值: `true`
   - 示例: `<Window ShowInTaskbar="false" />`

### **Topmost**
   - 描述: 控制窗口是否始终位于其他窗口之上。
   - 默认值: `false`
   - 示例: `<Window Topmost="true" />`

### **Icon**
   - 描述: 设置窗口的任务栏图标和标题栏图标。
   - 示例: `<Window Icon="appicon.ico" />`

### **Background**
   - 描述: 设置窗口的背景颜色或图像。
   - 示例: `<Window Background="LightGray" />`

### **AllowsTransparency**
   - 描述: 允许窗口具有透明效果。
   - 默认值: `false`
   - 注意: 如果设置为 `true`，则必须将 `WindowStyle` 设置为 `None`。
   - 示例: `<Window AllowsTransparency="true" WindowStyle="None" Background="#80FFFFFF" />`

### **WindowStyle**
   - 描述: 设置窗口的样式。
   - 可能的值:
     - `SingleBorderWindow`: 单边框窗口，默认值。
     - `ThreeDBorderWindow`: 三维边框窗口。
     - `ToolWindow`: 工具窗口样式。
     - `None`: 无边框窗口。
   - 示例: `<Window WindowStyle="None" />`

### **Opacity**
   - 描述: 设置窗口的不透明度。
   - 范围: 0.0（完全透明）到 1.0（完全不透明）。
   - 示例: `<Window Opacity="0.8" />`

### **SnapsToDevicePixels**

 描述: 控制元素是否对齐到设备像素边界以减少模糊。
- 默认值: `false`
- 示例: `<Window SnapsToDevicePixels="true" />`

### **Focusable**
- 描述: 控制窗口是否可以接收焦点。
- 默认值: `true`
- 示例: `<Window Focusable="false" />`

### Visibility 

Visibility 属性用于控制控件的可见性

- `Visible`: 控件是可见的，并且占据空间。  
- `Hidden`: 控件不可见，但仍然占据空间。  
- `Collapsed`: 控件不可见，并且不占据任何空间。

下面是一个包含多个常用属性的 WPF 窗口示例：

```C#
<!-- MainWindow.xaml -->
<Window x:Class="WpfAppProperties.MainWindow"
        xmlns="http://schemas.microsoft.com/winfx/2006/xaml/presentation"
        xmlns:x="http://schemas.microsoft.com/winfx/2006/xaml"
        Title="My WPF Application"
        Width="500"
        Height="300"
        MinWidth="400"
        MaxWidth="800"
        MinHeight="200"
        MaxHeight="600"
        ResizeMode="CanResizeWithGrip"
        WindowState="Normal"
        ShowInTaskbar="true"
        Topmost="false"
        Icon="appicon.ico"
        Background="LightGray"
        AllowsTransparency="false"
        WindowStyle="SingleBorderWindow"
        Opacity="1.0"
        SnapsToDevicePixels="true"
        Focusable="true">
    <Grid>
        <TextBlock Text="Welcome to my WPF application!" FontSize="24" HorizontalAlignment="Center" VerticalAlignment="Center"/>
    </Grid>
</Window>
```

## 程序的退出方式



---

> 作者: [hobby](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/wpf/892c051/  

