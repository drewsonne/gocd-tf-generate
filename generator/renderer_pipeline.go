package generator

import (
	"fmt"
	"strings"
	"bytes"
	"github.com/drewsonne/go-gocd/gocd"
	"text/template"
)

func RenderPipeline(pt *gocd.Pipeline, group string, debug bool) (string, error) {
	tplt := fmt.Sprintf(`## START pipeline.{{.Name}}
#CMD: terraform import gocd_pipeline.{{.Name}} "{{.Name}}";
{{$containerName := .Name -}}
{{$containerType := "pipeline" -}}
{{$defaultLabel := "${COUNT}"}}
resource "gocd_pipeline" "{{.Name}}" {
  name = "{{$containerName}}"
  group = "%s"{{if .Template}}
  template = "{{.Template}}"{{end}}{{if ne .LabelTemplate $defaultLabel}}
  label_template  = "{{.LabelTemplate | escapeDollar}}"{{end}}{{if .EnablePipelineLocking}}
  enable_pipeline_locking = "{{.EnablePipelineLocking}}"{{end}}{{if .Label}}
  label = "{{.Label}}"{{end}}{{if .Parameters}}
  parameters { {{range .Parameters}}
      {{.Name}} = "{{.Value}}",
  {{end}} }
  {{end}}{{if .EnvironmentVariables}}
  environment_variables = [{{range .EnvironmentVariables}}{
      name = "{{.Name}}",{{if .Value}}
      value = "{{.Value}}"{{end}}{{if .EncryptedValue}}
      encrypted_value = "{{.EncryptedValue}}"{{end}}{{if .Secure}}
      secure = "{{.Secure}}"{{end}}
	}, {{end}}
  ]{{end}}{{if .Materials}}
  materials = [{{range .Materials}}
    {
      type = "{{.Type}}"{{if .Description}}
      description = "{{.Description}}"{{end}}
      attributes { {{with .Attributes}}{{if .URL}}
        url = "{{.URL}}"{{end}}{{if .Destination}}
        destination = "{{.Destination}}"{{end}}{{if .Filter}}
        filter = {
          ignore = [{{.Filter.Ignore | stringJoin -}}]
        }{{end}}{{if .InvertFilter}}
        invert_filter = {{.InvertFilter}}{{end}}{{if .Name}}
        name = "{{.Name}}"{{end}}{{if .Branch}}
        branch = "{{.Branch}}"{{end}}{{if .SubmoduleFolder}}
        submodule_folder = "{{.SubmoduleFolder}}"{{end}}{{if .ShallowClone}}
        shallow_clone = {{.ShallowClone}}{{end}}{{if .Pipeline}}
        pipeline = "{{.Pipeline}}"{{end}}{{if .Stage}}
        stage = "{{.Stage}}"{{end}}
      }{{end}}
    }, {{end}}
  ]{{end}}
}

%s
## END`, group, fmt.Sprintf(STAGE_TEMPLATE, "pipeline"))

	fmap := template.FuncMap{
		"stringJoin": templateStringJoin,
		"escapeDollar": func(s string) (string, error) {
			str := strings.Replace(s, "$", "$$", -1)
			return str, nil
		},
	}
	if debug {
		fmt.Printf("Generating terraform configuration template...")
	}
	t, err := template.New("pipeline").
		Funcs(fmap).
		Parse(tplt)
	if err != nil {
		return "", err
	}
	if debug {
		fmt.Printf("Generated terraform configuration template.\nRendering template...")
	}

	for _, m := range pt.Materials {
		if m.Type == "dependency" {
			if m.Attributes.Name == m.Attributes.Pipeline {
				m.Attributes.Name = ""
			}
		}
	}

	w := new(bytes.Buffer)
	if err := t.Execute(w, pt); err != nil {
		return "", err
	}
	if debug {
		fmt.Printf("Template rendered.")
	}
	return w.String(), nil
}
