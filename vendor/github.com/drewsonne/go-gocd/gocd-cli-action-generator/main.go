package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const commandTemplate = `package cli

import (
	"github.com/urfave/cli"
	"context"
)

const (
  %[1]sCommandName  = "%[2]s"
  %[1]sCommandUsage = "%[3]s"
  %[5]sGroup = "%[4]s"
)

func %[1]sAction(c *cli.Context) error {
	_, _, err := cliAgent(c).%[4]s.Get(context.Background(), ...)
	return err
}

func %[1]sCommand() *cli.Command {
	return &cli.Command{
		Name:      %[1]sCommandName,
		Usage:     %[1]sCommandUsage,
		Action:    %[1]sAction,
		Category:  %[5]sGroup,
	}
}

`

func main() {
	var cn string
	flag.StringVar(&cn, "command", "", "Name of the command to generate in hypehn snake case (eg, 'get-resource")
	var dsc string
	flag.StringVar(&dsc, "description", "", "Description for the command")
	var category string
	flag.StringVar(&category, "category", "", "CLI category name (eg, 'Agents', 'Pipelines')")
	var toFile bool
	flag.BoolVar(&toFile, "to-file", false, "If true, print to file")
	var help bool
	flag.BoolVar(&help, "help", false, "Show usage.")
	flag.Parse()

	if cn == "" || dsc == "" || help {
		flag.Usage()
		os.Exit(1)
	}

	if cn == "" || dsc == "" {
		flag.Usage()
		os.Exit(1)
	}

	nameCapitalised := strings.Replace(strings.Title(cn), "-", "", -1)
	nameCapitalised = strings.ToLower(string(nameCapitalised[0])) + nameCapitalised[1:]
	nameLower := strings.ToLower(cn)
	categoryLocal := strings.ToLower(string(category[0])) + category[1:]

	if toFile {
		f, err := os.Create(fmt.Sprintf("./%s.go", cn))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = f.WriteString(fmt.Sprintf(commandTemplate, nameCapitalised, nameLower, dsc, category, categoryLocal))
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf(commandTemplate, nameCapitalised, nameLower, dsc, category, categoryLocal)
	}
}
