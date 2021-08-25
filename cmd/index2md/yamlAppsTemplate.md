{{ printf "# %s" .Title }}

{{- range $key, $value := .Entries }}
{{ printf "## %s" $key }}

{{- range $value}}
{{printf "#### Version **%s**" .Version}}
{{printf "> Generated %s" .Created}}

{{printf "App Version **%s**" .AppVersion}}

{{- range .Urls}}
{{printf "[%s](%s)" . .}}
{{- end}}
{{- end}}
{{- end}}