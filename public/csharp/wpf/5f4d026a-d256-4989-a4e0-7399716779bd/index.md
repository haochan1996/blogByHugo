# WPF内容控件


## Control类

定义:表示用户界面(UI)元素的基类，这些元素使用 **ControlTemplate** 来定义其外观。

Control是许多控件的基类。比如最常见的按钮(Button)、单选(RadioButton)、复选(CheckBox)、文本框(TextBox)、 ListBox、DataGrid、日期控件等等。这些控件通常用于展示程序的数据或获取用户输入的数据，我们可以将这一类型的控件称为内容控件或数据控件，它们与前面的布局控件有一定的区别，布局控件更专注于界面，而内容控件更专注于数据(业务)。

Control类虽然可以实例化，但是在界面上是不会有任何显示的。只有那些继承了Control的子类(控件)才会在界面上显示，而且所呈现的样子各不相同，为什么会是这样呢?

因为Control类提供了一个控件模板(ControlTemplate)，而几乎所有的子类都对这个ControlTemplate进行了各自的实现，所以在呈现子类时，我们才会看到Button拥有Button的样子，TextBox拥有TextBox的样子。

```xaml
<Control>
    <Control.Template>
        <ControlTemplate TargetType="Control">
            <Border Width="200" Height="200" Background="Red" CornerRadius="100">
                <TextBlock Text="内容控件" HorizontalAlignment="Center" VerticalAlignment="Center" Foreground="White" FontSize="24"/>
            </Border>
        </ControlTemplate>
    </Control.Template>

</Control>
```

效果

<img src="https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250806120428254.png" alt="image-20250806120428254" style="zoom:33%;" />



### 控件属性

|         **属性名称**         |                           **说明**                           |
| :--------------------------: | :----------------------------------------------------------: |
|        **字体与文本**        |                                                              |
|         `FontFamily`         |        设置控件的字体系列（如 `Arial`、`微软雅黑`）。        |
|          `FontSize`          |             设置字体大小（单位：设备无关像素）。             |
|         `FontStyle`          |           设置字体样式（如 `Normal`、`Italic`）。            |
|         `FontWeight`         |            设置字体粗细（如 `Normal`、`Bold`）。             |
|        `FontStretch`         |       控制字体拉伸程度（如 `Condensed`、`Expanded`）。       |
|        **颜色与背景**        |                                                              |
|         `Foreground`         | 设置前景色（文本/图标颜色），类型为 `Brush`（支持纯色、渐变等）。 |
|         `Background`         |        设置背景色（控件内部填充色），类型为 `Brush`。        |
|        `BorderBrush`         |                设置边框颜色，类型为 `Brush`。                |
|        **布局与尺寸**        |                                                              |
|      `Width` / `Height`      |       显式设置控件的宽度和高度（支持 `Auto` 自适应）。       |
|   `MinWidth` / `MinHeight`   |          设置控件的最小宽度/高度（约束尺寸范围）。           |
|   `MaxWidth` / `MaxHeight`   |                  设置控件的最大宽度/高度。                   |
|           `Margin`           |            设置控件的外边距（与其他元素的距离）。            |
|          `Padding`           |            设置控件的内边距（内容与边框的距离）。            |
|         **内容对齐**         |                                                              |
| `HorizontalContentAlignment` |     设置内容的水平对齐方式（`Left`/`Center`/`Right`）。      |
|  `VerticalContentAlignment`  |     设置内容的垂直对齐方式（`Top`/`Center`/`Bottom`）。      |
|         **边框设置**         |                                                              |
|      `BorderThickness`       |    设置边框的厚度（如 `"1"` 或 `"1,2,3,4"` 分方向设置）。    |
|         **交互行为**         |                                                              |
|         `IsTabStop`          |      控制是否可通过 Tab 键导航到该控件（默认 `True`）。      |
|          `TabIndex`          |       设置 Tab 键导航时的顺序（数值越小优先级越高）。        |
|           `Cursor`           |       设置鼠标悬停时的光标样式（如 `Hand`、`Wait`）。        |
|        **模板与样式**        |                                                              |
|          `Template`          | 定义控件的视觉模板（通过 `ControlTemplate` 完全自定义外观）。 |
|           `Style`            |       引用样式资源，批量设置属性值（如全局按钮样式）。       |
|           **其他**           |                                                              |
|          `ToolTip`           |           设置鼠标悬停时的提示信息（支持富文本）。           |
|        `ContextMenu`         |                     设置右键上下文菜单。                     |
|         `Visibility`         |        控制可见性（`Visible`/`Collapsed`/`Hidden`）。        |
|          `Opacity`           |          设置透明度（0.0 完全透明 ~ 1.0 不透明）。           |

