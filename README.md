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
cd blogByHugo
git submodule add https://github.com/hugo-fixit/FixIt.git themes/FixIt
```

配置主题

```bash
echo "theme = 'FixIt'" >> hugo.toml
```bash

在站点配置文件中添加一行，指定默认的内容语言。

```bash
echo "defaultContentLanguage = 'zh-cn'" >> hugo.toml
```

### 运行

在blogByHugo目录下执行以下命令：

```bash
hugo server -D
```

这将启动一个本地服务器，通常在`http://localhost:1313/`

## 快速上手

更多详情参考官网文档[FixIt - Hugo 主题](https://fixit.lruihao.cn/zh-cn/)

为了能完整地使用 FixIt 主题的所有功能，务必在站点配置文件中添加以下内容。

```toml
[markup]
  _merge = "shallow"
[outputs]
  _merge = "shallow"
[taxonomies]
  _merge = "shallow"
```

以上配置表示继承 FixIt 主题的 `markup`，`outputs` 和 `taxonomies` 配置。

### 添加内容

给你的网站添加新页面。

```bash
hugo new content posts/my-first-post.md
```

Hugo 在 content/posts 目录中创建了该文件，使用编辑器打开文件。

```markdown
---
title: My First Post
subtitle:
date: 2025-08-02T15:19:33+08:00
slug: 583bc6c
draft: true
---
```

> 请注意，front matter 中的 draft 值为 true。默认情况下，Hugo 在你构建网站时不会发布草稿内容。详细了解 草稿、未来和过期内容。
> 在帖子正文中添加一些 Markdown，但不要更改 draft 值。

保存文件，然后启动 Hugo 的开发服务器来查看站点。你可以运行以下任一命令来包含草稿内容。

```bash
hugo server --buildDrafts
hugo server -D
hugo server -D --disableFastRender
```

由于本主题使用了 Hugo 中的 .Store 来实现一些特性， 非常建议你为 hugo server 命令添加 --disableFastRender 参数来实时预览你正在编辑的文章页面

![image-20250802152215779](https://blog-1301697820.cos.ap-guangzhou.myqcloud.com/blog/image-20250802152215779.png)

当对新内容感到满意时，将 front matter 中的 `draft` 值更改为 `false`，然后保存文件。
### 发布网站

在此步骤中，你将发布你的网站，但不会*部署*它。

当你发布站点时，Hugo 在项目根目录的 `public` 目录中创建整个静态站点。这包括 HTML 文件以及图像、CSS 文件和 JavaScript 文件等资源。

当你发布网站时，你通常不希望包含 [草稿、未来或过期的内容](https://gohugo.io/getting-started/usage/#draft-future-and-expired-content)，命令很简单。

```bash
hugo
```

我们的大多数用户使用 CI/CD 工作流程部署他们的网站，通过推送到他们的 GitHub 或 GitLab 存储库会触发构建和部署。流行的提供商包括 [Vercel](https://vercel.com/)[2](https://fixit.lruihao.cn/zh-cn/documentation/getting-started/quick-start/#fn:2)、[Netlify](https://www.netlify.com/)[3](https://fixit.lruihao.cn/zh-cn/documentation/getting-started/quick-start/#fn:3)、[AWS Amplify](https://aws.amazon.com/amplify/)、[CloudCannon](https://cloudcannon.com/)、[Cloudflare Pages](https://pages.cloudflare.com/)、 [GitHub pages](https://pages.github.com/) 和 [GitLab pages](https://docs.gitlab.com/ee/user/project/pages/)。

要了解如何部署站点，请参阅 [托管和部署](https://gohugo.io/hosting-and-deployment/) 部分。

## 配置 FixIt

详细情况见官方文档[配置 FixIt | FixIt](https://fixit.lruihao.cn/zh-cn/documentation/getting-started/configuration/)

