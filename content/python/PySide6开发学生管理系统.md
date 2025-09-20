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

使用VsCode创建目录student_system。

创建项目虚拟环境(python -m venv .env)。

激活虚拟环境，安装pyside6。

```bash
pip install pyside6
```

`PyQt Fluent` 是一个用于创建现代风格 PyQt 应用程序的库，它提供了丰富的 UI 组件和主题支持。如果项目中使用的是 PySide2、PySide6 或者 PyQt6，需切换至 PySide2、PySide6 和 PyQt6 分支下载对应的代码。这里安装PySide6-Fluent-Widgets。

```bash
pip install PySide6-Fluent-Widgets -i https://pypi.org/simple/
```

项目结构如下：

```bash
student_system
- student	# 学生相关
	- student_interface.py 	# 学生用户界面基本布局
- classes	# 班级相关
- config # 配置相关
- database
	- base_db.py # 数据库操作基类
	- classes_db.py	# 班级数据操作
	- student_db.py	# 学生数据操作
	- user_db.py	# 用户数据操作
- logs # 记录日志
	- app.log	# 日志文件
- user		# 用户相关
- utils
	- custom_style.py	# 自定义样式
	- log.py 	#日志处理
main.py	# 程序的入口
```

## 实现用界面基本布局

student/student_interface.py中添加如下内容：

