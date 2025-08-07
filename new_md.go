// 该程序用于生成一个新的 Markdown 文件，包含预设的模板内容。
// 使用方法：
// go run new_md.go <文件路径/文件名.md>
// 例如：
// go run new_md.go content/csharp/basic/布局控件.md
// go run new_md.go content/csharp/wpf/布局控件.md
// go run new_md.go content/posts/创建预设模板内容的Markdown文件.md

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

const template = `---
title: "%s"
date: %s
slug: "%s"
draft: true
author: 
  name: hobby
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
---

## 标题二

正文内容...

`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("用法: go run py_new_md.go <文件路径/文件名.md>")
		fmt.Println("请确保输入的路径存在且有效，可以使用相对或绝对路径。")
		os.Exit(1)
	}

	filePath := os.Args[1]
	title := strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
	dateStr := time.Now().Format("2006-01-02T15:04:05-0700")
	slug := uuid.New().String()

	content := fmt.Sprintf(template, title, dateStr, slug)

	dir := filepath.Dir(filePath)
	// 检查目录是否存在，不存在则报错并退出
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Printf("目录不存在: %s\n", dir)
		os.Exit(1)
	}

	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		fmt.Printf("写入文件失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("已生成: %s\n", filePath)
}
