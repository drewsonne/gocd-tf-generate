package gocd

import (
	"context"
	"fmt"
)

// PipelineTemplatesService describes the HAL _link resource for the api response object for a pipeline configuration objects.
type PipelineTemplatesService service

// PipelineTemplateRequest describes a PipelineTemplate
type PipelineTemplateRequest struct {
	Name    string   `json:"name"`
	Stages  []*Stage `json:"stages"`
	Version string   `json:"version"`
}

// PipelineTemplateResponse describes an api response for a single pipeline templates
type PipelineTemplateResponse struct {
	Name     string `json:"name"`
	Embedded *struct {
		Pipelines []*struct {
			Name string `json:"name"`
		}
	} `json:"_embedded,omitempty"`
}

// PipelineTemplatesResponse describes an api response for multiple pipeline templates
type PipelineTemplatesResponse struct {
	Links    *HALLinks `json:"_links,omitempty"`
	Embedded *struct {
		Templates []*PipelineTemplate `json:"templates"`
	} `json:"_embedded,omitempty"`
}

type embeddedPipelineTemplate struct {
	Pipelines []*Pipeline `json:"pipelines,omitempty"`
}

// PipelineTemplate describes a response from the API for a pipeline template object.
type PipelineTemplate struct {
	Links    *HALLinks                 `json:"_links,omitempty"`
	Name     string                    `json:"name"`
	Embedded *embeddedPipelineTemplate `json:"_embedded,omitempty"`
	Version  string                    `json:"template_version"`
	Stages   []*Stage                  `json:"stages,omitempty"`
}

// Get a single PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Get(ctx context.Context, name string) (*PipelineTemplate, *APIResponse, error) {
	pt := PipelineTemplate{}
	_, resp, err := pts.client.getAction(ctx, &APIClientRequest{
		Path:         "admin/templates/" + name,
		APIVersion:   apiV3,
		ResponseBody: &pt,
	})
	//pt.Version = resp.HTTP.Header.Get("Etag")
	//pt.Version = strings.Replace(pt.Version, "\"", "", -1)
	return &pt, resp, err
}

// List all PipelineTemplate objects in the GoCD API.
func (pts *PipelineTemplatesService) List(ctx context.Context) ([]*PipelineTemplate, *APIResponse, error) {
	ptr := PipelineTemplatesResponse{}

	_, resp, err := pts.client.getAction(ctx, &APIClientRequest{
		Path:         "admin/templates",
		APIVersion:   apiV3,
		ResponseBody: &ptr,
	})

	return ptr.Embedded.Templates, resp, err
}

// Create a new PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Create(ctx context.Context, name string, st []*Stage) (*PipelineTemplate, *APIResponse, error) {

	pt := PipelineTemplateRequest{
		Name:   name,
		Stages: st,
	}
	ptr := PipelineTemplate{}

	_, resp, err := pts.client.postAction(ctx, &APIClientRequest{
		Path:         "admin/templates",
		APIVersion:   apiV3,
		RequestBody:  pt,
		ResponseBody: &ptr,
	})

	return &ptr, resp, err

}

// Update an PipelineTemplate object in the GoCD API.
func (pts *PipelineTemplatesService) Update(ctx context.Context, name string, template *PipelineTemplate) (*PipelineTemplate, *APIResponse, error) {
	pt := &PipelineTemplateRequest{
		Name:    name,
		Stages:  template.Stages,
		Version: template.Version,
	}
	ptr := PipelineTemplate{}

	_, resp, err := pts.client.putAction(ctx, &APIClientRequest{
		Path:         "admin/templates/" + name,
		APIVersion:   apiV3,
		RequestBody:  pt,
		ResponseBody: &ptr,
	})

	return &ptr, resp, err

}

// Delete a PipelineTemplate from the GoCD API.
func (pts *PipelineTemplatesService) Delete(ctx context.Context, name string) (string, *APIResponse, error) {
	return pts.client.deleteAction(ctx, fmt.Sprintf("admin/templates/%s", name), apiV3)
}