```python
from PySide6.QtWidgets import QWidget, QVBoxLayout, QHBoxLayout,QHeaderView, QTableWidgetItem, QCheckBox
from PySide6.QtCore import Qt
from qfluentwidgets import CardWidget, PushButton, SearchLineEdit, TableWidget, setCustomStyleSheet
from utils.custom_style import *

class StudentInterface(QWidget):

    def __init__(self):
        super().__init__()
        self.setWindowTitle("Student Interface")
        self.setGeometry(100, 100, 800, 600)
        self.students = []  # 用于存储学生数据的列表
        self.setupUI()  # 设置UI
        self.loadData() # 模拟加载数据
        self.populate_table() # 填充表格

    def setupUI(self):
        # 创建垂直布局, 用于放置顶部按钮组和表格
        layout = QVBoxLayout()

        # 顶部按钮组
        card_widget = CardWidget(self)
        button_layout = QHBoxLayout(card_widget)
        
        self.addButton = PushButton("新增", self)
        setCustomStyleSheet(self.addButton, ADD_BUTTON_STYLE, ADD_BUTTON_STYLE) # 设置新增按钮样式
        self.searchIput = SearchLineEdit(self)
        self.searchIput.setPlaceholderText("请输入学生姓名或学号")
        self.searchIput.setClearButtonEnabled(True)
        self.searchIput.setFixedWidth(500)
        self.batchDeleteButton = PushButton("批量删除", self)
        setCustomStyleSheet(self.batchDeleteButton, PATCH_DELETE_BUTTON_STYLE, PATCH_DELETE_BUTTON_STYLE) # 设置批量删除按钮样式
        
        button_layout.addWidget(self.addButton) # 添加新增按钮
        button_layout.addWidget(self.searchIput)    # 添加搜索框
        button_layout.addStretch(1) # 添加弹性空间, 默认填充满
        button_layout.addWidget(self.batchDeleteButton)     # 添加批量删除按钮

        layout.addWidget(card_widget)
        
        # 添加table
        self.tableWidget = TableWidget(self)
        self.tableWidget.setBorderRadius(8) # 设置圆角
        self.tableWidget.setBorderVisible(True) # 设置边框可见

        # 设置表头
        self.tableWidget.setColumnCount(11) # 设置列数
        self.tableWidget.setHorizontalHeaderLabels(["", "学生ID", "姓名","学号", "性别", "班级", "语文", "数学", "英语", "总分", "操作"]) # 设置表头标签
        self.tableWidget.horizontalHeader().setSectionResizeMode(QHeaderView.ResizeMode.Stretch)  # 设置表头自适应宽度
        
        
        layout.addWidget(self.tableWidget)
        
        self.setStyleSheet("""
            StudentInterface {
                background-color: white;
            }
        """)
        self.resize(1280, 760)
        self.setLayout(layout)

    def loadData(self):
        # 模拟加载数据
        self.students = [
            {"student_id": 1, "student_name": "张三", "student_number": "2025092001","gender": "1", "class_name": "高三1班", "chinese_score": 85, "math_score": 90, "english_score": 88},
            {"student_id": 2, "student_name": "李四", "student_number": "2025092002","gender": "2", "class_name": "高三2班", "chinese_score": 78, "math_score": 82, "english_score": 80},
            {"student_id": 3, "student_name": "王五", "student_number": "2025092003","gender": "1", "class_name": "高三1班", "chinese_score": 92, "math_score": 88, "english_score": 95},
            {"student_id": 4, "student_name": "赵六", "student_number": "2025092004","gender": "2", "class_name": "高三3班", "chinese_score": 70, "math_score": 75, "english_score": 72},
            {"student_id": 5, "student_name": "孙七", "student_number": "2025092005","gender": "1", "class_name": "高三2班", "chinese_score": 88, "math_score": 90, "english_score": 85},
        ]
    
    def populate_table(self):
        self.tableWidget.setRowCount(len(self.students)) # 设置行数
        for row, student_info in enumerate(self.students):
            self.setup_table_row(row, student_info)
    
    def setup_table_row(self, row, student_info):
        checkbox = QCheckBox()
        self.tableWidget.setCellWidget(row, 0, checkbox) # 添加复选框
        # 赋值其他列
        for col, key in enumerate(["student_id", "student_name", "student_number","gender", 
                                   "class_name", "chinese_score", "math_score", "english_score"]):
            value = student_info.get(key, "")
            if key == 'gender':
                value = "男" if value == "1" else "女"
            item = QTableWidgetItem(str(value)) # 创建表格项
            item.setTextAlignment(Qt.AlignCenter) # 设置文本居中
            self.tableWidget.setItem(row, col + 1, item)    # 设置表格项， col+1 因为第一列是复选框
        total_score = (student_info["chinese_score"] + 
                       student_info["math_score"] + 
                       student_info["english_score"]) # 计算总分
        total_item = QTableWidgetItem(str(total_score)) # 创建总分项
        # total_item.setTextAlignment(Qt.AlignCenter) # 设置文本居中
        self.tableWidget.setItem(row, 9, total_item) # 设置总分列
        
        # 添加操作按钮
        action_widget = QWidget()
        action_layout = QHBoxLayout(action_widget)
        action_layout.setContentsMargins(0, 0, 0, 0) #
        action_layout.setSpacing(5) #
        edit_button = PushButton("编辑")
        setCustomStyleSheet(edit_button, UPDATE_BUTTON_STYLE, UPDATE_BUTTON_STYLE) # 设置编辑按钮样式
        delete_button = PushButton("删除")
        setCustomStyleSheet(delete_button, DELETE_BUTTON_STYLE, DELETE_BUTTON_STYLE) # 设置删除按钮样式
        action_layout.addWidget(edit_button)
        action_layout.addWidget(delete_button)
        action_layout.addStretch() # 添加弹性空间, 默认填充满   
        self.tableWidget.setCellWidget(row, 10, action_widget) # 设置操作列
        self.tableWidget.setRowHeight(row, 40) # 设置行高
        
        # 设置所有列居中
        for col in range(self.tableWidget.columnCount()):
            item = self.tableWidget.item(row, col)
            if item:
                item.setTextAlignment(Qt.AlignCenter)
```

utils/custom_style.py添加如下内容：

