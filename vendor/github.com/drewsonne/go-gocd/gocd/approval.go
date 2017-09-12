package gocd

// Approval represents a request/response object describing the approval configuration for a GoCD Job
type Approval struct {
	Type          string         `json:"type,omitempty"`
	Authorization *Authorization `json:"authorization,omitempty"`
}

// Authorization describes the access control for a "manual" approval type. Specifies whoe (role or users) can approve
// the job to move to the next stage of the pipeline.
type Authorization struct {
	Users []string `json:"users,omitempty"`
	Roles []string `json:"roles,omitempty"`
}
