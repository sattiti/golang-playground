---
title: "Amber - Alternative Templating"
h1: "Amber"
date: 
draft: false
linkTitle: "Amber"
weight: 1
---

## Usage
[Amber Doc](https://github.com/eknkc/amber)

### Doctypes
`!!!` or `doctype` を利用する。  


Available options:  
`5`, `default`, `xml`, `transitional`, `strict`, `frameset`, `1.1`, `basic`, `mobile`

```amber
!!!5
!!! transitional
```

```html
<!DOCTYPE html>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
```


### id, class
```amber
.bk0
#main
```

```html
<div class="bk0"></div>
<div id="main"></div>
```

### Attributes
1つの属性につき、1つの `[]` で括って記述する。

```amber
p
  a[href="/"][title="Main Page"].main.link Click Link
```

```html
<p><a class="main link" href="/" title="Main Page">Click Link</a></p>
```

### Mutil line tag text
`|`
```amber
p
  | aaa
  | bbb
```

```html
```

### Expand Variables
`#{}` を使って変数を展開する。
```amber
$fullname = "aaa" + $.Name
p #{$fullname}
```

```html
<p>Hello</p>
```

### Conditions
#### if .. else if .. else
```amber
div
  if Friends > 10
    p 1
  else if Friends > 5
    p 2
  else
    p 3

div
  #foo ? Name == "Ekin"
  [bar=baz] ? len(Repositories) > 0
```

#### Iterations
```amber
each $r in Repositories
  p #{$r}

each $i, $r in Repositories
  p
    .even ? $i % 2 == 0
    .odd ? $i % 2 == 1
```

### Mixins
```amber
mixin link($href, $title, $text)
  a[href=$href][title=$title] #{$text}

p
  +link("https://www.", "AAA" , "LINKLINK")
```

### Imports
必要な部品を import する。(partials 的なもの)

```amber
# a.amber
p AAA

# b.amber
p BBB

# c.amber
import a
import b
```

```html
div
  p AAA
  p BBB
```

### Inheritance
```amber
# baseof.amber
!!!5
html
  head
    title
  body
    header
      block header

    main
      block main

    footer
      block footer


# section.amber
extends baseof
block header
  h1 MAIN

block main
  h2 CONTENTS

block footer
  p Copyright
```
