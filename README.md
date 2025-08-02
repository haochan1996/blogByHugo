# blogByHugo
基于Hugo搭建的博客网站

## 安装Hugo

基于Windows的安装方法：

[下载安装包](https://github.com/gohugoio/hugo/releases/tag/v0.148.2)

将Hugo.exe所在的路径添加到系统环境变量中。

## 快速入门

### 创建网站

使用git创建一个仓库blogByHugo，在该仓库所在路径中执行以下命令：

```bash
hugo new site blogByHugo --force
```

### 添加主题

初始化 Hugo 模块系统：`hugo mod init github.com/<your_user>/<your_project>`

导入主题

```bash
hugo mod get github.com/hugo-fixit/FixIt
```