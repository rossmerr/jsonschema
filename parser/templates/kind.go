package templates

const KindTemplate = `
{{- define "kind" -}}
	{{- if isReference . -}}
		{{- template "reference" . -}}
	{{end -}}
	{{- if isType . -}}
		{{- template "type" . -}}
	{{end -}}
	{{- if isInterface . -}}
		{{- template "interface" . -}}
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
	{{- if isEnum . -}}
		{{- template "enum" . -}}
	{{end -}}
	{{- if isConst . -}}
		{{- template "const" . -}}
	{{end -}}
	{{- if isMethod . -}}
		{{- template "method" . -}}
	{{end -}}
	{{- if isAllOf . -}}
		{{- template "allof" . -}}
	{{end -}}
	{{- if isAnyOf . -}}
		{{- template "anyof" . -}}
	{{end -}}
	{{- if isOneOf . -}}
		{{- template "oneof" . -}}
	{{end -}}
{{- end -}}`
