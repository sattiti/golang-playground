---
title: "Templates"
h1: "Templates"
date: 
draft: false
linkTitle: "Templates"
weight: 1
---

## Templates 

### index.md と _index.md の違い
_index.md は List タイプの index ページ。(branch bundle)  
`sections` `taxonomies` `taxonomy terms` `homepage` template の index になる。  

index.md は page タイプの index。(leaf bundle)

### TableOfContents
```
{{ .TableOfContents }}
```
で呼び出すことで目次を自動生成することができる。
