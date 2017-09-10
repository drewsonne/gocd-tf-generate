package generator

import (
	"fmt"
	"strings"
	"bytes"
	"github.com/drewsonne/go-gocd/gocd"
	"text/template"
)

func RenderPipeline(pt *gocd.Pipeline, group string) (string, error) {
	tplt := fmt.Sprintf(`## START pipeline.{{.Name}}
# CMD terraform import gocd_pipeline.{{.Name}} "{{.Name}}"
{{$containerName := .Name -}}
resource "gocd_pipeline" "{{.Name}}" {
  name = "{{$containerName}}"
  group = "%s"{{if .Template}}
  template = "{{.Template}}"{{end}}{{if .LabelTemplate}}
  label_template  = "{{.LabelTemplate}}"{{end}}{{if .EnablePipelineLocking}}
  enable_pipeline_locking = "{{.EnablePipelineLocking}}"{{end}}{{if .Label}}
  label = "{{.Label}}"{{end}}{{if .Parameters}}
  parameters = [{{range .Parameters}}{
      name = "{{.Name}}",
      value = "{{.Value}}"
    }, {{end}}
  ]{{end}}{{if .EnvironmentVariables}}
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
        filter = ""{{end}}{{if .InvertFilter}}
        invert_filter = {{.InvertFilter}}{{end}}{{if .Name}}
        name = "{{.Name}}"{{end}}{{if .Branch}}
        branch = "{{.Branch}}"{{end}}{{if .SubmoduleFolder}}
        submodule_folder = "{{.SubmoduleFolder}}"{{end}}{{if .ShallowClone}}
        shallow_clone = {{.ShallowClone}}{{end}}{{if .Pipeline}}
        pipeline = "{{.Pipeline}}"{{end}}{{if .Stage}}
        stage = "{{.Stage}}"{{end}}{{if .AutoUpdate}}
        auto_update = {{.AutoUpdate}}{{end}}
      }{{end}}
    }, {{end}}
  ]{{end}}
}

%s
## END`, group, STAGE_TEMPLATE)

	fmap := template.FuncMap{
		"stringJoin": func(s []string) (string, error) {
			if len(s) > 0 {
				return "\"" + strings.Join(s, "\",\n\"") + "\"", nil
			}
			return "", nil
		},
	}
	t, err := template.New("pipeline").
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
