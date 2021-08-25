{{ printf "# %s" .Title }}
{{ range $key, $value := .Entries }}
{{ printf "## %s" $key }}
{{- range $i, $v :=  $value}}
{{ if gt $i 0 }}
<details>
  <summary>Older versions</summary>
{{ end }}
{{printf "#### Version **%s**\n" $v.Version}}
{{printf "> Generated %s" $v.Created}}

{{printf "App Version **%s**" $v.AppVersion}}

{{- range $v.Urls}}
{{printf "[%s](%s)" . .}}
{{- end}}

{{- if gt $i 0 }}
</details>
{{- end }}
{{ end}}
{{- end}}