package generator

import (
	"github.com/drewsonne/go-gocd/gocd"
	"text/template"
	"bytes"
	"strings"
	"fmt"
)

const STAGE_TEMPLATE = `{{range .Stages}}
{{$stage := .Name -}}
# CMD terraform import gocd_pipeline_stage.{{.Name}} "{{$containerType}}/{{$containerName}}/{{.Name}}"
resource "gocd_pipeline_stage" "{{.Name}}" {
  name = "{{.Name}}"{{if .FetchMaterials}}
  %s = "{{$containerName}}"
  fetch_materials = {{.FetchMaterials}}{{end}}{{if .CleanWorkingDirectory}}
  clean_working_directory = {{.CleanWorkingDirectory}}{{end}}{{if .NeverCleanupArtifacts}}
  never_cleanup_artifacts = {{.NeverCleanupArtifacts}}{{end}}{{if .Jobs}}
  jobs = [{{range .Jobs}}
    "${data.gocd_job_definition.{{.Name}}.json}"{{end}}
  ]{{end}}{{if .EnvironmentVariables}}
  environment_variables = [{{range .EnvironmentVariables}}
    {
      name = "{{.Name}}"{{if .Value}}
      value = "{{.Value}}"{{end}}{{if .EncryptedValue}}
      encrypted_value = "{{.EncryptedValue}}"{{end}}{{if .Secure}}
      secure = {{.Secure}}{{end}}
    },{{end}}
  ]{{end}}
}
{{range .Jobs -}}
{{$job := .Name -}}
data "gocd_job_definition" "{{.Name}}" {
  name = "{{.Name}}"{{if .Tasks}}
  tasks = [{{range $i, $e := .Tasks}}
    "${data.gocd_task_definition.{{$containerName}}_{{$stage}}_{{$job}}_{{$i}}.json}",{{end}}
  ]{{end}}{{if .Timeout}}
  timeout = {{.Timeout}}{{end}}{{if .EnvironmentVariables -}}
  environment_variables = [{{range .EnvironmentVariables}}{
    name = "{{.Name}}"{{if .Value}}
    value = "{{.Value}}"{{end}}{{if .EncryptedValue}}
    encrypted_value = "{{.EncryptedValue}}"{{end}}{{if .Secure}}
    secure = {{.Secure}}{{end}}
  }{{end}}]
  {{- end}}{{if .Resources}}
  resources = [{{.Resources | stringJoin -}}]{{end}}{{if .ElasticProfileID}}
  elastic_profile_id = "{{ .ElasticProfileID }}"{{end}}{{if .Tabs}}
  tabs = [{{range .Tabs}}
    {
      name = "{{.Name}}",
      path = "{{.Path}}"
    },{{end}}
  ]{{end}}{{if .Artifacts}}
  artifacts = [{{range .Artifacts}}
    {
      type = "{{.Type}}",
      source = "{{.Source}}",{{if .Destination}}
      destination = "{{.Destination}}"{{end}}
    }, {{end}}
  ]{{end}}{{if .Properties -}}
  properties = [{{range .Properties}}{
      name = "{{.Name}}",
      source = "{{.Source}}",
      xpath = "{{.XPath}}"
    }, {{end}}
  ]{{- end}}
}
{{range $i, $task := .Tasks -}}
data "gocd_task_definition" "{{$containerName}}_{{$stage}}_{{$job}}_{{$i}}" {
  type = "{{.Type}}"{{with .Attributes}}{{if .RunIf}}
  run_if = [
	{{.RunIf | stringJoin -}}]{{if .Command}}{{end}}
  command = "{{.Command}}"{{end}}{{if .Arguments}}
  arguments = [
    {{.Arguments | stringJoin -}}]{{end}}{{if .WorkingDirectory}}
  working_directory = "{{.WorkingDirectory}}"{{end}}{{if .Target}}
  target = "{{.Target}}"{{end}}{{if .Pipeline}}
  pipeline = "{{.Pipeline}}"{{end}}{{if .Stage}}
  stage = "{{.Stage}}"{{end}}{{if .Job}}
  job = "{{.Job}}"{{end}}{{if .IsSourceAFile}}
  is_source_a_file = "{{.IsSourceAFile}}"{{end}}{{if .Destination}}
  destination = "{{.Destination}}"{{end}}{{if .Source}}
  source = "{{.Source}}"{{end}}
{{end}}}
{{end -}}
{{end}}
{{- end}}`

func RenderPipelineTemplate(pt *gocd.PipelineTemplate) (string, error) {
	tplt := fmt.Sprintf(`## START pipeline_template.{{.Name}}
# CMD terraform import gocd_pipeline_template.{{.Name}} "{{.Name}}"
{{$containerName := .Name -}}
{{$containerType := "template" -}}
resource "gocd_pipeline_template" "{{.Name}}" {
  name = "{{$containerName}}"
}

%s
## END`, fmt.Sprintf(STAGE_TEMPLATE, "pipeline_template"))

	fmap := template.FuncMap{
		"stringJoin": func(rawStrings []string) (string, error) {
			if len(rawStrings) > 0 {
				escapedStrings := []string{}
				for _, rawString := range rawStrings {
					escapedStrings = append(escapedStrings, strings.Replace(rawString, "\"", "\\\"", -1))
				}
				return "\"" + strings.Join(escapedStrings, "\",\n\"") + "\"", nil
			}
			return "", nil
		},
	}
	t, err := template.New("pipeline_template").
		Funcs(fmap).
		Parse(tplt)
	if err != nil {
		return "", err
	}

	w := new(bytes.Buffer)
	if err := t.Execute(w, pt); err != nil {
		return "", err
	}

	return w.String(), nil
}
