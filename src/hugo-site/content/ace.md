---
title: "Ace - Alternative Templating"
h1: "Ace"
date: 
draft: false
linkTitle: "Ace"
weight: 1
---

## Usage
Go Template の [Action][ActionLink] 機能も兼用可能。  
[Ace Doc](https://github.com/yosssi/ace/blob/master/documentation/syntax.md)  

[ActionLink]: https://golang.org/pkg/text/template/#hdr-Actions


### Doctype Helper Method
Available options:  
`html`, `xml`, `transitional`, `strict`, `frameset`, `1.1`, `basic`, `mobile`

```ace
= doctype html
```

```html
<!DOCTYPE html>
```

### Comment
```ace
/ The Comment is not rendered.

/
  The Comment is not rendered.
  The Comment is not rendered.

//
  The Comment is rendered as an HTML Comment.
  The Comment is rendered as an HTML Comment.
```

### id, class
```ace
.bk0
#main
```

```html
<div class="bk0"></div>
<div id="main"></div>
```

### Attribute
```ace
p
  a.main.link href=/ title=Main Page
    Click Link
```

```html
<p><a class="main link" href="/" title="Main Page">Click Link</a></p>
```


### Block Text
`Tag` + `..` でブロックテキストを利用することができる。

```ace
script..
  var a = 1;
  console.log(1);

p..
  text text text text
  text text text text
```

```html
<script>
var a = 1;
console.log(1);
</script>

<p>
text text text text
text text text text
</p>
```

### Plain Text
A line which starts with a pipe `|` or double pipe `||` is interpreted as a block of plain texts. 
`||` を使うと、改行した箇所の尾部は `<br>` がつく。
```ace
p
  |
    text text text
    text text text
p
  ||
    text text text
    text text text
```

```html
<p>text text text
text text text</p>

<p>text text text<br>
text text text</p>
```

### Conditional Comment Helper Method (IE)
```ace
= conditionalComment COMMENTtYPE CONDITION
```

CommentType	| Generated HTML
------------|----------------
hidden      |
revealed    | <![if expression]> HTML <![endif]>

```ace
= conditionalComment hidden IE 6
  <p>You are using Internet Explorer 6.</p>

= conditionalComment revealed !IE
  <link href="non-ie.css" rel="stylesheet">
```

```html
<!--[if IE 6]>
  <p>You are using Internet Explorer 6.</p>
<![endif]-->

<![if !IE]>
  <link href="non-ie.css" rel="stylesheet">
<![endif]>
```

### Content Helper Method
```ace
= yield main

= content main
  h1 xxxx
```

### CSS Helper Method
```ace
= css
  h1{
    color: red;
  }
```

```css
<style type="text/css">
  h1{
    color: red;
  }
```

### Javascript Helper Method
```ace
= javascript
  var msg = "hello world";
  console.log(msg);
```

```html
<script type="text/javascript">
var msg = "hello world";
console.log(msg);
</script>
```

### Include Helper Method
partials を呼び出す。
`= include TEMPLATE_PATH_WITHOUT_EXTENSION pipeline (parameter)`

```ace
= include header .
```
