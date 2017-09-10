package cli

import (
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/urakozz/go-json-schema-generator"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
)

// List of command name and descriptions
const (
	GenerateJSONSchemaCommandName  = "generate-json-schema"
	GenerateJSONSchemaCommandUsage = "Generates a JSON schema based on the structs in this library"
)

// Schemas is a collection of json structs to write to file as json schemas.
var Schemas = map[string]interface{}{
	"job":           gocd.Job{},
	"agent":         gocd.Agent{},
	"build-details": gocd.BuildDetails{},
	"stage":         gocd.Stage{},
}

// GenerateJSONSchemaAction will generate the list of files for the JSON Schema for the defined structs.
func GenerateJSONSchemaAction(c *cli.Context) error {
	directory := "schema"
	os.Mkdir(directory, os.FileMode(int(0777)))
	for k, s := range Schemas {
		fmt.Printf("Building '%s'...\n", k)
		schema := generator.Generate(s)
		schemaPath := fmt.Sprintf("%s/%s.json", directory, strings.ToLower(k))
		fmt.Printf("Writing '%s' to disk '%s'...\n", k, schemaPath)
		err := ioutil.WriteFile(schemaPath, []byte(schema), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// GenerateJSONSchemaCommand handles the interaction between the cli flags and the action handler for generate-json
func GenerateJSONSchemaCommand() *cli.Command {
	return &cli.Command{
		Name:     GenerateJSONSchemaCommandName,
		Usage:    GenerateJSONSchemaCommandUsage,
		Category: "Schema",
		Action:   GenerateJSONSchemaAction,
	}
}
