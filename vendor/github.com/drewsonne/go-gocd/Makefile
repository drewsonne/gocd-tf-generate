.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

GO_TARGETS= ./cli ./gocd ./gocd-*generator

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
	$(MAKE) -C ./gocd-cli-action-generator/ format
	$(MAKE) -C ./gocd-response-links-generator/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s main.go $(GO_TARGETS))
	golint -set_exit_status . ./cli ./gocd ./gocd-*-generator

test: lint
	go tool vet $(GO_TARGETS)
	go tool vet main.go
	bash ./go.test.sh
	$(MAKE) -C ./gocd test
	$(MAKE) -C ./cli test

before_install:
	@go get -t -v $$(go list ./... | grep -v vendor/)
	@go get github.com/golang/lint/golint

build: deploy_on_develop

deploy_on_tag:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser

deploy_on_develop:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc fpm
	go get
	goreleaser --rm-dist --snapshot