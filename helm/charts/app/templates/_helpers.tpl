
{{/*
-- APP LEVEL DEFINITIONS --
*/}}

{{/*
Expand the name of the chart.
*/}}
{{- define "app.name" -}}
{{- .Values.appName | default "unnamed-app" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 53 chars because some Kubernetes name fields are limited to 63 characters (by the DNS naming spec),
And 10 characters should be left for the controller name.
If release name contains chart name it will be used as a full name.
*/}}
{{- define "app.fullname" -}}
{{- $name := (include "app.name" .) }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 53 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 53 | trimSuffix "-" }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
App common labels
*/}}
{{- define "app.labels" -}}
helm.sh/chart: {{ include "app.chart" . }}
{{ include "app.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
App selector labels.
Never use directly, only a helper function for "app.labels" and "controller.selectorLabels"
*/}}
{{- define "app.selectorLabels" -}}
app.kubernetes.io/name: {{ include "app.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "app.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "app.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}


{{/*
-- CONTROLLER LEVEL DEFINITIONS --

All named templates under this scope should be called at the next way to include the base values as "root" and the controller values as "controller" :
{{ template "controller.X" (dict "root" . "controller" $controller ) }}
*/}}

{{/*
create a unique controller short name
*/}}
{{- define "controller.name" -}}
{{ $appname := include "app.name" .root }}
{{- printf "%s-%s" $appname .name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified controller name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
*/}}
{{- define "controller.fullname" -}}
{{ $appfullname := include "app.fullname" .root }}
{{- printf "%s-%s" $appfullname .name | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Controller common labels
*/}}
{{- define "controller.labels" -}}
{{ include "app.labels" .root }}
{{ include "controller.controllerOnlySelectorLabels" . }}
{{- end }}

{{/*
controller selector labels.
*/}}
{{- define "controller.selectorLabels" -}}
{{ include "app.selectorLabels" .root }}
{{ include "controller.controllerOnlySelectorLabels" . }}
{{- end }}


{{/*
controller selector labels without app selector labels.
*/}}
{{- define "controller.controllerOnlySelectorLabels" -}}
app.kubernetes.io/controller: {{ include "controller.name" . }}
{{- end }}


{{- define "utils.renderEnvVars" -}}
{{- range $e := . }}
-   name: {{ $e.name | quote }}
    value: {{ $e.value | quote }}
{{- end}}
{{- end}}