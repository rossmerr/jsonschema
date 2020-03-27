package parser

const KindTemplate = `
{{- define "kind" -}}
	{{- if isReference . -}}
		{{- template "reference" . -}}
	{{end -}}
	{{- if isInterfaceReference . -}}
		{{- template "interfacereference" . -}}
	{{end -}}
	{{- if isInterface . -}}
		{{- template "interface" . -}}
	{{end -}}
	{{- if isEmbeddedStruct . -}}
		{{- template "embeddedStruct" . -}}
	{{end -}}
	{{- if isStruct . -}}
		{{- template "struct" . -}}
	{{end -}}
	{{- if isArray . -}}
		{{- template "array" . -}}
	{{end -}}
	{{- if isNumber . -}}
		{{- template "number" . -}}
	{{end -}}
	{{- if isInteger . -}}
		{{- template "integer" . -}}
	{{end -}}
	{{- if isString . -}}
		{{- template "string" . -}}
	{{end -}}
	{{- if isBoolean . -}}
		{{- template "boolean" . -}}
	{{end -}}
	{{- if isCustomType . -}}
		{{- template "customtype" . -}}
	{{end -}}
{{- end -}}`
