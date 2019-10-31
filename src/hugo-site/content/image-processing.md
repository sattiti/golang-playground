---
title: "Image Processing"
keywords:
desscription:
isCJKLanguage: true
date: 
publishDate:
expriyDate:
lastmod: 
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

h1: Image Processing
---

## Image Resources
プロセスメソッドを使うには画像を page resources として認識させる必要があり、`static` フォルダに置かれた画像リソースには適応しない。

## Methods
### .Resize "WxH EXT BGCOLOR"
```
{{ $img := $resouce.Resize "600x" }}
{{ $img := $resouce.Resize "x400" }}
{{ $img := $resouce.Resize "600x400" }}

# Bgcolor
{{ $img := $resouce.Resize "600x400 jpg #ff0000" }}

# Quality
{{ $img := $resouce.Resize "600x400 q75" }}

# Ratate
{{ $img := $resouce.Resize "600x400 r90" }}

```

### .Fit "WxH"
```
{{ $img := $resouce.Fit "600x400" }}

# Anchor
#   Center Left Top Right Bottom
#   TopLeft TopRight
#   BottomLeft, BottomRight
{{ $img := $resouce.Fill "600x400 LeftTop" }}
```

### .Fill "WxH"
```
{{ $img := $resouce.Fill "600x400" }}
```

### .Filter "(FILTER) (FILTER) .."
Method        | Value                  | Range
--------------|------------------------|-------
.Brightness   | V%                     | -100, 100
.ColorBalance | R% G% B%               | -100, 500
.Colorize     | HUE SATURATION V%      | 0,360 0,100 0,100
.Contrast     | V%                     | -100,100
.Gamma        | V                      | >=1
.GaussianBlur | SIGMA                  |
.Grayscale    |                        |
.Hue          | SHIFT                  | -180, 180
.Invert       |                        |
.Pixelate     | SIZE                   |
.Saturation   | V%                     |
.Sepia        | V%                     |
.Sigmoid      | MIDPOINT FACTOR        |
.UnsharpMask  | SIGMA AMOUNT THRESHOLD |

```
{{ $img := $img.Filter (images.Brightness 1) }}
```

### .Exif
exif を提供する。
