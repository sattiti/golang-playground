---
title: "Templates"
h1: "Templates"
date: 
draft: false
linkTitle: "Templates"
weight: 1
---

## Taxonomies Definitions
### Taxonomy (categories, category, list)
コンテンツを分類するために使用できるカテゴリ

### Term (tags, tag, list)
分類されたカテゴリのキーの集合

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