大部分的属性都比较好理解，这里着重介绍一下Template属性。属性 Template 是，ControlTemplate指定的外观 Control。如果要更改控件的外观但保留其功能，应考虑创建新 ControlTemplate 控件，而不是创建新类。如果把人比作是一个Control(控件)，那么”着装“就是Template(模板)。在大街上，我们会看到不同着装的人来来往往。

### 事件处理

|     **事件名称**      |                         **触发时机**                         |                       **常见用途**                        |
| :-------------------: | :----------------------------------------------------------: | :-------------------------------------------------------: |
|   **生命周期事件**    |                                                              |                                                           |
|     `Initialized`     |      控件初始化完成，属性已设置但尚未布局或渲染时触发。      |             初始化非依赖属性或执行早期配置。              |
|       `Loaded`        |       控件完成布局、数据绑定和渲染，可用于交互时触发。       |         加载数据、启动动画或执行依赖布局的逻辑。          |
|      `Unloaded`       |                  控件从视觉树中移除时触发。                  |          释放资源、取消订阅事件或清理后台任务。           |
|     **鼠标事件**      |                                                              |                                                           |
|     `MouseEnter`      |                 鼠标指针进入控件边界时触发。                 |            显示悬停提示（ToolTip）或高亮控件。            |
|     `MouseLeave`      |                 鼠标指针离开控件边界时触发。                 |               隐藏提示或恢复控件默认状态。                |
| `MouseLeftButtonDown` | 鼠标左键在控件上按下时触发（注意：`Button` 等控件默认抑制此事件）。 | 启动拖拽操作或自定义点击逻辑（需配合 `e.Handled` 处理）。 |
|  `MouseLeftButtonUp`  |                 鼠标左键在控件上释放时触发。                 |                 完成拖拽或确认点击操作。                  |
|      `MouseMove`      |                 鼠标在控件上移动时持续触发。                 |             实时追踪鼠标位置（如绘图工具）。              |
|  `MouseDoubleClick`   |          双击控件时触发（仅限 `Control` 派生类）。           |          快速编辑操作（如列表项双击打开详情）。           |
|     **键盘事件**      |                                                              |                                                           |
|       `KeyDown`       |              控件获得焦点且键盘按键按下时触发。              |            捕获快捷键（如 `Enter` 提交表单）。            |
|        `KeyUp`        |              控件获得焦点且键盘按键释放时触发。              |                响应按键释放后的状态更新。                 |
|     **焦点事件**      |                                                              |                                                           |
|      `GotFocus`       |                   控件获得逻辑焦点时触发。                   |                激活编辑模式或显示辅助UI。                 |
|      `LostFocus`      |                   控件失去逻辑焦点时触发。                   |               验证输入数据或保存编辑内容。                |
|     **拖拽事件**      |                                                              |                                                           |
|      `DragEnter`      |  拖拽对象进入控件边界时触发（需设置 `AllowDrop="True"`）。   |      检查拖拽数据类型并显示视觉反馈（如高亮边框）。       |
|      `DragOver`       |               拖拽对象在控件上移动时持续触发。               |          动态更新拖拽位置（如实时调整插入点）。           |
|        `Drop`         |                 拖拽对象在控件上释放时触发。                 |         处理拖拽数据（如文件导入或控件重定位）。          |

WPF 的 Control 事件基于路由事件模型，分为**冒泡路由（从子控件向父容器传递）**和**隧道路由（从父容器向子控件传递，以 Preview 前缀标识，如 PreviewMouseLeftButtonDown）**。实际应用中需注意：

- 事件抑制：部分控件（如Button）会标记`e.Handled=true`抑制底层事件（如MouseLeftButtonDown），此时需改用Preview事件或显式调用AddHandler。
- 拖拽必要条件：接收拖拽事件的控件必须设置`AllowDrop="True"`且Background非null（建议设为 Transparent）。
- 性能优化：高频事件（如 MouseMove）中避免复杂逻辑，或使用去抖机制（Debounce）减少处理频率。
- 命令与行为：为解耦 UI 与逻辑，建议将事件处理封装为 Behavior（如拖拽行为）或绑定到 ICommand 实现 MVVM 模式。

## ControlContent（内容控件）

`ContentControl`是WPF中用于承载任意类型内容的基类，其核心是通过`Content` 属性动态显示数据或UI元素。

### 核心特性

- 内容灵活性：Content 属性类型为 object，可接受字符串、图像、其他控件（如 Button）甚至复杂数据对象。
- 模板化支持：通过 ContentTemplate 属性自定义内容的视觉呈现方式（如将数据对象转换为富文本布局）。
- 数据绑定：支持将 Content 绑定到数据源（如数据库字段、XML 节点），实现动态更新。

