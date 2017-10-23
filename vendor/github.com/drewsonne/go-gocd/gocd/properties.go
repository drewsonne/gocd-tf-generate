package gocd

import (
	"bytes"
	"context"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"strings"
)

// PropertiesService describes Actions which can be performed on agents
type PropertiesService service

// PropertyRequest describes the parameters to be submitted when calling/creating properties.
type PropertyRequest struct {
	Pipeline        string
	PipelineCounter int
	Stage           string
	StageCounter    int
	Job             string
	LimitPipeline   string
	Limit           int
	Single          bool
}

// PropertyCreateResponse handles the parsing of the response when creating a property
type PropertyCreateResponse struct {
	Name  string
	Value string
}

// List the properties for the given job/pipeline/stage run.
func (ps *PropertiesService) List(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job,
	)
	log.Info("Calling `PropertiesServices.List`")
	return ps.commonPropertiesAction(ctx, path, pr.Single)
}

// Get a specific property for the given job/pipeline/stage run.
func (ps *PropertiesService) Get(ctx context.Context, name string, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	path := fmt.Sprintf("/properties/%s/%d/%s/%d/%s/%s",
		pr.Pipeline, pr.PipelineCounter,
		pr.Stage, pr.StageCounter,
		pr.Job, name,
	)
	return ps.commonPropertiesAction(ctx, path, true)
}

// Create a specific property for the given job/pipeline/stage run.
func (ps *PropertiesService) Create(ctx context.Context, name string, value string, pr *PropertyRequest) (bool, *APIResponse, error) {

	log.Info("Calling `PropertiesServices.Create`")
	responseBuffer := bytes.NewBuffer([]byte(""))
	_, resp, err := ps.client.postAction(ctx, &APIClientRequest{
		Path: fmt.Sprintf("/properties/%s/%d/%s/%d/%s/%s",
			pr.Pipeline, pr.PipelineCounter,
			pr.Stage, pr.StageCounter,
			pr.Job, name,
		),
		ResponseType: responseTypeText,
		ResponseBody: responseBuffer,
		RequestBody: strings.Join(
			[]string{name, value},
			"=",
		),
		Headers: map[string]string{
			"Confirm": "true",
		},
	})
	resp.Body = responseBuffer.String()
	responseIsValid := resp.Body == fmt.Sprintf("Property '%s' created with value '%s'", name, value)

	return responseIsValid, resp, err
}

// ListHistorical properties for a given pipeline, stage, job.
func (ps *PropertiesService) ListHistorical(ctx context.Context, pr *PropertyRequest) (*Properties, *APIResponse, error) {
	u := ps.client.BaseURL
	q := u.Query()
	q.Set("pipelineName", pr.Pipeline)
	q.Set("stageName", pr.Stage)
	q.Set("jobName", pr.Job)
	if pr.Limit >= 0 && pr.LimitPipeline != "" {
		q.Set("limitCount", fmt.Sprintf("%d", pr.Limit))
		q.Set("limitPipeline", pr.LimitPipeline)
	}
	u.RawQuery = q.Encode()
	return ps.commonPropertiesAction(ctx, "/properties/search", false)
}

func (ps *PropertiesService) commonPropertiesAction(ctx context.Context, path string, isDatum bool) (*Properties, *APIResponse, error) {
	p := Properties{
		UnmarshallWithHeader: true,
		IsDatum:              isDatum,
	}
	_, resp, err := ps.client.getAction(ctx, &APIClientRequest{
		Path:         path,
		ResponseBody: &p,
	})

	return &p, resp, err
}
