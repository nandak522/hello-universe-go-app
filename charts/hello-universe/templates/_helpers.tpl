{{/* vim: set filetype=mustache: */}}

{{- define "app.probe" -}}
{{- .type -}}:
  {{- $probeInfo := .probe -}}
  {{- if $probeInfo.exec }}
  exec: {{ $probeInfo.exec }}
  {{- end }}
  {{- if $probeInfo.failureThreshold }}
  failureThreshold: {{ $probeInfo.failureThreshold }}
  {{- end }}
  {{- if $probeInfo.initialDelaySeconds }}
  initialDelaySeconds: {{ $probeInfo.initialDelaySeconds }}
  {{- end }}
  {{- if $probeInfo.periodSeconds }}
  periodSeconds: {{ $probeInfo.periodSeconds }}
  {{- end }}
  {{- if $probeInfo.successThreshold }}
  successThreshold: {{ $probeInfo.successThreshold }}
  {{- end }}
  {{- if $probeInfo.timeoutSeconds }}
  timeoutSeconds: {{ $probeInfo.timeoutSeconds }}
  {{- end }}
  {{- if $probeInfo.httpGet }}
  httpGet:
    path: {{ $probeInfo.httpGet.path }}
    port: {{ $probeInfo.httpGet.port }}
  {{- end }}
  {{- if $probeInfo.tcpSocket }}
  tcpSocket:
    port: {{ $probeInfo.tcpSocket.port }}
  {{- end }}
{{- end }}

{{/*
Renders nodeSelector block in the manifests
*/}}
{{- define "app.nodeLabels" }}
{{- if . }}
nodeSelector:
  {{ toYaml . }}
{{- end -}}
{{- end -}}

{{/*
Renders Pod's Annotations.
*/}}
{{- define "app.podAnnotations" -}}
{{- toYaml .podAnnotations -}}
{{- end }}
