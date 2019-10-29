---
title: "Directory Structure"
h1: "Directory Structure"
date: 
draft: false
linkTitle: "Directory Structure"
weight: 1
---

## Directory Structure

### /archetypes
記事の雛形の定義ファイル用。

### /assets
Hugo Pipes で処理が必要なファイル用。  
`resources` よりアクセス。

### /config
環境設定ファイル用。

```
/config
  /_default
    config.toml
    languages.toml
    menus.en.toml
    menus.ja.toml
    params.toml
  /production
    config.toml
    params.toml
  /staging
    config.toml
    params.toml
```

### /content
md ファイル(記事)用。

### /data
Hugo 内で使用するすべてのデータ用。

### /layouts
レイアウト(html)用。

### /public
公開用一式。

### /resources
ユーザーが触る必要ない。Hugo が処理をする際に使う。

### /static
img, css, js 等の静的ファイル用。  
デプロイ時にすべてのファイルが public にコピーされる。

### /themes
テーマ一式を格納する。

### config.{yaml,toml,json}
設定ファイル。
