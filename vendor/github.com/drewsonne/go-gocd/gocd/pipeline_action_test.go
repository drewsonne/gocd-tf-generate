package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func testPipelineServicePause(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/pause", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		assert.Equal(t, "true", r.Header.Get("Confirm"))
		fmt.Fprint(w, "")
	})
	pp, _, err := client.Pipelines.Pause(context.Background(), "test-pipeline")
	if err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, pp)
}

func testPipelineServiceUnpause(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/unpause", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		assert.Equal(t, "true", r.Header.Get("Confirm"))
		fmt.Fprint(w, "")
	})
	pp, _, err := client.Pipelines.Unpause(context.Background(), "test-pipeline")
	if err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, pp)
}

func testPipelineServiceReleaseLock(t *testing.T) {
	mux.HandleFunc("/api/pipelines/test-pipeline/releaseLock", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		assert.Equal(t, "true", r.Header.Get("Confirm"))
		fmt.Fprint(w, "")
	})
	pp, _, err := client.Pipelines.ReleaseLock(context.Background(), "test-pipeline")
	if err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, pp)
}
