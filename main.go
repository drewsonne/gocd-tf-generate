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
	// Version for the cli tool
	var Version string

	var id, resource, output, profile string
	var outputToFile, debug, printVersion bool
	flag.StringVar(&resource, "resource", "",
		"Name of the resource to generate tf config for. Use '*' for all resources. Default: * ."+
			"Allowed Values: (pipeline_template)")
	flag.StringVar(&id, "id", "*",
		"Specify the ID of the resource to generate a specific resource. Omit this to generate all.")
	flag.StringVar(&profile, "profile", "default", "Specify the gocd profile to use.")
	flag.BoolVar(&outputToFile, "to-file", false, "If supplied, output to file.")
	flag.BoolVar(&debug, "debug", false, "Print debug logging")
	flag.BoolVar(&printVersion, "version", false, "Print version and exit")
	flag.Parse()

	if (printVersion) {
		fmt.Printf("gocd-tf-generate '%s'", Version)
		return
	}

	cfg := gocd.Configuration{}
	if err := gocd.LoadConfigByName(profile, &cfg); err != nil {
		panic(err)
	}
	c := cfg.Client()
	ctx := context.Background()

	doFilter := id != "" && id != "*"

	if debug {
		fmt.Printf("Searching for resource: `%s`\n", resource)
	}
	if resource == "pipeline_template" {
		if debug {
			fmt.Printf("Listing pipeline templates...")
		}
		templates, _, err := c.PipelineTemplates.List(ctx)
		if err != nil {
			panic(err)
		}

		if doFilter {
			if debug {
				fmt.Printf("Searching for '%s' with id '%s'\n", resource, id)
			}
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
			if err != nil {
				panic(err)
			}
			writeOutput(outputToFile, template.Name, []byte(output))
		}
	} else if resource == "pipeline" {
		if debug {
			fmt.Printf("Listing pipelne groups...\n")
		}
		pipelinesGroups, _, err := c.PipelineGroups.List(ctx, "")
		if err != nil {
			panic(err)
		}

		for _, group := range (*pipelinesGroups) {
			if debug {
				fmt.Printf("Searching group '%s' for pipeline with id '%s'.\n", group.Name, id)
			}
			for _, pipeline := range group.Pipelines {
				if doFilter {
					if pipeline.Name == id {
						if debug {
							fmt.Printf("Getting config for pipeline '%s.%s'", group.Name, pipeline.Name)
						}
						pipelineCfg, _, err := c.PipelineConfigs.Get(ctx, pipeline.Name)
						if err != nil {
							panic(err)
						}
						if output, err = generator.RenderPipeline(pipelineCfg, group.Name, debug); err != nil {
							panic(err)
						}
						if err := writeOutput(outputToFile, pipelineCfg.Name, []byte(output)); err != nil {
							os.Exit(1)
						}
						os.Exit(0)
					}
				} else {
					pipelineCfg, _, err := c.PipelineConfigs.Get(ctx, pipeline.Name)
					if err != nil {
						panic(err)
					}
					output, err = generator.RenderPipeline(pipelineCfg, group.Name, debug)
					if err != nil {
						panic(err)
					}
					if err := writeOutput(outputToFile, pipelineCfg.Name, []byte(output)); err != nil {
						os.Exit(1)
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
