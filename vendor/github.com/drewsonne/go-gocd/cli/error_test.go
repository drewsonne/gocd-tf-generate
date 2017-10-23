package cli

import (
	"github.com/drewsonne/go-gocd/gocd"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"net/http"
	"net/url"
	"testing"
)

func TestError(t *testing.T) {
	t.Run("Basic", testErrorBasic)
	t.Run("Type", testErrorType)
	t.Run("UnexpectedError", testErrorUnexpectedError)
}

func testErrorType(t *testing.T) {
	var err cli.ExitCoder
	err = JSONCliError{}
	_, ok := err.(cli.ExitCoder)
	assert.True(t, ok)
}

func testErrorBasic(t *testing.T) {
	err := NewCliError("TestReqType", nil, errors.New("test-error"))
	assert.Equal(t, `{
  "error": "test-error",
  "request": "TestReqType"
}`, err.Error())

	assert.Equal(t, 1, err.ExitCode())
}

func testErrorUnexpectedError(t *testing.T) {
	u, e := url.Parse("http://example.com/")
	if e != nil {
		t.Error(e)
	}
	err := NewCliError("TestReqType", &gocd.APIResponse{
		HTTP: &http.Response{
			StatusCode: 201,
		},
		Request: &gocd.APIRequest{
			HTTP: &http.Request{
				URL: u,
			},
		},
	}, errors.New("test-error"))
	assert.Equal(t, `{
  "error": "An error occurred while retrieving the resource.",
  "request-body": "",
  "request-endpoint": "http://example.com/",
  "request-header": "null",
  "response-body": "",
  "response-header": "null",
  "status": 201
}`, err.Error())

}
