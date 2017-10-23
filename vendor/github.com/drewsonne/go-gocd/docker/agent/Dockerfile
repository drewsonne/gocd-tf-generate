FROM gocd/gocd-agent-ubuntu-16.04:v17.10.0

ARG GO_GOCD_VERSION=0.6.1

ENV GO_VERSION=1.9
ENV PATH=${PATH}:/usr/lib/go-${GO_VERSION}/bin/:/go-path/bin/:/usr/local/bin/
ENV GOPATH=/go-path/
ENV GOCD_SKIP_SSL_CHECK=1


# # USER root
RUN apt-get update
RUN apt-get -y upgrade && \
    apt-get install git wget software-properties-common -y && \
    add-apt-repository ppa:gophers/archive && \
    apt-get update
RUN apt-get install golang-${GO_VERSION} -y

RUN find / -name go -type f && \
    echo $GOROOT

RUN go get github.com/golang/lint/golint && \
	go install github.com/golang/lint/golint

RUN mkdir -p /go-path/ && \
	mkdir -p /go/ && \
	chown go:go /go-path/ && \
	chown go:go /go/

RUN ls -lah /usr/lib/go-${GO_VERSION}/bin/

COPY gocd /usr/local/bin/gocd

#RUN wget -O /tmp/go-gocd.deb https://github.com/drewsonne/go-gocd/releases/download/${GO_GOCD_VERSION}/gocd-${GO_GOCD_VERSION}-linux-x86_64.deb && \
#    dpkg -i /tmp/go-gocd.deb && \
#    rm -f /tmp/go-gocd.deb


# USER go
# ENV PATH=$PATH:/go/bin
RUN go version && \
    gocd --version