```python
BUTTON_STYLE = """
QPushButton {
    color: white;
    border: none;
    font-family: "Segoe UI", "Microsoft Yahei";
    padding: 5px 10px;
    border-radius: 4px;
    font-size: 14px;
}
QPushButton:hover {
    background-color: rgba(255, 255, 255, 0.1);
}
QPushButton:pressed {
    background-color: rgba(255, 255, 255, 0.2);
}
"""

ADD_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {
    background-color: #0d6efd;
}
QPushButton:hover {
    background-color: #0b5ed7;
}
QPushButton:pressed {
    background-color: #0a58ca;
}
"""

DELETE_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {
    background-color: #D70022;
}   
QPushButton:hover {
    background-color: #A3001F;
}
QPushButton:pressed {
    background-color: #7F001A;
}
"""

PATCH_DELETE_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {
    background-color: #E81123;
}   
QPushButton:hover {
    background-color: #C50F1F;
}
QPushButton:pressed {
    background-color: #A8001C;
}   
"""

UPDATE_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {
    background-color: #FF8C00;
}   
QPushButton:hover {
    background-color: #E87700;
}   
QPushButton:pressed {
    background-color: #C86400;  
}
"""

IMPORT_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {   
    background-color: #107C10;
}   
QPushButton:hover { 
    background-color: #0E6C0E;
}
QPushButton:pressed {
    background-color: #0C5E0C;
}
"""    

EXPORT_BUTTON_STYLE = BUTTON_STYLE + """
QPushButton {
    background-color: #5C2D91;
}   
QPushButton:hover { 
    background-color: #4B237A;
}
QPushButton:pressed {
    background-color: #3A1A5E;
}
"""

SEARCH_LINE_EDIT_STYLE = """
QLineEdit {
    border: 1px solid #CCCCCC;
    border-radius: 4px;
    padding: 5px 10px;
    font-size: 14px;
}   
QLineEdit:focus {
    border: 1px solid #0078D7;
}
"""
```

main.py添加如下内容：

```python
from PySide6.QtWidgets import QApplication
from student.student_interface import StudentInterface

if __name__ == "__main__":
    app = QApplication([])
    window = StudentInterface() 
    window.show()
    app.exec()
```

终端输入`python main.py`运行效果如下：

![image-20250920223234710](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250920223234710.png)

## 数据库集成创建基类

上面基础用户界面中实现了基本的界面，在表格中填充的是自定义的假数据，接下来实现从数据库中读取真实的数据。

在python中操作mysql数据库，可以使用一个常用包叫做pymysql，使用前先需要安装`pip install pymysql`。

### 实现日志记录功能

在utils/log.py中创建日志记录器，用来记录日志，内容如下：

```python
import logging
import os
from logging.handlers import RotatingFileHandler 

# 定义日志文件的路径
LOG_DIR = 'logs'
LOG_FILE_DATABASE = os.path.join(LOG_DIR, 'database.log')

# 确保日志目录存在
os.makedirs(LOG_DIR, exist_ok=True)

format='%(asctime)s - %(filename)s:%(lineno)d - %(funcName)s - %(levelname)s - %(message)s'

# 创建一个RotatingFileHandler对象
handler_database = RotatingFileHandler(
    filename=LOG_FILE_DATABASE, # 日志文件路径
    mode= 'a',  # 追加模式
    encoding='utf-8', # 文件编码
    maxBytes=1024*1024*5,  # 每个日志文件最大5MB
    backupCount=3   # 保留3个备份日志文件
    )
handler_database.setFormatter(logging.Formatter(format))

# 获取根日志记录器并添加处理器
logger_database = logging.getLogger("database")
logger_database.setLevel(logging.DEBUG)
logger_database.addHandler(handler_database)
logger_database.addHandler(logging.StreamHandler())  # 同时输出到控制台

```

在base_bd.py中引入logger_database日志记录器，用于记录数据库操作基类中的日志信息。

### 定义配置文件

在config/settings.py中定义数据库了连接的参数，后续可以扩展中环境变量中读取。

```python
# 数据库连接配置
DB_HOST = 'localhost'
DB_USER = 'root'
DB_PASSWORD = 'root'
DB_NAME = 'student_system'
DB_PORT = 3306
DB_CHARSET = 'utf8mb4'
DB_POOL_SIZE = 10
DB_TIMEOUT = 5  # 连接超时时间，单位为秒
DB_RETRY_ATTEMPTS = 3  # 重试连接次数
```

### 实现数据操作基类

在database/base_db.py中添加如下内容：

