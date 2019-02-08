<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="viewport">
<title>{{ .Title }}</title>
</head>
<body>
<header id="hBk" class="hBk">
  <ul id="hnav">
    {{ range $index, $li := .Li -}}
      {{- $attr := "" -}}
      {{- range $key, $val := $li.Attr -}}
        {{- $attr = printf "%s %s=\"%s%v\"" $attr $key $val $index -}}
      {{- end -}}
      <li{{ safeAttr $attr }}><a href="{{- $li.Href -}}">{{- $li.Value -}}</a></li>
    {{ end -}}
  </ul>
</header>
<div id="container" class="wrapper">
  <h1>{{ .H1  }}</h1>
</div>
<footer id="fBk" class="fBk">
</footer>
</body>
</html>
