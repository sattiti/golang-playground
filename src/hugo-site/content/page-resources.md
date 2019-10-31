---
title: "Page Resources"
keywords:
desscription:
isCJKLanguage: true
date: 2019-10-31T09:58:55+09:00
publishDate:
expriyDate:
lastmod: 2019-10-31T09:58:55+09:00
draft: false
linkTitle: ""
weight: 2
# tags:
#   -
# categories:
#   -
# images:
#   -

h1: "Page Resources"
---

## Resources
`content` フォルダ内のすべてのファイルは `.Resources` からアクセスすることができる。

## Properties

### .ResourceType
リソースのメインタイプ。

### .Name
front matter で指定することができる。デフォルト値はファイル名。

### .Title
front matter で指定することができる。デフォルト値は `.Name` と同じくファイル名。

### .Peramlink
Absolute URL.

### .RelPeramlink
Relative URL.

### .Content
リソースの本体。

```html
<script>{{ (.Resources.GetMatch "myscript.js").Content | safeJS }}</script>
<img src="{{ (.Resources.GetMatch "mylogo.png").Content | base64Encode }}">
```

### .MediaType
`image/jpg` のような MIME type.

### .MediaType.MainType
MIME type のメインタイプ。  
`application/pdf` の `application` 部分。

### .MediaType.SubType
MIME type のサブタイプ。  
`application/pdf` の `pdf` 部分。

### .MediaType.Suffixes
MIME type のスタフィックス。


## Methods

### .ByType "TYPE_NAME"
Returns the page resources of the given type.

### .Match "PATTERN"
`.Name` の値と比較し、マッチしたものを返す。

### .GetMatch "PATTERN"
`.Match` と同じだが、ファーストマッチだけ返す。


## Page resources metadata
Page resouces metadata はページ内の front matter `resources` 項目で指定することができる。 
`.GetMatch "mainVisual"` を使い、取得する。
```yaml
resources:
  - name: mainVisual
    title: Top Brand
    src: /images/mv.png

  # :counter をつけると自動連番することができる。
  - name: pdf-file:counter
    title: "Specification #:counter"
    src: "**.pdf"

  # Custom key-values.
  - params:
    icon: image
    width: 1000
    height: 400
```
