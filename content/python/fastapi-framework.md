---
title: "fastapi framework"
date: 2025-09-13T21:59:20+0800
slug: "34962319-3dcf-4e30-ac14-5b0e2ef020e6"
draft: true
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

## Intro

Automtic docs

- Swagger UI
- ReDoc

Just Modern Python

- Python 3.6 +
- Pydantic

Based on open standards

- JSON Schema
- Open API

## Install and Setup

Install and setup python 3.6 +

Create a virtual environment by venv

```bash
python3 -m venv env
```

Activate virtual environment:

```bash
env\Scripts\activate.bat
```

Install fastapi

```bash
pip3 install fastapi
```

Install uvicorn

```bash
pip3 install uvicorn
```

Create a file main.py

```bash
from fastapi import FastAPI

app = FastAPI()

@app.get("/")
async def index():
    return {"Hello": "World"}
```

Run it

```bash
uvicorn main:app --reload
```

## Break it down





