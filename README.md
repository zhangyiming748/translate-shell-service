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
  "keyword": "Alohomora"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» src|body|string| 是 |原文|
|» proxy|body|string| 否 |代理地址|
|» keyword|body|string| 否 |关键词|

> 返回示例

> 200 Response

```json
{
  "dst": "你好",
  "msg": {
    "status": "fail",
    "message": "reserved range",
    "continent": "",
    "continentCode": "",
    "country": "",
    "countryCode": "",
    "region": "",
    "regionName": "",
    "city": "",
    "district": "",
    "zip": "",
    "lat": 0,
    "lon": 0,
    "timezone": "",
    "offset": 0,
    "currency": "",
    "isp": "",
    "org": "",
    "as": "",
    "asname": "",
    "reverse": "",
    "mobile": false,
    "proxy": false,
    "hosting": false,
    "query": "127.0.0.1"
  }
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
|» msg|object|true|none||返回通过gin自带方法检测的ip|
|»» status|string|true|none||none|
|»» message|string|true|none||none|
|»» continent|string|true|none||none|
|»» continentCode|string|true|none||none|
|»» country|string|true|none||none|
|»» countryCode|string|true|none||none|
|»» region|string|true|none||none|
|»» regionName|string|true|none||none|
|»» city|string|true|none||none|
|»» district|string|true|none||none|
|»» zip|string|true|none||none|
|»» lat|integer|true|none||none|
|»» lon|integer|true|none||none|
|»» timezone|string|true|none||none|
|»» offset|integer|true|none||none|
|»» currency|string|true|none||none|
|»» isp|string|true|none||none|
|»» org|string|true|none||none|
|»» as|string|true|none||none|
|»» asname|string|true|none||none|
|»» reverse|string|true|none||none|
|»» mobile|boolean|true|none||none|
|»» proxy|boolean|true|none||none|
|»» hosting|boolean|true|none||none|
|»» query|string|true|none||none|
