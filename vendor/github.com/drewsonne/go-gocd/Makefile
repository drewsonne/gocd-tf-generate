.DEFAULT: test
SHELL:=/bin/bash
TEST?=$$(go list ./... |grep -v 'vendor')

GO_TARGETS= ./cli ./gocd ./gocd-*generator

format:
	gofmt -w -s .
	$(MAKE) -C ./cli/ format
	$(MAKE) -C ./gocd/ format
	$(MAKE) -C ./gocd-cli-action-generator/ format

lint:
	diff -u <(echo -n) <(gofmt -d -s main.go $(GO_TARGETS))
	golint -set_exit_status . $(go list ./... | grep -v vendor/)

test: lint
	go tool vet $(GO_TARGETS)
	go tool vet main.go
	bash ./go.test.sh
	cat coverage.out
#	$(MAKE) -C ./gocd test
#	$(MAKE) -C ./cli test

before_install:
	@go get -t -v $$(go list ./... | grep -v vendor/)
	@go get github.com/golang/lint/golint
	@go install github.com/golang/lint/golint
#	curl -i https://api.github.com/user
#	curl -i https://api.github.com/rate_limit
	curl https://raw.githubusercontent.com/drewsonne/glide.sh/feature/github_auth_get/get | bash -x

build: deploy_on_develop

deploy_on_tag:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser --debug

deploy_on_develop:
	go get github.com/goreleaser/goreleaser
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser --debug --rm-dist --snapshot