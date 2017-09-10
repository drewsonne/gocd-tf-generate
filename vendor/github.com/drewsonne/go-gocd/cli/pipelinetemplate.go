package cli

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urfave/cli"
)

// List of command name and descriptions
const (
	ListPipelineTemplatesCommandName   = "list-pipeline-templates"
	ListPipelineTemplatesCommandUsage  = "List Pipeline Templates"
	GetPipelineTemplateCommandName     = "get-pipeline-template"
	GetPipelineTemplateCommandUsage    = "Get Pipeline Templates"
	CreatePipelineTemplateCommandName  = "create-pipeline-template"
	CreatePipelineTemplateCommandUsage = "Create Pipeline Templates"
	UpdatePipelineTemplateCommandName  = "update-pipeline-template"
	UpdatePipelineTemplateCommandUsage = "Update Pipeline template"
	DeletePipelineTemplateCommandName  = "delete-pipeline-template"
	DeletePipelineTemplateCommandUsage = "Delete Pipeline template"
)

// ListPipelineTemplatesAction lists all pipeline templates.
func ListPipelineTemplatesAction(c *cli.Context) error {
	ts, r, err := cliAgent(c).PipelineTemplates.List(context.Background())
	if err != nil {
		return handleOutput(nil, r, "ListPipelineTemplates", err)
	}

	type p struct {
		Name string `json:"name"`
	}

	type ptr struct {
		Name      string `json:"name"`
		Pipelines []*p   `json:"pipelines"`
	}
	responses := []ptr{}
	for _, pt := range ts {
		pt.RemoveLinks()
		ps := []*p{}
		for _, pipe := range pt.Pipelines() {
			ps = append(ps, &p{pipe.Name})
		}

		responses = append(responses, ptr{
			Name:      pt.Name,
			Pipelines: ps,
		})
	}
	return handleOutput(responses, r, "ListPipelineTemplates", err)
}

// GetPipelineTemplateAction checks template-name is provided, and that the response is 2xx.
func GetPipelineTemplateAction(c *cli.Context) error {
	var name string
	if name = c.String("template-name"); name == "" {
		return handleOutput(nil, nil, "GetPipelineTemplate", errors.New("'--template-name' is missing"))
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Get(context.Background(), name)
	if r.HTTP.StatusCode != 404 {
		pt.RemoveLinks()
	}
	return handleOutput(pt, r, "GetPipelineTemplate", err)
}

// CreatePipelineTemplateAction checks stages and template-name is provided, and that the response is 2xx.
func CreatePipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handleOutput(nil, nil, "CreatePipelineTemplate", errors.New("'--template-name' is missing"))
	}
	if len(c.StringSlice("stage")) < 1 {
		return handleOutput(nil, nil, "CreatePipelineTemplate", errors.New("At least 1 '--stage' must be set"))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return handleOutput(nil, nil, "CreatePipelineTemplate", err)
		}
		stages = append(stages, &st)
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Create(context.Background(), c.String("template-name"), stages)
	return handleOutput(pt, r, "CreatePipelineTemplate", err)
}

// UpdatePipelineTemplateAction checks stages, template-name and template-version is provided, and that the response is
// 2xx.
func UpdatePipelineTemplateAction(c *cli.Context) error {
	if c.String("template-name") == "" {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("'--template-name' is missing"))
	}
	if c.String("template-version") == "" {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("'--version' is missing"))
	}
	if len(c.StringSlice("stage")) < 1 {
		return handeErrOutput("UpdatePipelineTemplate", errors.New("At least 1 '--stage' must be set"))
	}

	stages := []*gocd.Stage{}
	for _, stage := range c.StringSlice("stage") {
		st := gocd.Stage{}
		json.Unmarshal([]byte(stage), &st)

		if err := st.Validate(); err != nil {
			return handeErrOutput("UpdatePipelineTemplate", err)
		}
		stages = append(stages, &st)
	}

	ptr := gocd.PipelineTemplate{
		Version: c.String("template-version"),
		Stages:  stages,
	}

	pt, r, err := cliAgent(c).PipelineTemplates.Update(context.Background(), c.String("template-name"), &ptr)
	return handleOutput(pt, r, "UpdatePipelineTemplate", err)
}

// DeletePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// delete-pipeline-template and checks a template-name is provided and that the response is a 2xx response.
func DeletePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     DeletePipelineTemplateCommandName,
		Usage:    DeletePipelineTemplateCommandUsage,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."}},
		Action: func(c *cli.Context) error {
			if c.String("template-name") == "" {
				return handleOutput(nil, nil, "DeletePipelineTemplate", errors.New("'--template-name' is missing"))
			}

			deleteResponse, r, err := cliAgent(c).PipelineTemplates.Delete(context.Background(), c.String("template-name"))
			if r.HTTP.StatusCode == 406 {
				err = errors.New(deleteResponse)
			}
			return handleOutput(deleteResponse, r, "DeletePipelineTemplate", err)
		},
	}
}

// ListPipelineTemplatesCommand handles the interaction between the cli flags and the action handler for
// list-pipeline-templates
func ListPipelineTemplatesCommand() *cli.Command {
	return &cli.Command{
		Name:     ListPipelineTemplatesCommandName,
		Usage:    ListPipelineTemplatesCommandUsage,
		Action:   ListPipelineTemplatesAction,
		Category: "Pipeline Templates",
	}
}

// GetPipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// get-pipeline-template
func GetPipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     GetPipelineTemplateCommandName,
		Usage:    GetPipelineTemplateCommandUsage,
		Action:   GetPipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Name of the Pipeline Template configuration."},
		},
	}
}

// CreatePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// create-pipeline-template
func CreatePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     CreatePipelineTemplateCommandName,
		Usage:    CreatePipelineTemplateCommandUsage,
		Action:   CreatePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}

// UpdatePipelineTemplateCommand handles the interaction between the cli flags and the action handler for
// update-pipeline-template
func UpdatePipelineTemplateCommand() *cli.Command {
	return &cli.Command{
		Name:     UpdatePipelineTemplateCommandName,
		Usage:    UpdatePipelineTemplateCommandUsage,
		Action:   UpdatePipelineTemplateAction,
		Category: "Pipeline Templates",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template-version", Usage: "Pipeline template version."},
			cli.StringFlag{Name: "template-name", Usage: "Pipeline Template name."},
			cli.StringSliceFlag{Name: "stage", Usage: "JSON encoded stage object."},
		},
	}
}
