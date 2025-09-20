---
title: "PySide6开发学生管理系统"
date: 2025-09-20T18:29:02+0800
slug: "86b4bffc-1d91-4c49-a0f9-5989ec3060fc"
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
  - Python
  - PySide6
  - GUI
categories:
  - PySide6
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

## 技术要点

1. PySide6核心技术：

- QWidget,QApplication基础应用结构
- QVBoxLayout,QHBoxLayout布局管理
- QTableWidget高级表格控件使用
- QHeaderView实现复杂表头
- QCheckBox,QPushButton,QlineEdit等基础控件应用
- 信号与槽机制应用

2. 自定义组件开发

- 继承QHeaderView实现CustomHeaderView

3. 数据库集成（MySQL8）

- 设计DatabaseManager基类
- 实现StudentDB,ClassB,UserDB等特定数据库操作
- 执行复杂SQL查询（JOIN，WHERE，ORDER BY等）
- 参数化防止SQL注入
- 使用上下文管理器（with语句）确保数据库连接正确关闭

4. 多线程处理

- QThreadPoll和自定义Work类实现
- 使用PySide6信号机制在线程间通讯
- 异步加载数据，保证UI响应性

5. 用户界面设计

- 使用QFluentWidgets实现现代化UI
- 响应式设计，自适应窗口大小变化
- 自定义样式表（StyleSheet)美化界面

6.用户认证与权限管理

- 实现登录系统login.py
- 使用bcrypt进行密码加密和验证
- 基于角色的访问控制（RBAC)实现

7. 高级Python特性应用

- 上下文管理器（with语句）
- 装饰器使用
- 列表推导式和生成器表达式
- 异常处理和自定义异常

8. 软件架构设计

- 模块化设计
- MVC（模型-视图-控制器）架构思想
- 单例模式实现（user_session.py)

9. 数据处理和验证

- 输入验证和错误处理
- 数据转换（如班级ID与名称的映射）
- 批量数据操作（批量删除功能）

10. 高级表格功能

- 自定义排序
- 动态行高调整
- 集成操作按钮到表格单元格

11. 导入、导出功能

- 使用openpyxl处理Excel文件
- 数据验证和错误处理
- 批量数据导入和导出

## 开发环境

操作系统：Windows、Mac、Linux均可

Python解释器：Python3.9及以上

开发工具：Pycharm、VsCode都可

数据库：MySQL8数据库

数据库管理软件：DBeaver Community

## 创建数据库和设计数据表

### 创建数据库

数据库名称：student_system

字符集：utf8mb4

排序规则：utf8mb4_0900_ai_ci

![image-20250920190614674](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250920190614674.png)

### MySQL数据库设计

1. classes表

classes表用于存储班级的基本信息。

| 字段名     | 数据类型     | 约束条件                    | 描述               |
| ---------- | ------------ | --------------------------- | ------------------ |
| class_id   | INT          | PRIMARY KRY,AUTO_INCREAMENT | 班级ID、主键、自增 |
| class_name | VARCHAR(255) | NOT NULL                    | 班级名称           |

创建classes表的SQL语句：

```sql
create table classes(
	class_id INT primary key auto_increment,
	class_name VARCHAR(255) not null
);
```

2. student表

student表用于存储学生的基本信息及成绩。字段及属性如下：

| 字段名         | 数据类型     | 约束条件                                                     | 描述                                  |
| -------------- | ------------ | ------------------------------------------------------------ | ------------------------------------- |
| student_id     | INT          | PRIMARY KRY,AUTO_INCREAMENT                                  | 学生ID、主键、自增                    |
| student_name   | VARCHAR(255) | NOT NULL                                                     | 学生姓名                              |
| student_number | VARCHAR(255) | NOT NULL UNIQUE                                              | 学号，唯一性约束                      |
| gender         | INT          | NOT NULL                                                     | 性别，1表示“男”，2表示”女“            |
| class_id       | INT          | NOT NULL，FOREIGN KEY(class_id) REFERENCES classes(class_id) | 班级ID，外键，关联classes表的class_id |
| chinese_score  | FLOAT        |                                                              | 语文分数                              |
| math_score     | FLOAT        |                                                              | 数学分数                              |
| english_score  | FLOAT        |                                                              | 英语分数                              |

创建sutdent表的SQL语句：

```sql
create table student(
	student_id INT primary key auto_increment,
	student_name VARCHAR(255) not null,
	student_number VARCHAR(255) not null unique,
	gender INT not null,
	class_id INT not null,
	chinese_score FLOAT,
	math_score FLOAT,
	english_score FLOAT,
	foreign key(class_id) references classes(class_id)
);
```

3. user表

user表用于存储用户的基本信息及成绩。字段及属性如下：

| 字段名    | 数据类型     | 约束条件                                                     | 描述                                                         |
| --------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| user_id   | INT          | PRIMARY KRY,AUTO_INCREAMENT                                  | 用户ID、主键、自增                                           |
| user_name | VARCHAR(255) | NOT NULL UNIQUE                                              | 用户姓名，唯一性约束                                         |
| password  | VARCHAR(255) | NOT NULL                                                     | 密码                                                         |
| nickname  | VARCHAR(255) |                                                              | 昵称                                                         |
| user_role | INT          |                                                              | 用户角色，1表示管理员，2表示老师                             |
| class_id  | INT          | NOT NULL，FOREIGN KEY(class_id) REFERENCES classes(class_id) | 管理班级，存储为‘1，2，3’形式的字符串，关联classes的class_id字段 |

创建user表的SQL语句：

```sql
create table user(
	user_id INT primary key auto_increment,
	user_name VARCHAR(255) not null unique,
	password VARCHAR(255) not null,
	nickname VARCHAR(255),
	user_role INT,
	class_id VARCHAR(255)
);
```

## 创建项目

使用Pycharm创建项目，选择创建项目虚拟环境。

![image-20250920194021485](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250920194021485.png)

Alt+F12打开终端，虚拟环境下安装pyside6

```bash
pip install pyside6
```