### 自定义内容模板

将数据对象转换为复杂 UI（如员工卡片）

```xaml
<ContentControl Content="{Binding Employee}">  
    <ContentControl.ContentTemplate>  
        <DataTemplate>  
            <StackPanel>  
                <Image Source="{Binding Avatar}" Width="50"/>  
                <TextBlock Text="{Binding Name}" FontWeight="Bold"/>  
            </StackPanel>  
        </DataTemplate>  
    </ContentControl.ContentTemplate>  
</ContentControl>  
```

## Button(按钮控件)

Button继承自ButtonBase，是用户交互的核心控件，用于触发命令或事件。

### 核心属性详解

|      **属性**      |                   **说明**                    |                           **示例**                           |
| :----------------: | :-------------------------------------------: | :----------------------------------------------------------: |
|     `Content`      |         按钮显示内容（支持任意对象）          | `<Button Content="提交"/>` 或 `<Button> <Image Source="icon.png"/> </Button>` |
|     `Command`      |       绑定执行命令（如 `SaveCommand`）        |         `<Button Command="{Binding SaveCommand}"/>`          |
| `CommandParameter` |                 传递命令参数                  | `<Button Command="{Binding DeleteCommand}" CommandParameter="{Binding SelectedItem}"/>` |
|    `IsDefault`     |           设为默认按钮（回车触发）            |         `<Button Content="确定" IsDefault="True"/>`          |
|     `IsCancel`     |           设为取消按钮（Esc 触发）            |          `<Button Content="取消" IsCancel="True"/>`          |
|    `ClickMode`     | 控制点击响应时机（`Press`/`Release`/`Hover`） |                `<Button ClickMode="Press"/>`                 |

### 事件处理

#### 常用事件

- Click：点击触发（最常用）。
- MouseEnter/MouseLeave：鼠标悬停/离开时改变视觉状态。
- PreviewMouseDown/PreviewMouseUp：高级交互控制（如长按检测）。

#### 事件绑定

```xaml
<Button Click="Button_Click" MouseEnter="Button_MouseEnter"/>  
```

```c#
private void Button_Click(object sender, RoutedEventArgs e) 
{  
    MessageBox.Show("操作执行！");  
} 
```

### 样式与模板定制

#### 基础样式

通过 `Style` 统一设置外观

```xaml
<Style TargetType="Button">  
    <Setter Property="Background" Value="LightBlue"/>  
    <Setter Property="Foreground" Value="Black"/>  
    <Setter Property="Padding" Value="10,5"/>  
</Style>
```

#### 自定义模板（ControlTemplate）

重构视觉树实现个性化设计（如圆形按钮）

```xaml
<Button Content="确认">
    <Button.Template>
        <ControlTemplate TargetType="Button">
            <Grid Width="200" Height="200">
                <Ellipse Fill="Orange" Stroke="DarkBlue" StrokeThickness="2"/>
                <ContentPresenter  HorizontalAlignment="Center" VerticalAlignment="Center"/>
            </Grid>
        </ControlTemplate>
    </Button.Template>
</Button>
```

显示效果

<img src="https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250806223839614.png" alt="image-20250806223839614" style="zoom:33%;" />

#### 状态触发器（Triggers）

响应交互状态（悬停、禁用等）

```xaml
<Style.Triggers>  
    <Trigger Property="IsMouseOver" Value="True">  
        <Setter Property="Background" Value="DarkGray"/>  
    </Trigger>  
    <Trigger Property="IsEnabled" Value="False">  
        <Setter Property="Opacity" Value="0.5"/>  
    </Trigger>  
</Style.Triggers>
```

### 命令绑定与 MVVM

- 命令绑定示例：

  ```xml
  <Button Content="保存" Command="{Binding SaveCommand}"/>  
  ```

  ```csharp
  public ICommand SaveCommand => new RelayCommand(() => {  
      // 保存逻辑  
  }, canExecute: true); [1,5](@ref)  
  ```

- **优势**：分离业务逻辑与 UI，提升可测试性。

### 事件冒泡控制

使用`e.Handled = true`阻止事件向上传递

```c#
private void Button_PreviewMouseDown(object sender, MouseButtonEventArgs e) {  
    e.Handled = true; // 阻止父容器接收事件  
}
```



---

> 作者: [hao](https://github.com/haochan1996)  
> URL: http://localhost:1313/csharp/wpf/5f4d026a-d256-4989-a4e0-7399716779bd/  