```python
import pymysql
import logging
from config.settings import DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT
from utils.log import logger_database as logger

class DatabaseManager:
    def __init__(self):
        self.db_connection = None
        
    def __enter__(self):
        '''上下文管理器入口，自动连接数据库'''
        # print("Entering context manager, connecting to database...")
        self.connect()
        return self
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        '''上下文管理器出口，自动断开数据库连接'''
        # print("Exiting context manager, disconnecting from database...")
        self.disconnect()

    def connect(self):
        # 连接数据库
        if not self.db_connection or not self.db_connection.open: # 检查连接是否已存在且打开
            try:
                # 连接数据库
                self.db_connection = pymysql.connect(
                    host=DB_HOST,
                    user=DB_USER,
                    password=DB_PASSWORD,
                    database=DB_NAME,
                    port=DB_PORT,
                    cursorclass=pymysql.cursors.DictCursor # 使用字典游标
                )
                logger.info("Database connection established.")
                # print(self.db_connection)
            except pymysql.MySQLError as e:
                logger.error(f"Error connecting to database: {e}") # 记录错误日志

    def disconnect(self):
        # 断开数据库连接
        if self.db_connection and self.db_connection.open:
            try:
                self.db_connection.close()
                logger.info("Database connection closed.")
            except pymysql.MySQLError as e:
                logger.error(f"Error closing database connection: {e}")
     
    def  fetch_query(self, query, params=None, fetch_one=False): 
        '''查询数据
        :param query: SQL查询语句
        :param params: 可选的查询参数
        :param fetch_one: 是否只获取一条记录
        :return: 查询结果
        '''
        if self.db_connection and self.db_connection.open:
            try:
                with self.db_connection.cursor() as cursor:
                    cursor.execute(query, params)
                    if fetch_one:
                        return cursor.fetchone() # 获取单条记录
                    else:
                        return cursor.fetchall() # 获取所有记录
                logger.info(f"Query executed: {query} with params: {params}")
            except pymysql.MySQLError as e:
                logger.error(f"Error executing query: {e}")
                return None
        else:
            logger.error("Database connection is not established.")
            return None

    def execute_query(self, query, params=None)-> bool:
        '''执行插入、更新或删除操作
        :param query: SQL语句
        :param params: 可选的参数
        :return: 执行是否成功
        '''
        if self.db_connection and self.db_connection.open:
            try:
                with self.db_connection.cursor() as cursor:
                    affected_rows = cursor.execute(query, params)
                    self.db_connection.commit() # 提交事务
                    logger.info(f"Query executed: {query} with params: {params}, affected rows: {affected_rows}")
                    return True
            except pymysql.MySQLError as e:
                self.db_connection.rollback() # 回滚事务
                logger.error(f"Error executing query: {e}")
                return False
        else:
            logger.error("Database connection is not established.")
            return False
```

connect方法实现了连接数据库，disconnect方法实现了断开数据库。fetch_query方法实现了数据查询，query是要执行的SQL语句，fetch_one默认为False，获取所有记录。

execute_query实现数据库插入、更新或删除操作，使用事务方式执行，错误回滚。

`from utils.log import logger_database as logger`引入了之前定义的日志处理器，所有的日志会输出到logs/database.log文件中。

```tex
2025-09-21 00:28:59,413 - base_db.py:73 - fetch_query - ERROR - Database connection is not established.
2025-09-21 00:29:32,096 - base_db.py:36 - connect - INFO - Database connection established.
2025-09-21 00:29:32,099 - base_db.py:46 - disconnect - INFO - Database connection closed.
2025-09-21 00:29:47,373 - base_db.py:34 - connect - INFO - Database connection established.
2025-09-21 00:29:47,375 - base_db.py:44 - disconnect - INFO - Database connection closed.
```

在根目录下新建test_base_bd.py文件，填入如下内容测试基类基本功能：

```python
from database.base_db import DatabaseManager

if __name__ == "__main__":
    with DatabaseManager() as db_manager:
        db_manager.execute_query(
            query=
            "INSERT INTO user (user_name, password, user_role, class_id) VALUES (%s, %s,%s, %s)",
            params=("Alice", "666666", 1, "[1,2,3,4]"))
    with DatabaseManager() as db_manager:
        res = db_manager.fetch_query(query="SELECT * FROM user",
                                     fetch_one=False)
        print(res)
```

输出结果：

```bash
>>python test_base_db.py
Database connection established.
Query executed: INSERT INTO user (user_name, password, user_role, class_id) VALUES (%s, %s,%s, %s) with params: ('Alice', '666666', 1, '[1,2,3,4]'), affected rows: 1
Database connection closed.
Database connection established.
[{'user_id': 6, 'user_name': 'Alice', 'password': '666666', 'nickname': None, 'user_role': 1, 'class_id': '[1,2,3,4]'}]
Database connection closed.
```

