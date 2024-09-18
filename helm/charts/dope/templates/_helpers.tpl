{{/*
Create the name of the apps namespace/project
*/}}
{{- define "appsProjectName" -}}
{{- if hasKey .Values "environment" }}
{{- default "apps" .Values.environment.appsProjectNameOverride }}
{{- else }}
{{- "apps" }}
{{- end }}
{{- end }}