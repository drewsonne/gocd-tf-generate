#! /usr/bin/env bash

# set -x

GOCLI_CMD=${GOCLI_CMD:-gocd}

${GOCLI_CMD} delete-pipeline-config --name go-gocd
${GOCLI_CMD} create-pipeline-config --group go-gocd --pipeline-file examples/go-gocd.json
${GOCLI_CMD} unpause-pipeline --name go-gocd