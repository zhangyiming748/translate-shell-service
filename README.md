---
title: 默认模块
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

* <a href="http://trans.zhangyiming748.eu.org/api/v1">正式环境: http://trans.zhangyiming748.eu.org/api/v1</a>

# Authentication

# Default

## GET 心跳检测

GET /health

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user|query|string| 否 |none|

> 返回示例

> 200 Response

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

## POST 翻译功能

POST /translate

> Body 请求参数

```json
{
  "src": "hello",
  "proxy": "http://127.0.0.1:8889"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» src|body|string| 是 |原文|
|» proxy|body|string| 否 |代理地址|

> 返回示例

> 200 Response

```json
{
  "dst": "string",
  "msg": "string"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» dst|string|true|none||译文|
|» msg|string|true|none||返回通过gin自带方法检测的ip|
