# gocd-tf-generate
Utility to generate terraform configuration for gocd 

[![GoDoc](https://godoc.org/github.com/drewsonne/gocd-tf-generate/gocd?status.svg)](https://godoc.org/github.com/drewsonne/gocd-tf-generate/gocd)
[![Build Status](https://travis-ci.org/drewsonne/gocd-tf-generate.svg?branch=master)](https://travis-ci.org/drewsonne/gocd-tf-generate)

## Usage

### Installation

#### Homebrew

    $ brew tap drewsonne/tap
    $ brew install gocd-tf-generate

#### Manual
Download the latest release from [https://github.com/drewsonne/go-gocd/releases](https://github.com/drewsonne/go-gocd/releases),
and place the binary in your `$PATH`.

### Quickstart

    $ gocd
    $ gocd list-agents

### Importing State

Each generated terraform config has an import statement prefix with "CMD". You can extract all the import commands with:

    $ grep -r CMD . | sed 's/^.*CMD //p'

in the directory you are generating your configs in.