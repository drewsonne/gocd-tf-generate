# gocd-cli-action-generator

Used to simplify the creation of new actions in the cli utility.

__NOTE__: Used for development of this package only

## Usage

Generate a command template.

    $ cd gocd-cli-action-generator
    $ go install
    $ cd ../cli
    $ ls -l mytask.go
    > ls: mytask.go: No such file or directory
    $ gocd-cli-action-generator -stdout \
    ... -command=mytask \
    ... -description="This is what my task does"
    >  package main
    ...  
    ...  import "github.com/urfave/cli"
    ...
    ... type Mytask struct {}
    ...
    ... func MytaskCommand() *cli.Command {
    ... 	return &cli.Command{
    ...
    ... 
    $ gocd-cli-action-generator -command=mytask
    $  ls -l mytask.go
    >  -rw-r--r--  1 drews  125639865  294 29 Jul 11:19 mytask.go

