{{ printf "# %s" .Title }}
{{ range $key, $value := .Entries }}
{{ printf "## %s" $key }}
{{- range $i, $v :=  $value}}
{{ if eq $i 0 }}
{{printf "#### Version **%s**\n" $v.Version}}
{{printf "> Generated %s" $v.Created}}

{{printf "App Version **%s**" $v.AppVersion}}

{{- range $v.Urls}}
{{printf "[%s](%s)" . .}}
{{- end}}
{{- else }}
<details>
  <summary>Older versions</summary>
{{printf "  <h4>Version <strong>%s</strong></h4>\n" $v.Version}}
{{printf "  <blockquote><p>Generated %s</p></blockquote>" $v.Created}}

{{printf "  <p>App Version <strong>%s</strong></p>" $v.AppVersion}}

{{- range $v.Urls}}
{{printf "<a href=\"%s\">%s</a>" . .}}
{{- end}}

</details>
{{- end }}
{{ end}}
{{- end}}