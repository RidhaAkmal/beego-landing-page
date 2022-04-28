{{ template "about/layout.tpl" . }}
{{ define "css" }}
        <link rel="stylesheet" href="/static/css/about.css">
{{ end}}


{{ define "content" }}
        <h2>{{ .Title }}</h2>
        <p> THIS IS PAGE ABOUT WITH TEMPLATE CONTENT</p>
{{ end }}

{{ define "js" }}
   <script src="/static/js/reload.min.js"></script>
{{ end}}