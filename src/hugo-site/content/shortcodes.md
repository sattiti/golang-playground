---
title: "Shortcodes"
keywords:
desscription:
isCJKLanguage: true
date: 2019-10-31T12:21:33+09:00
publishDate:
expriyDate:
lastmod: 2019-10-31T12:21:33+09:00
draft: false
linkTitle: ""
weight: 2
# resources:
#  - name:
#    title:
#    params:
#      width:
#      height:
# tags:
#   -
# categories:
#   -
# images:
#   -

h1: shortcodes
---

## Shortcodes
Markdown 中で呼び出すことができるスニペット。
2種類の呼び出し方がある。  
中の記法が md の場合 `%%`  
中の記法が md 以外の場合、 `<>`

```html
{{\% SHORTCODE PARAM1 PARAM2 .. %}}
  # aaa
  ## aaa
{{\% /SHORTCODE %}}

{{\< SHORTCODE PARAM1 PARAM2 .. >}}
  <p>111</p>
{{\< /SHORTCODE >}}
```

### Nested Shortcodes
`.Parent` を使って親コンテキストを確認することができる。

### Built-in shortcodes
- figure (html5 figure に相当)
- gist (gist を取得する)
- hightlight (hightlight language)
- instagram (instagram を取得する)
- param (現在のページの `.Params` を取得する)
- ref, relref (`.Peramlink` `.RelPeramlink` を取得する)
- tweet (twitter を取得する)
- vimeo
- youtube
