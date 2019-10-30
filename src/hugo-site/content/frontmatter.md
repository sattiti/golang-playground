---
title: "Front matter"
h1: "Front matter"
date: 
draft: false
linkTitle: "Front matter"
weight: 1
---

## User-Defined Variables
任意のキー、値を指定することができる。  
`.Params` よりアクセスすることができる。
```yaml
h1: aaa
h2: bbb
```

```
<h1>{{ .Params.H1 }}</h1>
```


## Front matter Pre-defined Variables

### cascade
オーバーライトしない限り、子孫ページまで辿りづくことができる map。  
_index.md で設定すると便利。

```
cascade:
  banner: 
  age:
```

### draft
下書きの設定。 true にすると、-D オプションをつけないとレンダリングされない。

#### weight
ページ表示順の設定。小さいほど表示優先される。

### headless
true にすると、公開されないように構成される。  
`Permalink` を持たないことと、レンダリングされない。  
`.Site.RegularPages` に属さないが、`.Site.GetPage` よりアクセス可能。

### isCJKLanguage
CJK であるかどうかの設定。

### layout
レイアウトの指定。

### linkTitle
リンク表示用テキストの指定。

### markup
markup 種類の指定。  
`rst` reStructuredText (requiresrst2html)  
`md`  Markdown (default)

### outputs
ebook, Google AMP, JSON 等の outputs format の指定。
```yaml
outputs:
  - html
  - json
  - rss
  - amp
```

outputs format 一覧表示
```
{{ range .AlternativeOutputFormats -}}
  <link rel="{{ .Rel }}" type="{{ .MediaType.Type }}" href="{{ .Permalink | safeURL }}">
{{ end -}}
```

### summary
`.Summary` で使われるテキストの指定。

### type
コンテンツタイプの指定。


### taxonomies
#### tags
`tags` 用の指定。
```yaml
tags:
  -
  -
```

#### categories
`categories` 用の指定。
```yaml
categories:
  -
  -
```

### Meta
#### title
title タグ設定。

#### description
meta description.

#### keywords
meta keywords.

### og
#### audio
og:audio 用 audio ファイルパス設定。
```yaml
audio:
  - 
  - 
```

#### videos
og:video 用の指定。
```yaml
videos:
  -
  -
```

#### series
og:see_also 関連用配列。


### 日付関連
#### date
ページの作成日。

#### publishDate
公開日の設定。将来の日付をセットすると、 --buildFuture オプションを利用しない限り、レンダリングされない。

### expiryDate
有効期限の設定。期限が過ぎたコンテンツは、 --buildExpired オプションをつけないとレンダリングされない。

### lastmod
最終更新日の指定。


### resources 関連
#### resources
該当ページ用の画像、ドキュメント等のリソースの指定。
```yaml
resources:
  - resourceType: image/jpg
    name: header
    src: images/sunset.jpg
```

#### images
該当ページに関連するイメージパスの配列。
```yaml
images:
  -
  -
```


### URL 関連
#### aliases
別名を設定することができる。

```yaml
orginal: /content/posts/my-awesome-post.md

aliases:
  - /posts/my-original-url/
  - /2019/10/31/post00001.html
```

#### slug
url (basename) 尾部の指定。指定した場合、オーバーライトされる。

#### url
root パスから url の設定。
