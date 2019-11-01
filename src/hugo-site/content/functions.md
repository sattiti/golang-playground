---
title: "Functions"
h1: "Functions"
date: 
draft: false
linkTitle: "Functions"
weight: 1
---

## ReadDir "PATH"
The `readDir` function returns an array of `os.FileInfo`.

## readFile "PATH"
Reads a file from disk.


## .Get
shortcode 内でパラメータを取得する。

```plain
.Get 0
.Get "class"
```

## .GetPage
`Page` を取得する。

```plain
.GetPage PATH
```

## .Scratch
変更可能 `map` 的なもの?  
`page` `shortcode` `template` で使用可能。


```plain
# Create a new scratch
$s := newScratch

# .Set
$s.Set "greeting" "Hello"

# .Get
$s.Get "greeting" #> hello

# .Add
$s.Add "greeting" "world" #> Helloworld

# .SetInMap (map を追加)
$s.SetInMap "KEY" "MAPKEY" "VALUE"

# .GetSortedMapValues
$s.GetSortedMapValues "KEY"

# .Delete
$s.Delete "KEY"
```

## anchorize
ハイフンで区切って、anchor で使えるように変換する。

```plain
{{ "hello world" | nchorize }} #> hello-world
```

## default
デフォルト値を設定する。

```plain
<title>{{ .Params.seo_title | default .Title }}</title>
```

## echoParam
`Params` の中身を出力する。

```plain
{{ echoParams DICT KEY }}
```

## errorf
error 出力

```plain
errorf FORMAT INPUT
```

## findRE
Returns a list of strings that match the regular expression.

```plain
findRE PATTERN INPUT [LIMIT]
```

### getenv
Returns the value of an environment variable.

```plain
getenv "HOME"
```

## querify
クエリ生成。

```plain
<a href="https://www.google.com?{{ (querify "q" "test" "page" 3) | safeURL }}">Search</a>
```
