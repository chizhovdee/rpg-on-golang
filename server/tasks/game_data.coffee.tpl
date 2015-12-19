# Этот файл создан автоматически с помощью gofer game_data:export

{{ range $key, $value := . }}
{{$key}} = require("./game_data/{{$key}}.coffee")
{{end}}

module.exports =
  define: ->
    {{ range $key, $value := . }}
    {{$key}}.define({{marshal $value}})
    {{end}}
