    {{/*
Expand the name of the chart.
*/}}
{{- define "versiago.name" -}}
{{- .Chart.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "versiago.fullname" -}}
{{- $name := .Chart.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "versiago.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "versiago.labels" -}}
helm.sh/chart: {{ include "versiago.chart" . }}
{{ include "versiago.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "versiago.selectorLabels" -}}
app.kubernetes.io/name: {{ include "versiago.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "versiago.instanceHostname"}}
{{- first (regexSplit ":" (get (urlParse .) "host") 2) }}
{{- end }}
