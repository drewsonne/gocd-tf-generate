# GoCD SDK 0.3.13

[![GoDoc](https://godoc.org/github.com/drewsonne/go-gocd/gocd?status.svg)](https://godoc.org/github.com/drewsonne/go-gocd/gocd)
[![Build Status](https://travis-ci.org/drewsonne/go-gocd.svg?branch=master)](https://travis-ci.org/drewsonne/go-gocd)
[![codecov](https://codecov.io/gh/drewsonne/go-gocd/branch/master/graph/badge.svg)](https://codecov.io/gh/drewsonne/go-gocd)
[![codebeat badge](https://codebeat.co/badges/1ea74899-2337-4ea6-aaeb-2cc8037fe362)](https://codebeat.co/projects/github-com-drewsonne-go-gocd-master)

## CLI

CLI tool to interace with GoCD Server.

### Usage

#### Installation

##### Homebrew

    $ brew tap drewsonne/tap
    $ brew install go-gocd

##### Manual
Download the latest release from [https://github.com/drewsonne/go-gocd/releases](https://github.com/drewsonne/go-gocd/releases),
and place the binary in your `$PATH`.

#### Quickstart

    $ gocd configure
    $ gocd list-agents

#### Configuration
The library can either be configured using environment variables, cli flags, or a yaml config file.

##### Environment Variables

 - `$GOCD_SERVER`
 - `$GOCD_USERNAME`
 - `$GOCD_PASSWORD`
 - `$GOCD_SSL_CHECK`
 
##### CLI Flags

 - `--server`
 - `--username`
 - `--password`
 - `--ssl_check`
 
##### YAML Config File

Run `gocd configure` to launch a wizard which will create a file at `~/.gocd.conf`, or create the file manually:

```yaml
server: https://goserver:8154/go
username: admin
password: mypassword
ssl_check: false
```

#### Help

    $ gocd -help

## Library

### Usage

```go
package main
import "github.com/drewsonne/go-gocd/gocd"
```

Construct a new GoCD client and supply the URL to your GoCD server and if required, username and password. Then use the
various services on the client to access different parts of the GoCD API.
For example:

```go
package main
import (
    "github.com/drewsonne/go-gocd/gocd"
    "context"
)

func main() {
    cfg := gocd.Configuration{
        Server: "https://my_gocd/go/",
        Username: "ApiUser",
        Password: "MySecretPassword",
    }
    
    client := gocd.NewClient(&cfg,nil)
    
    // list all agents in use by the GoCD Server
    agents, _, err := client.Agents.List(context.Background())

    ...
}
```

### Usage

## Roadmap ##
This library is still in pre-release. It was initially developed to be an interface for a [gocd terraform provider](https://github.com/drewsonne/terraform-provider-gocd),
which, at this stage, will heavily influence the direction of this library. A list of new features and the expected release
schedule for those features can be found in the project for this github repository.

## Background ##
This library's structure was initially inspired by [https://github.com/google/go-github](https://github.com/google/go-github).
There may still be some vestigial code and structures from this library which will be removed in future revisions and 
before v1.0.0 of this library.
 
## License ##

This library is distributed under the LGPL-style license found in [LICENSE](./LICENSE) file.