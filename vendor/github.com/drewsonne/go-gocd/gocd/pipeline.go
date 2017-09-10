package gocd

import (
	"context"
	"fmt"
	"net/url"
)

// PipelinesService describes the HAL _link resource for the api response object for a pipelineconfig
type PipelinesService service

// PipelineRequest describes a pipeline request object
type PipelineRequest struct {
	Group    string    `json:"group"`
	Pipeline *Pipeline `json:"pipeline"`
}

// Pipeline describes a pipeline object
type Pipeline struct {
	Group                 string                 `json:"group"`
	Links                 *PipelineLinks         `json:"_links,omitempty"`
	Name                  string                 `json:"name"`
	LabelTemplate         string                 `json:"label_template,omitempty"`
	EnablePipelineLocking bool                   `json:"enable_pipeline_locking,omitempty"`
	Template              string                 `json:"template,omitempty"`
	Origin                *PipelineConfigOrigin  `json:"origin,omitempty"`
	Parameters            []*Parameter           `json:"parameters"`
	EnvironmentVariables  []*EnvironmentVariable `json:"environment_variables"`
	Materials             []Material             `json:"materials,omitempty"`
	Label                 string                 `json:"label,omitempty"`
	Stages                []*Stage               `json:"stages"`
	Version               string                 `json:"version,omitempty"`
	//TrackingTool          string                 `json:"tracking_tool"`
	//Timer                 string                 `json:"timer"`
}

// Parameter represents a key/value
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PipelineConfigOrigin describes where a pipeline config is being loaded from
type PipelineConfigOrigin struct {
	Type string `json:"type"`
	File string `json:"file"`
}

// PipelineLinks describes the HAL _link resource for the api response object for a collection of pipeline objects.
//go:generate gocd-response-links-generator -type=PipelineLinks
type PipelineLinks struct {
	Self *url.URL `json:"self"`
	Doc  *url.URL `json:"doc"`
	Find *url.URL `json:"find"`
}

// Material describes an artifact dependency for a pipeline object.
type Material struct {
	Type        string             `json:"type"`
	Fingerprint string             `json:"fingerprint,omitempty"`
	Description string             `json:"description,omitempty"`
	Attributes  MaterialAttributes `json:"attributes"`
}

// MaterialAttributes describes a material type
type MaterialAttributes struct {
	URL             string          `json:"url"`
	Destination     string          `json:"destination,omitempty"`
	Filter          *MaterialFilter `json:"filter,omitempty"`
	InvertFilter    bool            `json:"invert_filter"`
	Name            string          `json:"name,omitempty"`
	AutoUpdate      bool            `json:"auto_update,omitempty"`
	Branch          string          `json:"branch,omitempty"`
	SubmoduleFolder string          `json:"submodule_folder,omitempty"`
	ShallowClone    bool            `json:"shallow_clone,omitempty"`
	Pipeline        string          `json:"pipeline,omitempty"`
	Stage           string          `json:"stage"`
}

// MaterialFilter describes which globs to ignore
type MaterialFilter struct {
	Ignore []string `json:"ignore"`
}

// PipelineHistory describes the history of runs for a pipeline
type PipelineHistory struct {
	Pipelines []*PipelineInstance `json:"pipelines"`
}

// PipelineInstance describes a single pipeline run
type PipelineInstance struct {
	BuildCause   BuildCause `json:"build_cause"`
	CanRun       bool       `json:"can_run"`
	Name         string     `json:"name"`
	NaturalOrder int        `json:"natural_order"`
	Comment      string     `json:"comment"`
	Stages       []*Stage   `json:"stages"`
}

// BuildCause describes the triggers which caused the build to start.
type BuildCause struct {
	Approver          string             `json:"approver,omitempty"`
	MaterialRevisions []MaterialRevision `json:"material_revisions"`
	TriggerForced     bool               `json:"trigger_forced"`
	TriggerMessage    string             `json:"trigger_message"`
}

