build:
	go build

install:
	go install

deploy_on_tag:
	gem install --no-ri --no-rdoc -v "1.8.1" fpm
	go get
	goreleaser

deploy_on_develop:
	gem install --no-ri --no-rdoc fpm
	go get
	goreleaser --rm-dist --snapshot