# gocd-tf-generate
Utility to generate terraform configuration for gocd 

[![GoDoc](https://godoc.org/github.com/drewsonne/gocd-tf-generate/gocd?status.svg)](https://godoc.org/github.com/drewsonne/gocd-tf-generate/gocd)
[![Build Status](https://travis-ci.org/drewsonne/gocd-tf-generate.svg?branch=master)](https://travis-ci.org/drewsonne/gocd-tf-generate)

## CLI

CLI tool to interace with GoCD Server.

### Usage

#### Installation

##### Homebrew

    $ brew tap drewsonne/tap
    $ brew install gocd-tf-generate

##### Manual
Download the latest release from [https://github.com/drewsonne/go-gocd/releases](https://github.com/drewsonne/go-gocd/releases),
and place the binary in your `$PATH`.

#### Quickstart

    $ gocd
    $ gocd list-agents