// MaterialRevision describes the uniquely identifiable version for the material which was pulled for this build
type MaterialRevision struct {
	Modifications []Modification `json:"modifications"`
	Material      struct {
		Description string `json:"description"`
		Fingerprint string `json:"fingerprint"`
		Type        string `json:"type"`
		ID          int    `json:"id"`
	} `json:"material"`
	Changed bool `json:"changed"`
}

// Modification describes the commit/revision for the material which kicked off the build.
type Modification struct {
	EmailAddress string `json:"email_address"`
	ID           int    `json:"id"`
	ModifiedTime int    `json:"modified_time"`
	UserName     string `json:"user_name"`
	Comment      string `json:"comment"`
	Revision     string `json:"revision"`
}

// PipelineStatus describes whether a pipeline can be run or scheduled.
type PipelineStatus struct {
	Locked      bool `json:"locked"`
	Paused      bool `json:"paused"`
	Schedulable bool `json:"schedulable"`
}

// GetStatus returns a list of pipeline instanves describing the pipeline history.
func (pgs *PipelinesService) GetStatus(ctx context.Context, name string, offset int) (*PipelineStatus, *APIResponse, error) {
	ps := PipelineStatus{}
	_, resp, err := pgs.client.getAction(ctx, &APIClientRequest{
		Path:         fmt.Sprintf("pipelines/%s/status", name),
		ResponseBody: &ps,
	})

	return &ps, resp, err
}

// Pause allows a pipeline to handle new build events
func (pgs *PipelinesService) Pause(ctx context.Context, name string) (bool, *APIResponse, error) {
	return pgs.pipelineAction(ctx, name, "pause")
}

// Unpause allows a pipeline to handle new build events
func (pgs *PipelinesService) Unpause(ctx context.Context, name string) (bool, *APIResponse, error) {
	return pgs.pipelineAction(ctx, name, "unpause")
}

// ReleaseLock frees a pipeline to handle new build events
func (pgs *PipelinesService) ReleaseLock(ctx context.Context, name string) (bool, *APIResponse, error) {
	return pgs.pipelineAction(ctx, name, "releaseLock")
}

// GetInstance of a pipeline run.
func (pgs *PipelinesService) GetInstance(ctx context.Context, name string, offset int) (*PipelineInstance, *APIResponse, error) {
	stub := pgs.buildPaginatedStub("admin/pipelines/%s/instance", name, offset)

	pt := PipelineInstance{}
	_, resp, err := pgs.client.getAction(ctx, &APIClientRequest{
		Path:         stub,
		ResponseBody: &pt,
	})

	return &pt, resp, err
}

// GetHistory returns a list of pipeline instances describing the pipeline history.
func (pgs *PipelinesService) GetHistory(ctx context.Context, name string, offset int) (*PipelineHistory, *APIResponse, error) {
	stub := pgs.buildPaginatedStub("pipelines/%s/history", name, offset)

	pt := PipelineHistory{}
	_, resp, err := pgs.client.getAction(ctx, &APIClientRequest{
		Path:         stub,
		ResponseBody: &pt,
	})

	return &pt, resp, err
}

func (pgs *PipelinesService) pipelineAction(ctx context.Context, name string, action string) (bool, *APIResponse, error) {

	_, resp, err := pgs.client.postAction(ctx, &APIClientRequest{
		Path:         fmt.Sprintf("pipelines/%s/%s", name, action),
		ResponseType: responseTypeJSON,
		Headers: map[string]string{
			"Confirm": "true",
		},
	})

	return resp.HTTP.StatusCode == 200, resp, err
}

func (pgs *PipelinesService) buildPaginatedStub(format string, name string, offset int) string {
	stub := fmt.Sprintf(format, name)
	if offset > 0 {
		stub = fmt.Sprintf("%s/%d", stub, offset)
	}
	return stub
}
