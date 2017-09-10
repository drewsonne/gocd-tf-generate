package main

import (
	"flag"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/drewsonne/gocd-tf-generate/generator"
	"context"
	"fmt"
)

func main() {
	var id, resource string
	flag.StringVar(&resource, "resource", "",
		"Name of the resource to generate tf config for. Use '*' for all resources. Default: * ."+
			"Allowed Values: (pipeline_template)")
	flag.StringVar(&id, "id", "*",
		"Specify the ID of the resource to generate a specific resource. Omit this to generate all.")
	flag.Parse()

	cfg, err := gocd.LoadConfig()
	if err != nil {
		panic(err)
	}
	c := cfg.Client()

	var output string
	if resource == "pipeline_template" {
		ctx := context.Background()
		templates, _, err := c.PipelineTemplates.List(ctx)
		if err != nil {
			panic(err)
		}

		if id != "" && id != "*" {
			for _, template := range templates {
				if template.Name == id {
					templates = []*gocd.PipelineTemplate{template}
					break
				}
			}
		}

		for _, templateDescription := range templates {
			template, _, err := c.PipelineTemplates.Get(ctx, templateDescription.Name)
			if err != nil {
				panic(err)
			}
			output, err = generator.RenderPipelineTemplate(template)
			fmt.Println(output)
		}
	}

}
