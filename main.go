package main

import (
	"flag"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/drewsonne/gocd-tf-generate/generator"
	"context"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var id, resource, output string
	var outputToFile bool
	flag.StringVar(&resource, "resource", "",
		"Name of the resource to generate tf config for. Use '*' for all resources. Default: * ."+
			"Allowed Values: (pipeline_template)")
	flag.StringVar(&id, "id", "*",
		"Specify the ID of the resource to generate a specific resource. Omit this to generate all.")
	flag.BoolVar(&outputToFile, "to-file", false, "If supplied, output to file.")
	flag.Parse()

	cfg, err := gocd.LoadConfig()
	if err != nil {
		panic(err)
	}
	c := cfg.Client()
	ctx := context.Background()

	doFilter := id != "" && id != "*"

	if resource == "pipeline_template" {
		templates, _, err := c.PipelineTemplates.List(ctx)
		if err != nil {
			panic(err)
		}

		if doFilter {
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

			writeOutput(outputToFile, template.Name, []byte(output))
		}
	} else if resource == "pipeline" {
		pipelinesGroups, _, err := c.PipelineGroups.List(ctx, "")
		if err != nil {
			panic(err)
		}

		for _, group := range (*pipelinesGroups) {
			for _, pipeline := range group.Pipelines {
				if doFilter {
					if pipeline.Name == id {
						pipelineCfg, _, err := c.PipelineConfigs.Get(ctx, pipeline.Name)
						if err != nil {
							panic(err)
						}
						output, err = generator.RenderPipeline(pipelineCfg, group.Name)
						if err := writeOutput(outputToFile, pipelineCfg.Name, []byte(output)); err != nil {
							os.Exit(1)
						}
						os.Exit(0)
					}
				}
			}
		}

	}

}

func writeOutput(outputToFile bool, name string, output []byte) error {
	if outputToFile {
		if err := ioutil.WriteFile(name+".tf", output, 0644); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Println(string(output)); err != nil {
			return err
		}
	}
	return nil
}
