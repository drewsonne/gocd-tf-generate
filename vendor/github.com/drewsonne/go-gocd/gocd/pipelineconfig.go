package gocd

import (
	"context"
	"fmt"
	"strings"
)

// PipelineConfigsService describes the HAL _link resource for the api response object for a pipelineconfig
type PipelineConfigsService service

// PipelineConfigRequest describes a request object for creating or updating pipelines
type PipelineConfigRequest struct {
	Group    string    `json:"group,omitempty"`
	Pipeline *Pipeline `json:"pipeline"`
}

// Get a single PipelineTemplate object in the GoCD API.
func (pcs *PipelineConfigsService) Get(ctx context.Context, name string) (*Pipeline, *APIResponse, error) {
	p := Pipeline{}
	_, resp, err := pcs.client.getAction(ctx, &APIClientRequest{
		Path:         "admin/pipelines/" + name,
		APIVersion:   apiV4,
		ResponseBody: &p,
	})
	if err == nil {
		p.Version = strings.Replace(resp.HTTP.Header.Get("Etag"), "\"", "", -1)
	}

	return &p, resp, err
}

// Update a pipeline configuration
func (pcs *PipelineConfigsService) Update(ctx context.Context, name string, p *Pipeline) (*Pipeline, *APIResponse, error) {

	pt := &PipelineConfigRequest{
		Pipeline: p,
	}
	pr := Pipeline{}

	_, resp, err := pcs.client.putAction(ctx, &APIClientRequest{
		Path:         "admin/pipelines/" + name,
		APIVersion:   apiV4,
		RequestBody:  pt,
		ResponseBody: &pr,
		Headers: map[string]string{
			"If-Match": fmt.Sprintf("\"%s\"", p.Version),
		},
	})

	return &pr, resp, err

}

// Create a pipeline configuration
func (pcs *PipelineConfigsService) Create(ctx context.Context, group string, p *Pipeline) (*Pipeline, *APIResponse, error) {
	pt := &PipelineConfigRequest{
		Group:    group,
		Pipeline: p,
	}
	pc := Pipeline{}

	_, resp, err := pcs.client.postAction(ctx, &APIClientRequest{
		Path:         "admin/pipelines",
		APIVersion:   apiV4,
		RequestBody:  pt,
		ResponseBody: &pc,
	})
	if err == nil {
		pc.Version = strings.Replace(resp.HTTP.Header.Get("Etag"), "\"", "", -1)
	}

	return &pc, resp, err
}

// Delete a pipeline configuration
func (pcs *PipelineConfigsService) Delete(ctx context.Context, name string) (string, *APIResponse, error) {
	return pcs.client.deleteAction(ctx, "admin/pipelines/"+name, apiV4)
}
