---
title: "Templates"
h1: "Templates"
date: 
draft: false
linkTitle: "Templates"
weight: 1
---

## Section
「フォルダ」による分類。  

## Taxonomies Definitions
「タグ」による分類。デフォルトでは `tags` `categoies` の2つが Taxonomy として扱う。

```plain
Actor                    <- Taxonomy
    Bruce Willis         <- Term
        The Sixth Sense  <- Value
        Unbreakable      <- Value
        Moonrise Kingdom <- Value
    Samuel L. Jackson    <- Term
        Unbreakable      <- Value
        The Avengers     <- Value
        xXx              <- Value
```

### Taxonomy (categories, category, list)
コンテンツ(キーワード)を分類するために使用できるカテゴリ(キーワードの総称)。

### Term (tags, tag, list)
分類されたカテゴリのキーの集合(キーワードの集合)。

### Value (single)
用語に割り当てられたコンテンツ


## Basic Structures
```plain
/layouts
    /_default
        /baseof.html
        /list.html
        /single.html
    /index.html
    /404.html
    /robots.txt
    /sitemap.xml
```

## 404
**.htaccess**  
```plain
ErrorDocument 404 /404.html
```

**Nginx**
```plain
error_page 404 /404.html;
```
