---
title: "Markdown"
h1: "Markdown"
date: 
draft: false
linkTitle: "Markdown"
weight: 1
---
## Markdown

### hx

``` plain
h1 #
h2 ##
h3 ###
h4 ####
h5 #####
h6 ######
```

### br
`<SPACE><SPACE>` 行末にスペース2つ。

### p
```plain
TEXT TEXT TEXT  

TEXT TEXT TEXT  
```

### text-decoration

#### italic
``` plain
*WORD*
_WORD_
```

#### strong
``` plain
**WORD**
__WORD__
```

#### italic + strong
``` plain
***WORD***
___WORD___
```

#### through
``` plain
~~WORD~~
```

### List
#### ul
{{% highlight md %}}
``` plain
- 1
- 2
  + 2-1
  + 2-2
  + 2-3
    - 2-3-1
    - 2-3-2
      + 2-3-2-1
      + 2-3-2-2
- 3
- 4
```
{{% /highlight %}}

#### ol
``` plain
1. a
2. b
  1. b-a
  2. b-b
  3. b-c
  4. b-d
```

### blockquote
```plain
> aaa
> ccc
```

### pre
先頭に半角スペース 4つもしくは、tab
```plain
    class A
        def aa
        end
    end
```

### code
``` plain
`CODE CODE`
```

### hr
``` plain
---
***
* * *
```

### link
```plain
[NAME](https://www.)
[NAME](https://www. "TITLE")

[NAME][1]
[NAME][2]
[1]: https:/www.
[2]: https:/www.

[NAME][LINK_NAME0]
[NAME][LINK_NAME1]
[LINK_NAME0]: https://www.
[LINK_NAME1]: https://www.
```

### img
```plain
![ALT](https://sourceforge.net/images/icon_linux.gif)
![ALT](https://sourceforge.net/images/icon_linux.gif "TITLE")

![ALT][id]
[id]: https://?height=50px&width=300px&classes=a,b,c "TITLE"
```

### input[type="checkbox"]
- [x] qqq
- [ ] aaa


### table
```plain
  a |            b | c      | d 
 ---|-------------:|:------:|---
  1 | align right  | center | 4 
```

### Comments
```plain
<!-- COMMENTS -->
```

### raw html
```html
<div class="bk">
  <p>hello world</p>
</div>
```
