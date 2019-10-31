---
title: "Home"
h1: "Home"
date: 
draft: false
linkTitle: "Home"
weight: 1
---

## Links
- [Hugo Offical Documentation][HugoOfficalDoc]
- [Netlify Doc][netlify]
- [Hugo 日本語][HugoJaDoc]
- {{< extlink href="https://www.google.co.jp/" linkTitle="Google" >}}

[netlify]: https://learn.netlify.com/en/basics/
[HugoOfficalDoc]: https://gohugo.io/documentation/
[HugoJADoc]: https://note.yuuniworks.com/study/hugo.html


## Basic Usages

### Includes
#### partial
`/layouts/partials` 内にあるファイルを呼び出す。
```html
{{ partial "<PATH>/<PARTIAL>.<EXTENSION>" . }}
```

#### template
internal template を呼び出す。
`/layouts/_internal/`
```html
{{ template "_internal/<TEMPLATE>.<EXTENSION>" . }}
```

### Iteration
#### range
```html
# array
{{ range $IDX, $VAL := $ARRAY }}
{{ end }}

# map
{{ range $KEY, $VAL := $MAP }}
{{ end }}
```

### Conditionals
#### with
```html
{{ with .Param "description" }}
  <h1>{{ . }}</h1>
{{ else }}
  <p>{{ .Summary }}</p>
{{ end }}
```

#### if .. else if .. else
```html
{{ if (isset .Params "description") }}
  {{ index .Params "description" }}
{{ else if (isset .Params "summary") }}
  {{ index .Params "summary" }}
{{ else }}
  {{ .Summary }}
{{ end }}
```

#### and, or
```html
{{ if (and (or (isset .Params "title") (isset .Params "caption")) (isset .Params "attr")) }}
```

### Pipes
As like as Unix pipes.  
example:
```html
# shuffle
<p>{{ seq(1 10) | shuffle }}</p>
```

### Comments
```html
{{/* COMMENTS */}}
```
