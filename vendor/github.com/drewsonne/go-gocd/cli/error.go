package cli

import (
	"encoding/json"
	"fmt"
	"github.com/drewsonne/go-gocd/gocd"
)

// JSONCliError describes an error which outputs JSON on the CLI.
type JSONCliError struct {
	ReqType string
	data    dataJSONCliError
	resp    *gocd.APIResponse
}
type dataJSONCliError map[string]interface{}

// NewFlagError creates an error when a flag is missing.
func NewFlagError(flag string) (err error) {
	return fmt.Errorf("'--%s' is missing", flag)
}

// NewCliError creates an error which can be returned from a cli action
func NewCliError(reqType string, hr *gocd.APIResponse, err error) (jerr JSONCliError) {
	data := dataJSONCliError{
		"error": err.Error(),
	}
	if hr != nil {
		data["status"] = hr.HTTP.StatusCode
		data["response-body"] = hr.Body
		data["request-endpoint"] = hr.Request.HTTP.URL.String()

		if hr.HTTP.StatusCode == 404 {
			data["error"] = "Resource not found"
		} else {
			b1, _ := json.Marshal(hr.HTTP.Header)
			b2, _ := json.Marshal(hr.Request.HTTP.Header)
			data["error"] = "An error occurred while retrieving the resource."
			data["response-header"] = string(b1)
			data["request-body"] = hr.Request.Body
			data["request-header"] = string(b2)
		}
	} else {
		data["request"] = reqType
	}
	return JSONCliError{
		data: data,
		resp: hr,
	}
}

// Error encodes the error as a JSON string
func (e JSONCliError) Error() string {
	b, err := json.MarshalIndent(e.data, "", "  ")
	if err != nil {
		panic(err)
	}

	return string(b)
}

// ExitCode returns the cli statusin the event of an error
func (e JSONCliError) ExitCode() int {
	if e.resp == nil {
		return 1
	}
	code := e.resp.HTTP.StatusCode
	if code >= 100 && code < 200 {
		return 10
	} else if code >= 200 && code < 300 {
		return 20
	} else if code >= 300 && code < 400 {
		return 30
	} else if code >= 400 && code < 500 {
		return 40
	} else if code >= 500 && code < 600 {
		return 50
	}
	return 2
}
